<template>
  <div class="ssh-terminal-page">
    <div class="header">
      <div class="host-info">
        <el-icon class="host-icon"><Monitor /></el-icon>
        <div>
          <div class="host-name">{{ hostInfo.hostname }}</div>
          <div class="host-address">{{ hostInfo.ip_address }}:{{ hostInfo.port }}</div>
        </div>
      </div>
      <div class="actions">
        <el-select
          v-model="selectedCredentialId"
          placeholder="选择登录凭据"
          :disabled="connected"
          style="width: 200px; margin-right: 10px;"
        >
          <el-option
            v-for="cred in credentials"
            :key="cred.id"
            :label="`${cred.name} (${cred.username})`"
            :value="cred.id"
          />
        </el-select>
        <el-button 
          type="primary" 
          @click="handleConnect"
          :disabled="connected || !selectedCredentialId"
        >
          <el-icon><Connection /></el-icon>
          连接
        </el-button>
        <el-button 
          type="danger" 
          @click="handleDisconnect" 
          :disabled="!connected"
        >
          <el-icon><Close /></el-icon>
          断开连接
        </el-button>
        <el-button @click="handleClose">
          <el-icon><Back /></el-icon>
          返回
        </el-button>
      </div>
    </div>

    <div class="terminal-container" ref="terminalContainerRef">
      <div v-if="!connected" class="terminal-placeholder">
        <el-icon class="loading-icon" v-if="connecting"><Loading /></el-icon>
        <el-icon class="waiting-icon" v-else><Connection /></el-icon>
        <p>{{ connectionMessage }}</p>
        <p class="hint" v-if="!connecting && selectedCredentialId">请点击"连接"按钮建立 SSH 连接</p>
      </div>
      <div v-show="connected" ref="terminalRef" class="terminal"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { Monitor, Close, Back, Loading, Connection } from '@element-plus/icons-vue'
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import { WebLinksAddon } from 'xterm-addon-web-links'
import 'xterm/css/xterm.css'
import { getHost } from '@/api/host'
import { getCredential, getCredentials } from '@/api/credential'
import type { Host } from '@/api/host'
import type { Credential } from '@/api/credential'

const route = useRoute()
const router = useRouter()

const hostId = ref<number>(parseInt(route.query.host_id as string))
const selectedCredentialId = ref<number>()

const terminalRef = ref<HTMLDivElement>()
const terminalContainerRef = ref<HTMLDivElement>()
const connected = ref(false)
const connecting = ref(false)
const connectionMessage = ref('请选择登录凭据并点击连接')

let terminal: Terminal | null = null
let fitAddon: FitAddon | null = null
let socket: WebSocket | null = null
let hostInfo = ref<Host>({ id: 0, hostname: '', ip_address: '', port: 22, os_type: 'linux', group_id: 0, status: 1, monitoring_enabled: 1, description: '' })
let credentialInfo = ref<Credential>({ id: 0, name: '', username: '', password: '', description: '' })
let credentials = ref<Credential[]>([])

// 初始化终端
const initTerminal = () => {
  terminal = new Terminal({
    theme: {
      background: '#1e1e1e',
      foreground: '#d4d4d4',
      cursor: '#ffffff',
      black: '#000000',
      red: '#cd3131',
      green: '#0dbc79',
      yellow: '#e5e510',
      blue: '#2472c8',
      magenta: '#bc3fbc',
      cyan: '#11a8cd',
      white: '#e5e5e5',
      brightBlack: '#666666',
      brightRed: '#f14c4c',
      brightGreen: '#23d18b',
      brightYellow: '#f5f543',
      brightBlue: '#3b8eea',
      brightMagenta: '#d670d6',
      brightCyan: '#29b8db',
      brightWhite: '#ffffff'
    },
    fontFamily: 'Consolas, Monaco, "Courier New", monospace',
    fontSize: 14,
    lineHeight: 1.5,
    cursorBlink: true,
    cursorStyle: 'block'
  })

  fitAddon = new FitAddon()
  terminal.loadAddon(fitAddon)
  terminal.loadAddon(new WebLinksAddon())

  nextTick(() => {
    if (terminalRef.value) {
      terminal.open(terminalRef.value)
      fitAddon.fit()
    }
  })

  // 监听终端输入
  terminal.onData((data) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.send(JSON.stringify({ type: 'input', data }))
    }
  })

  // 监听窗口大小变化
  window.addEventListener('resize', handleResize)
}

// 处理窗口大小变化
const handleResize = () => {
  if (fitAddon && terminal) {
    fitAddon.fit()
    const dims = { cols: terminal.cols, rows: terminal.rows }
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.send(JSON.stringify({ type: 'resize', ...dims }))
    }
  }
}

// 初始化页面
const initPage = async () => {
  try {
    connectionMessage.value = '正在获取主机信息...'
    
    // 获取主机信息
    const hostRes = await getHost(hostId.value)
    hostInfo.value = hostRes.data

    connectionMessage.value = '正在加载凭据列表...'
    
    // 获取主机关联的所有凭据
    if (hostInfo.value.credentials && hostInfo.value.credentials.length > 0) {
      credentials.value = hostInfo.value.credentials
      // 如果 URL 中有指定的凭据 ID，则选中它
      const urlCredentialId = parseInt(route.query.credential_id as string)
      if (urlCredentialId && credentials.value.some(c => c.id === urlCredentialId)) {
        selectedCredentialId.value = urlCredentialId
      } else {
        // 默认选中第一个凭据
        selectedCredentialId.value = credentials.value[0].id
      }
      
      connectionMessage.value = '准备就绪，请点击连接按钮'
    } else {
      ElMessage.error('该主机没有关联的凭据')
      connectionMessage.value = '无法连接：该主机没有关联的凭据'
    }
  } catch (error) {
    console.error('初始化页面失败:', error)
    ElMessage.error('加载信息失败')
    connectionMessage.value = '加载失败，请刷新页面重试'
  }
}

// 连接 SSH
const handleConnect = async () => {
  if (!selectedCredentialId.value) {
    ElMessage.error('请选择登录凭据')
    return
  }

  connecting.value = true
  connectionMessage.value = '正在连接 SSH 服务器...'
  
  try {
    // 初始化终端
    initTerminal()

    // 建立 WebSocket 连接
    await connectWebSocket()
  } catch (error) {
    console.error('连接 SSH 失败:', error)
    ElMessage.error('连接 SSH 失败')
    connectionMessage.value = '连接失败，请稍后重试'
    connecting.value = false
  }
}

// 建立 WebSocket 连接
const connectWebSocket = () => {
  return new Promise<void>((resolve, reject) => {
    if (!selectedCredentialId.value) {
      reject(new Error('请选择登录凭据'))
      return
    }

    const token = localStorage.getItem('token')
    const wsUrl = `ws://localhost:8080/api/v1/ssh/ws?host_id=${hostId.value}&credential_id=${selectedCredentialId.value}&token=${encodeURIComponent(token!)}`

    socket = new WebSocket(wsUrl)

    socket.onopen = () => {
      connected.value = true
      connecting.value = false
      connectionMessage.value = ''
      ElMessage.success('SSH 连接成功')

      // 发送初始终端尺寸
      if (terminal) {
        const dims = { cols: terminal.cols, rows: terminal.rows }
        socket!.send(JSON.stringify({ type: 'resize', ...dims }))
      }
      
      resolve()
    }

    socket.onmessage = (event) => {
      if (terminal) {
        try {
          const message = JSON.parse(event.data)
          if (message.type === 'output') {
            terminal.write(message.data)
          } else if (message.type === 'error') {
            ElMessage.error(message.data)
          }
        } catch (error) {
          console.error('解析消息失败:', error)
        }
      }
    }

    socket.onerror = (error) => {
      console.error('WebSocket 错误:', error)
      ElMessage.error('SSH 连接错误')
      connected.value = false
      connecting.value = false
      connectionMessage.value = '连接错误'
      reject(error)
    }

    socket.onclose = () => {
      connected.value = false
      connecting.value = false
      ElMessage.info('SSH 连接已断开')
      connectionMessage.value = '连接已断开'
    }
  })
}

// 凭证选择改变
const handleCredentialChange = () => {
  // 凭证改变时，如果已连接，需要提示用户
  if (connected.value) {
    ElMessage.info('请先断开当前连接，然后重新选择凭证')
  }
}

// 断开连接
const handleDisconnect = () => {
  if (socket) {
    socket.send(JSON.stringify({ type: 'close' }))
    socket.close()
  }
}

// 关闭页面
const handleClose = () => {
  handleDisconnect()
  router.back()
}

// 清理资源
onBeforeUnmount(() => {
  handleDisconnect()
  window.removeEventListener('resize', handleResize)
  if (terminal) {
    terminal.dispose()
  }
})

onMounted(() => {
  initPage()
})
</script>

<style scoped>
.ssh-terminal-page {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background: #1e1e1e;
  color: #d4d4d4;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 20px;
  background: #2d2d2d;
  border-bottom: 1px solid #3e3e3e;
}

.host-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.host-icon {
  font-size: 24px;
  color: #0dbc79;
}

.host-name {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 4px;
}

.host-address {
  font-size: 12px;
  color: #888;
}

.actions {
  display: flex;
  gap: 8px;
}

.terminal-container {
  flex: 1;
  overflow: hidden;
  position: relative;
}

.terminal {
  width: 100%;
  height: 100%;
}

.terminal-placeholder {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
  gap: 16px;
}

.loading-icon {
  font-size: 48px;
  animation: rotate 1s linear infinite;
}

.waiting-icon {
  font-size: 48px;
  color: #888;
}

.hint {
  font-size: 12px;
  color: #888;
  margin-top: 8px;
}

@keyframes rotate {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>