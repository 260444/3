<template>
  <div class="host-manage-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>主机管理</span>
          <div>
            <el-button type="danger" :disabled="selection.length === 0" @click="handleBatchDelete">批量删除</el-button>
            <el-button type="primary" @click="handleAdd">新建主机</el-button>
          </div>
        </div>
      </template>
      
      <el-form :inline="true" :model="queryParams" class="demo-form-inline">
        <el-form-item label="主机名">
          <el-input v-model="queryParams.hostname" placeholder="请输入主机名" clearable @keyup.enter="handleQuery" />
        </el-form-item>
        <el-form-item label="IP地址">
          <el-input v-model="queryParams.ip_address" placeholder="请输入IP地址" clearable @keyup.enter="handleQuery" />
        </el-form-item>
        <el-form-item label="主机组">
          <el-select v-model="queryParams.group_id" placeholder="请选择主机组" clearable>
            <el-option
              v-for="item in hostGroupOptions"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="queryParams.status" placeholder="请选择状态" clearable>
            <el-option label="在线" :value="1" />
            <el-option label="离线" :value="0" />
            <el-option label="故障" :value="-1" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleQuery">查询</el-button>
          <el-button @click="resetQuery">重置</el-button>
        </el-form-item>
      </el-form>

      <el-table 
        v-loading="loading" 
        :data="hostList" 
        style="width: 100%"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="hostname" label="主机名" min-width="150" />
        <el-table-column prop="ip_address" label="IP地址" width="140" />
        <el-table-column prop="port" label="端口" width="80" />
        <el-table-column prop="os_type" label="系统" width="100" />
        <el-table-column prop="group.name" label="主机组" width="120" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.status === 1" type="success">在线</el-tag>
            <el-tag v-else-if="scope.row.status === 0" type="info">离线</el-tag>
            <el-tag v-else type="danger">故障</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="monitoring_enabled" label="监控" width="80">
          <template #default="scope">
             <el-switch
                v-model="scope.row.monitoring_enabled"
                :active-value="1"
                :inactive-value="2"
                @change="handleMonitoringChange(scope.row)"
              />
          </template>
        </el-table-column>
        <el-table-column label="操作" width="250" fixed="right">
          <template #default="scope">
            <el-button link type="primary" @click="handleEdit(scope.row)">编辑</el-button>
            <el-button link type="primary" @click="handleDetail(scope.row)">详情</el-button>
            <el-button link type="danger" @click="handleDelete(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination-container">
        <el-pagination
          v-model:current-page="queryParams.page"
          v-model:page-size="queryParams.page_size"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
          @size-change="handleQuery"
          @current-change="handleQuery"
        />
      </div>
    </el-card>

    <!-- 添加/修改对话框 -->
    <el-dialog :title="dialog.title" v-model="dialog.visible" width="600px" append-to-body>
      <el-form ref="formRef" :model="form" :rules="rules" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="主机名" prop="hostname">
              <el-input v-model="form.hostname" placeholder="请输入主机名" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="IP地址" prop="ip_address">
              <el-input v-model="form.ip_address" placeholder="请输入IP地址" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="端口" prop="port">
              <el-input-number v-model="form.port" :min="1" :max="65535" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
             <el-form-item label="操作系统" prop="os_type">
              <el-select v-model="form.os_type" placeholder="请选择操作系统" style="width: 100%">
                <el-option label="Linux" value="linux" />
                <el-option label="Windows" value="windows" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="用户名" prop="username">
              <el-input v-model="form.username" placeholder="请输入用户名" />
            </el-form-item>
          </el-col>
           <el-col :span="12">
            <el-form-item label="密码" prop="password">
              <el-input v-model="form.password" type="password" show-password placeholder="请输入密码" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
             <el-form-item label="主机组" prop="group_id">
              <el-select v-model="form.group_id" placeholder="请选择主机组" style="width: 100%">
                <el-option
                  v-for="item in hostGroupOptions"
                  :key="item.id"
                  :label="item.name"
                  :value="item.id"
                />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-divider content-position="left">硬件配置（可选）</el-divider>
        <el-row :gutter="20">
           <el-col :span="8">
             <el-form-item label="CPU核心" prop="cpu_cores" label-width="80px">
              <el-input-number v-model="form.cpu_cores" :min="1" />
            </el-form-item>
           </el-col>
           <el-col :span="8">
             <el-form-item label="内存(GB)" prop="memory_gb" label-width="80px">
              <el-input-number v-model="form.memory_gb" :min="1" />
            </el-form-item>
           </el-col>
           <el-col :span="8">
             <el-form-item label="磁盘(GB)" prop="disk_space_gb" label-width="80px">
              <el-input-number v-model="form.disk_space_gb" :min="1" />
            </el-form-item>
           </el-col>
        </el-row>

        <el-form-item label="描述" prop="description">
          <el-input v-model="form.description" type="textarea" placeholder="请输入描述信息" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialog.visible = false">取 消</el-button>
          <el-button type="primary" @click="submitForm">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 详情对话框 -->
    <el-dialog title="主机详情" v-model="detailVisible" width="600px">
       <el-descriptions :column="2" border>
        <el-descriptions-item label="主机名">{{ currentHost?.hostname }}</el-descriptions-item>
        <el-descriptions-item label="IP地址">{{ currentHost?.ip_address }}</el-descriptions-item>
        <el-descriptions-item label="端口">{{ currentHost?.port }}</el-descriptions-item>
        <el-descriptions-item label="操作系统">{{ currentHost?.os_type }}</el-descriptions-item>
        <el-descriptions-item label="主机组">{{ currentHost?.group?.name }}</el-descriptions-item>
        <el-descriptions-item label="状态">
           <el-tag v-if="currentHost?.status === 1" type="success">在线</el-tag>
            <el-tag v-else-if="currentHost?.status === 0" type="info">离线</el-tag>
            <el-tag v-else type="danger">故障</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="CPU核心">{{ currentHost?.cpu_cores }}</el-descriptions-item>
        <el-descriptions-item label="内存(GB)">{{ currentHost?.memory_gb }}</el-descriptions-item>
        <el-descriptions-item label="磁盘(GB)">{{ currentHost?.disk_space_gb }}</el-descriptions-item>
        <el-descriptions-item label="最后心跳">{{ formatDate(currentHost?.last_heartbeat) }}</el-descriptions-item>
        <el-descriptions-item label="描述" :span="2">{{ currentHost?.description }}</el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox, type FormInstance } from 'element-plus'
import { 
  getHosts, 
  createHost, 
  updateHost, 
  deleteHost,
  batchDeleteHosts,
  updateHostMonitoring,
  type Host, 
  type HostQuery 
} from '@/api/host'
import { getHostGroups, type HostGroup } from '@/api/hostGroup'

const loading = ref(false)
const total = ref(0)
const hostList = ref<Host[]>([])
const hostGroupOptions = ref<HostGroup[]>([])
const selection = ref<Host[]>([])
const formRef = ref<FormInstance>()
const detailVisible = ref(false)
const currentHost = ref<any>({})

const queryParams = reactive<HostQuery>({
  page: 1,
  page_size: 10,
  hostname: '',
  ip_address: '',
  group_id: undefined,
  status: undefined,
  os_type: ''
})

const dialog = reactive({
  visible: false,
  title: ''
})

const form = reactive({
  id: 0,
  hostname: '',
  ip_address: '',
  port: 22,
  username: '',
  password: '',
  os_type: 'linux',
  cpu_cores: undefined as number | undefined,
  memory_gb: undefined as number | undefined,
  disk_space_gb: undefined as number | undefined,
  group_id: undefined as number | undefined,
  description: ''
})

const rules = {
  hostname: [{ required: true, message: '请输入主机名', trigger: 'blur' }],
  ip_address: [{ required: true, message: '请输入IP地址', trigger: 'blur' }],
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  group_id: [{ required: true, message: '请选择主机组', trigger: 'change' }]
}

const getList = async () => {
  loading.value = true
  try {
    const res = await getHosts(queryParams)
    hostList.value = res.data.list
    total.value = res.data.pagination.total
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

const getGroupList = async () => {
  try {
    const res = await getHostGroups({ page: 1, page_size: 100 })
    hostGroupOptions.value = res.data.list
  } catch (error) {
    console.error(error)
  }
}

const handleQuery = () => {
  queryParams.page = 1
  getList()
}

const resetQuery = () => {
  queryParams.hostname = ''
  queryParams.ip_address = ''
  queryParams.group_id = undefined
  queryParams.status = undefined
  queryParams.os_type = ''
  handleQuery()
}

const handleSelectionChange = (val: Host[]) => {
  selection.value = val
}

const handleAdd = () => {
  dialog.title = '添加主机'
  dialog.visible = true
  form.id = 0
  form.hostname = ''
  form.ip_address = ''
  form.port = 22
  form.username = ''
  form.password = ''
  form.os_type = 'linux'
  form.cpu_cores = undefined
  form.memory_gb = undefined
  form.disk_space_gb = undefined
  form.group_id = undefined
  form.description = ''
  
  setTimeout(() => {
    formRef.value?.clearValidate()
  }, 0)
}

const handleEdit = (row: Host) => {
  dialog.title = '修改主机'
  dialog.visible = true
  form.id = row.id
  form.hostname = row.hostname
  form.ip_address = row.ip_address
  form.port = row.port
  form.username = row.username
  form.password = '' // 编辑时不显示密码，留空表示不修改
  form.os_type = row.os_type
  form.cpu_cores = row.cpu_cores
  form.memory_gb = row.memory_gb
  form.disk_space_gb = row.disk_space_gb
  form.group_id = row.group_id
  form.description = row.description
}

const submitForm = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (form.id === 0) {
           if (!form.password) {
             ElMessage.warning('创建主机时密码不能为空')
             return
           }
          await createHost(form)
          ElMessage.success('添加成功')
        } else {
          await updateHost(form.id, form)
          ElMessage.success('修改成功')
        }
        dialog.visible = false
        getList()
      } catch (error) {
        console.error(error)
      }
    }
  })
}

const handleDelete = (row: Host) => {
  ElMessageBox.confirm('确认要删除该主机吗？', '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await deleteHost(row.id)
      ElMessage.success('删除成功')
      getList()
    } catch (error) {
      console.error(error)
    }
  })
}

const handleBatchDelete = () => {
  ElMessageBox.confirm(`确认要删除选中的 ${selection.value.length} 个主机吗？`, '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const ids = selection.value.map(item => item.id)
      await batchDeleteHosts(ids)
      ElMessage.success('批量删除成功')
      getList()
    } catch (error) {
      console.error(error)
    }
  })
}

const handleMonitoringChange = async (row: Host) => {
   // row.monitoring_enabled 已经由 v-model 更新
   // 这里需要发送请求到后端
   try {
     await updateHostMonitoring(row.id, row.monitoring_enabled)
     ElMessage.success('监控状态更新成功')
   } catch (error) {
     // 如果失败，回滚状态
     row.monitoring_enabled = row.monitoring_enabled === 1 ? 2 : 1
     console.error(error)
   }
}

const handleDetail = (row: Host) => {
  currentHost.value = row
  detailVisible.value = true
}

const formatDate = (dateStr: string | undefined) => {
  if (!dateStr) return '-'
  return new Date(dateStr).toLocaleString()
}

onMounted(() => {
  getGroupList()
  getList()
})
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.pagination-container {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>
