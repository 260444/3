<template>
  <el-dialog
    v-model="visible"
    title="SSH 登录"
    width="800px"
    :close-on-click-modal="false"
    :close-on-press-escape="false"
    @close="handleClose"
  >
    <div v-if="!connected" class="login-form">
      <el-form :model="form" label-width="100px">
        <el-form-item label="主机信息">
          <el-descriptions :column="2" border>
            <el-descriptions-item label="主机名">{{ hostInfo.hostname }}</el-descriptions-item>
            <el-descriptions-item label="IP地址">{{ hostInfo.ip_address }}</el-descriptions-item>
            <el-descriptions-item label="端口">{{ hostInfo.port }}</el-descriptions-item>
            <el-descriptions-item label="操作系统">{{ hostInfo.os_type }}</el-descriptions-item>
          </el-descriptions>
        </el-form-item>

        <el-form-item label="选择凭证">
          <el-select
            v-model="form.credential_id"
            placeholder="请选择登录凭证"
            style="width: 100%"
            @change="handleCredentialChange"
          >
            <el-option
              v-for="credential in credentials"
              :key="credential.id"
              :label="`${credential.name} (${credential.username})`"
              :value="credential.id"
            >
              <div style="display: flex; justify-content: space-between;">
                <span>{{ credential.name }}</span>
                <span style="color: #8492a6; font-size: 12px;">{{ credential.username }}</span>
              </div>
            </el-option>
          </el-select>
        </el-form-item>

        <el-form-item v-if="selectedCredential" label="凭证详情">
          <el-alert :title="`用户名: ${selectedCredential.username}`" type="info" :closable="false" />
        </el-form-item>
      </el-form>

      <div class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button :disabled="!form.credential_id" :loading="testing" @click="handleTest">
          测试连接
        </el-button>
        <el-button type="primary" :disabled="!form.credential_id" :loading="connecting" @click="handleConnect">
          连接
        </el-button>
      </div>
    </div>

    <SSHTerminal
      v-else
      :host-info="hostInfo"
      :credential-id="form.credential_id"
      @disconnect="handleDisconnect"
      @reconnect="handleReconnect"
    />
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import SSHTerminal from './SSHTerminal.vue'
import type { Credential } from '@/api/credential'
import { testSSHConnection } from '@/api/ssh'

interface Props {
  modelValue: boolean
  hostInfo: {
    id: number
    hostname: string
    ip_address: string
    port: number
    os_type: string
  }
  credentials: Credential[]
}

const props = defineProps<Props>()
const emit = defineEmits(['update:modelValue'])

const visible = ref(false)
const connected = ref(false)
const connecting = ref(false)
const testing = ref(false)
const form = ref({
  credential_id: undefined as number | undefined
})
const selectedCredential = ref<Credential | null>(null)

watch(() => props.modelValue, (val) => {
  visible.value = val
  if (val) {
    connected.value = false
    form.value.credential_id = undefined
    selectedCredential.value = null
  }
})

watch(visible, (val) => {
  emit('update:modelValue', val)
})

const handleCredentialChange = (credentialId: number) => {
  selectedCredential.value = props.credentials.find(c => c.id === credentialId) || null
}

const handleTest = async () => {
  if (!form.value.credential_id) {
    ElMessage.warning('请选择登录凭证')
    return
  }

  testing.value = true
  try {
    const res = await testSSHConnection({
      host_id: props.hostInfo.id,
      credential_id: form.value.credential_id
    })

    console.log('测试连接响应:', res)

    // 处理响应数据
    // 情况1: 直接返回 {success: true, message: "..."}
    // 情况2: 统一格式 {success: true, message: "...", data: {success: true, ...}}
    let success = false
    let message = '未知错误'

    if (res.success !== undefined) {
      // 统一格式，success 在顶层
      if (res.data && typeof res.data === 'object') {
        // 如果 data 中也有 success，使用 data 中的
        success = res.data.success !== undefined ? res.data.success : res.success
        message = res.data.message || res.message || '未知错误'
      } else {
        success = res.success
        message = res.message || '未知错误'
      }
    } else if (res.data && typeof res.data === 'object') {
      // data 格式
      success = res.data.success
      message = res.data.message || '未知错误'
    }

    if (success) {
      ElMessage.success(message || 'SSH 连接测试成功')
    } else {
      ElMessage.error(`SSH 连接测试失败: ${message}`)
    }
  } catch (error: any) {
    console.error('测试连接错误:', error)
    const errorMsg = error.response?.data?.message || error.response?.data?.error || error.message || '测试失败'
    ElMessage.error(`测试失败: ${errorMsg}`)
  } finally {
    testing.value = false
  }
}

const handleConnect = () => {
  if (!form.value.credential_id) {
    ElMessage.warning('请选择登录凭证')
    return
  }
  connected.value = true
}

const handleDisconnect = () => {
  connected.value = false
  ElMessage.warning('SSH 连接已断开')
}

const handleReconnect = () => {
  connected.value = true
}

const handleClose = () => {
  visible.value = false
  connected.value = false
  form.value.credential_id = undefined
  selectedCredential.value = null
}
</script>

<style scoped>
.login-form {
  padding: 10px 0;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 20px;
}
</style>