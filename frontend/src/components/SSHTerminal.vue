<template>
  <div class="ssh-terminal-container">
    <div class="terminal-header">
      <div class="terminal-title">
        <el-icon><Monitor /></el-icon>
        <span>{{ hostInfo.hostname }} ({{ hostInfo.ip_address }}:{{ hostInfo.port }})</span>
      </div>
      <div class="terminal-actions">
        <el-button size="small" @click="handleCopy">复制</el-button>
        <el-button size="small" @click="handlePaste">粘贴</el-button>
        <el-button size="small" type="danger" @click="handleDisconnect">断开连接</el-button>
      </div>
    </div>
    <div ref="terminalRef" class="terminal-content" @contextmenu.prevent="handleContextMenu"></div>
    <div v-if="!connected" class="terminal-overlay">
      <el-result icon="warning" title="连接已断开" sub-title="请重新连接">
        <template #extra>
          <el-button type="primary" @click="$emit('reconnect')">重新连接</el-button>
        </template>
      </el-result>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Monitor } from '@element-plus/icons-vue'
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import { WebLinksAddon } from 'xterm-addon-web-links'

interface Props {
  hostInfo: {
    hostname: string
    ip_address: string
    port: number
  }
  credentialId: number
}

const props = defineProps<Props>()
const emit = defineEmits(['disconnect', 'reconnect'])

const terminalRef = ref<HTMLElement>()
let terminal: Terminal | null = null
let fitAddon: FitAddon | null = null
let socket: WebSocket | null = null
let connected = ref(true)

// 初始化终端
const initTerminal = () => {
  if (!terminalRef.value) return

  terminal = new Terminal({
    cursorBlink: true,
    fontSize: 14,
    fontFamily: 'Consolas, Monaco, "Courier New", monospace',
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
    }
  })

  fitAddon = new FitAddon()
  terminal.loadAddon(fitAddon)
  terminal.loadAddon(new WebLinksAddon())

  terminal.open(terminalRef.value)
  fitAddon.fit()

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

// 连接 SSH WebSocket
const connect = () => {
  // WebSocket 需要直接连接后端，因为 Vite 代理不支持 WebSocket
  const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
  const host = 'localhost' // 直接连接后端服务器
  const port = '8080'      // 后端端口

  // 获取 JWT token
  const token = localStorage.getItem('token')
  console.log('Token:', token ? token.substring(0, 20) + '...' : '未找到')

  // 构建 WebSocket URL，直接连接后端
  const wsUrl = `${protocol}//${host}:${port}/api/v1/ssh/ws?host_id=${props.hostInfo.id}&credential_id=${props.credentialId}&token=${encodeURIComponent(token)}`

  console.log('连接 WebSocket:', wsUrl)
  console.log('WebSocket 就绪状态:', WebSocket ? '支持' : '不支持')

  try {
    socket = new WebSocket(wsUrl)

    // 设置连接超时
    const timeout = setTimeout(() => {
      if (socket && socket.readyState !== WebSocket.OPEN) {
        console.error('WebSocket 连接超时')
        ElMessage.error('SSH 连接超时，请检查网络')
        connected.value = false
      }
    }, 10000) // 10秒超时

    socket.onopen = () => {
      clearTimeout(timeout)
      connected.value = true
      ElMessage.success('SSH 连接成功')
      console.log('WebSocket 连接已建立，状态:', socket.readyState)

      // 清空终端并显示欢迎信息
      if (terminal) {
        terminal.clear()

        // 延迟一下，确保终端准备好
        setTimeout(() => {
          terminal.writeln('\x1b[32m正在连接 SSH 服务器...\x1b[0m')

          // 发送初始终端尺寸
          const dims = { cols: terminal.cols, rows: terminal.rows }
          socket.send(JSON.stringify({ type: 'resize', ...dims }))
          console.log('发送终端尺寸:', dims)
        }, 100)
      }
    }

    socket.onmessage = (event) => {
      if (terminal) {
        try {
          const message = JSON.parse(event.data)
          console.log('收到消息:', message)

          if (message.type === 'output') {
            terminal.write(message.data)
          } else if (message.type === 'error') {
            ElMessage.error(message.data)
            connected.value = false
          }
        } catch (error) {
          console.error('解析消息失败:', error)
        }
      }
    }

    socket.onerror = (error) => {
      clearTimeout(timeout)
      console.error('WebSocket 错误:', error)
      console.error('WebSocket 状态:', socket ? socket.readyState : '未创建')
      ElMessage.error('SSH 连接错误')
      connected.value = false
    }

    socket.onclose = (event) => {
      clearTimeout(timeout)
      console.log('WebSocket 关闭:', event.code, event.reason)
      console.log('关闭时状态:', event.code)
      connected.value = false
      emit('disconnect')
    }
  } catch (error) {
    console.error('创建 WebSocket 失败:', error)
    ElMessage.error('无法创建 WebSocket 连接')
    connected.value = false
  }
}

// 断开连接
const handleDisconnect = () => {
  if (socket) {
    socket.close()
  }
}

// 复制
const handleCopy = async () => {
  try {
    const selection = terminal?.getSelection()
    if (selection) {
      await navigator.clipboard.writeText(selection)
      ElMessage.success('复制成功')
    }
  } catch (error) {
    ElMessage.error('复制失败')
  }
}

// 粘贴
const handlePaste = async () => {
  try {
    const text = await navigator.clipboard.readText()
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.send(JSON.stringify({ type: 'input', data: text }))
    }
  } catch (error) {
    ElMessage.error('粘贴失败')
  }
}

// 右键菜单
const handleContextMenu = (event: MouseEvent) => {
  event.preventDefault()
  // 可以在这里添加自定义右键菜单
}

// 监听凭证变化
watch(() => props.credentialId, () => {
  if (socket) {
    socket.close()
  }
  connect()
})

onMounted(() => {
  initTerminal()
  connect()
})

onBeforeUnmount(() => {
  if (socket) {
    socket.close()
  }
  if (terminal) {
    terminal.dispose()
  }
  window.removeEventListener('resize', handleResize)
})
</script>

<style scoped>
.ssh-terminal-container {
  position: relative;
  width: 100%;
  height: 600px;
  background: #1e1e1e;
  border-radius: 4px;
  overflow: hidden;
}

.terminal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 15px;
  background: #2d2d2d;
  border-bottom: 1px solid #3e3e3e;
}

.terminal-title {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #d4d4d4;
  font-size: 14px;
}

.terminal-actions {
  display: flex;
  gap: 8px;
}

.terminal-content {
  width: 100%;
  height: calc(100% - 48px);
  padding: 10px;
}

.terminal-overlay {
  position: absolute;
  top: 48px;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(30, 30, 30, 0.95);
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>