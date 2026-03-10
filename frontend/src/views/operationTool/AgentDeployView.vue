<template>
  <div class="agent-deploy-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>Agent 部署</span>
          <el-button type="primary" @click="loadUndeployedHosts" :icon="Refresh">
            刷新列表
          </el-button>
        </div>
      </template>

      <el-alert
        title="说明"
        type="info"
        :closable="false"
        show-icon
        style="margin-bottom: 20px"
      >
        此页面用于向未部署监控的主机部署 Agent。请选择需要部署的主机，然后选择对应的凭据进行部署。
      </el-alert>

      <el-table
        v-loading="loading"
        :data="undeployedHosts"
        style="width: 100%"
        @selection-change="handleSelectionChange"
        empty-text="暂无未部署监控的主机"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="hostname" label="主机名" width="150" />
        <el-table-column prop="ip_address" label="IP地址" width="140" />
        <el-table-column prop="port" label="端口" width="80" />
        <el-table-column prop="os_type" label="操作系统" width="100">
          <template #default="{ row }">
            <el-tag :type="row.os_type === 'linux' ? 'success' : 'warning'" size="small">
              {{ row.os_type }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="group.name" label="所属主机组" width="150" />
        <el-table-column label="可用凭据" min-width="200">
          <template #default="{ row }">
            <div v-if="row.credentials && row.credentials.length > 0">
              <el-tag
                v-for="cred in row.credentials"
                :key="cred.id"
                size="small"
                style="margin-right: 5px; margin-bottom: 5px"
              >
                {{ cred.name }} ({{ cred.username }})
              </el-tag>
            </div>
            <el-text v-else type="danger" size="small">暂无可用凭据</el-text>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)" size="small">
              {{ getStatusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button
              type="primary"
              size="small"
              :disabled="!row.credentials || row.credentials.length === 0"
              @click="handleDeploySingle(row)"
            >
              部署
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <div v-if="selection.length > 0" style="margin-top: 20px">
        <el-divider>批量部署</el-divider>
        <div style="display: flex; align-items: center; gap: 15px">
          <span>已选择 {{ selection.length }} 台主机</span>
          <el-button
            type="success"
            :loading="deploying"
            :disabled="!canBatchDeploy"
            @click="handleBatchDeploy"
          >
            批量部署
          </el-button>
          <el-text v-if="!canBatchDeploy" type="warning" size="small">
            部分主机没有可用凭据，无法批量部署
          </el-text>
        </div>
      </div>
    </el-card>

    <!-- 单个主机部署对话框 -->
    <el-dialog
      v-model="deployDialogVisible"
      title="选择凭据部署 Agent"
      width="500px"
      :close-on-click-modal="false"
    >
      <el-form :model="deployForm" label-width="80px">
        <el-form-item label="主机信息">
          <el-descriptions :column="1" border size="small">
            <el-descriptions-item label="主机名">{{ deployForm.hostname }}</el-descriptions-item>
            <el-descriptions-item label="IP地址">{{ deployForm.ip_address }}</el-descriptions-item>
          </el-descriptions>
        </el-form-item>
        <el-form-item label="选择凭据" required>
          <el-select
            v-model="deployForm.credential_id"
            placeholder="请选择凭据"
            style="width: 100%"
          >
            <el-option
              v-for="cred in deployForm.credentials"
              :key="cred.id"
              :label="`${cred.name} (${cred.username})`"
              :value="cred.id"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="deployDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="deploying" @click="confirmDeploy">确认部署</el-button>
      </template>
    </el-dialog>

    <!-- 批量部署对话框 -->
    <el-dialog
      v-model="batchDeployDialogVisible"
      title="批量部署 Agent"
      width="800px"
      :close-on-click-modal="false"
    >
      <el-alert
        title="批量部署将为每台主机使用其第一个可用凭据进行部署"
        type="warning"
        :closable="false"
        show-icon
        style="margin-bottom: 20px"
      />
      <el-table :data="selection" style="width: 100%" max-height="400">
        <el-table-column prop="hostname" label="主机名" width="150" />
        <el-table-column prop="ip_address" label="IP地址" width="140" />
        <el-table-column label="使用凭据" min-width="200">
          <template #default="{ row }">
            <el-tag size="small">
              {{ row.credentials[0].name }} ({{ row.credentials[0].username }})
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="部署状态" width="120">
          <template #default="{ row }">
            <el-tag v-if="batchDeployStatus[row.id] === 'success'" type="success" size="small">
              成功
            </el-tag>
            <el-tag v-else-if="batchDeployStatus[row.id] === 'failed'" type="danger" size="small">
              失败
            </el-tag>
            <el-tag v-else type="info" size="small">待部署</el-tag>
          </template>
        </el-table-column>
      </el-table>
      <template #footer>
        <el-button @click="batchDeployDialogVisible = false" :disabled="deploying">
          取消
        </el-button>
        <el-button type="primary" :loading="deploying" @click="confirmBatchDeploy">
          开始部署
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Refresh } from '@element-plus/icons-vue'
import { getUndeployedHosts, updateHostMonitoringDeploy, deployAgent } from '@/api/host'

interface Credential {
  id: number
  name: string
  username: string
}

interface Host {
  id: number
  hostname: string
  ip_address: string
  port: number
  os_type: string
  status: number
  group?: {
    name: string
  }
  credentials?: Credential[]
}

const loading = ref(false)
const deploying = ref(false)
const undeployedHosts = ref<Host[]>([])
const selection = ref<Host[]>([])

// 单个部署对话框
const deployDialogVisible = ref(false)
const deployForm = ref({
  host_id: 0,
  hostname: '',
  ip_address: '',
  credential_id: 0,
  credentials: [] as Credential[]
})

// 批量部署对话框
const batchDeployDialogVisible = ref(false)
const batchDeployStatus = ref<Record<number, 'success' | 'failed'>>({})

// 检查是否可以批量部署
const canBatchDeploy = computed(() => {
  return selection.value.length > 0 && selection.value.every(host => 
    host.credentials && host.credentials.length > 0
  )
})

// 加载未部署监控的主机
const loadUndeployedHosts = async () => {
  try {
    loading.value = true
    const response: any = await getUndeployedHosts()
    undeployedHosts.value = response.data || []
  } catch (error: any) {
    ElMessage.error(error.message || '加载主机列表失败')
  } finally {
    loading.value = false
  }
}

// 选择变化
const handleSelectionChange = (val: Host[]) => {
  selection.value = val
}

// 获取状态类型
const getStatusType = (status: number) => {
  switch (status) {
    case 1:
      return 'success'
    case 0:
      return 'info'
    case -1:
      return 'danger'
    default:
      return 'info'
  }
}

// 获取状态文本
const getStatusText = (status: number) => {
  switch (status) {
    case 1:
      return '在线'
    case 0:
      return '离线'
    case -1:
      return '故障'
    default:
      return '未知'
  }
}

// 单个主机部署
const handleDeploySingle = (host: Host) => {
  deployForm.value = {
    host_id: host.id,
    hostname: host.hostname,
    ip_address: host.ip_address,
    credential_id: host.credentials?.[0]?.id || 0,
    credentials: host.credentials || []
  }
  deployDialogVisible.value = true
}

// 确认单个部署
const confirmDeploy = async () => {
  if (!deployForm.value.credential_id) {
    ElMessage.warning('请选择凭据')
    return
  }

  try {
    deploying.value = true
    
    // 调用部署接口
    await deployAgent(deployForm.value.host_id, deployForm.value.credential_id)
    
    // 更新监控部署状态
    await updateHostMonitoringDeploy(deployForm.value.host_id, 1)
    
    ElMessage.success(`主机 ${deployForm.value.hostname} 部署成功`)
    deployDialogVisible.value = false
    
    // 刷新列表
    await loadUndeployedHosts()
  } catch (error: any) {
    ElMessage.error(error.message || `主机 ${deployForm.value.hostname} 部署失败`)
  } finally {
    deploying.value = false
  }
}

// 批量部署
const handleBatchDeploy = () => {
  batchDeployStatus.value = {}
  batchDeployDialogVisible.value = true
}

// 确认批量部署
const confirmBatchDeploy = async () => {
  try {
    deploying.value = true
    const total = selection.value.length
    let successCount = 0
    let failCount = 0

    for (const host of selection.value) {
      try {
        // 使用第一个可用凭据进行部署
        const credentialId = host.credentials![0].id
        
        // 调用部署接口
        await deployAgent(host.id, credentialId)
        
        // 更新监控部署状态
        await updateHostMonitoringDeploy(host.id, 1)
        
        batchDeployStatus.value[host.id] = 'success'
        successCount++
      } catch (error: any) {
        batchDeployStatus.value[host.id] = 'failed'
        failCount++
        console.error(`主机 ${host.hostname} 部署失败:`, error)
      }
    }

    if (successCount === total) {
      ElMessage.success(`批量部署成功，共部署 ${successCount} 台主机`)
    } else if (successCount > 0) {
      ElMessage.warning(`批量部署完成，成功 ${successCount} 台，失败 ${failCount} 台`)
    } else {
      ElMessage.error(`批量部署失败，所有主机部署均失败`)
    }

    // 刷新列表
    await loadUndeployedHosts()
  } catch (error: any) {
    ElMessage.error(error.message || '批量部署失败')
  } finally {
    deploying.value = false
    batchDeployDialogVisible.value = false
  }
}

onMounted(() => {
  loadUndeployedHosts()
})
</script>

<style scoped>
.agent-deploy-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>