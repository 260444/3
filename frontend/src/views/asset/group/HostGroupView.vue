<template>
  <div class="host-group-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>主机组管理</span>
          <el-button type="primary" @click="handleAdd">新建主机组</el-button>
        </div>
      </template>
      
      <el-form :inline="true" :model="queryParams" class="demo-form-inline">
        <el-form-item label="名称">
          <el-input v-model="queryParams.name" placeholder="请输入主机组名称" clearable @keyup.enter="handleQuery" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="queryParams.status" placeholder="请选择状态" clearable>
            <el-option label="启用" :value="1" />
            <el-option label="禁用" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleQuery">查询</el-button>
          <el-button @click="resetQuery">重置</el-button>
        </el-form-item>
      </el-form>

      <el-table v-loading="loading" :data="hostGroupList" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="名称" min-width="150" />
        <el-table-column prop="description" label="描述" min-width="200" show-overflow-tooltip />
        <el-table-column prop="host_count" label="主机数量" width="100" align="center" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
              {{ scope.row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="250" fixed="right">
          <template #default="scope">
            <el-button link type="primary" @click="handleEdit(scope.row)">编辑</el-button>
            <el-button 
              link 
              :type="scope.row.status === 1 ? 'warning' : 'success'" 
              @click="handleStatusChange(scope.row)"
            >
              {{ scope.row.status === 1 ? '禁用' : '启用' }}
            </el-button>
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
    <el-dialog :title="dialog.title" v-model="dialog.visible" width="500px" append-to-body>
      <el-form ref="formRef" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入主机组名称" />
        </el-form-item>
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
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox, type FormInstance } from 'element-plus'
import { 
  getHostGroups, 
  createHostGroup, 
  updateHostGroup, 
  deleteHostGroup,
  updateHostGroupStatus,
  type HostGroup, 
  type HostGroupQuery 
} from '@/api/hostGroup'

const loading = ref(false)
const total = ref(0)
const hostGroupList = ref<HostGroup[]>([])
const formRef = ref<FormInstance>()

const queryParams = reactive<HostGroupQuery>({
  page: 1,
  page_size: 10,
  name: '',
  status: undefined
})

const dialog = reactive({
  visible: false,
  title: ''
})

const form = reactive({
  id: 0,
  name: '',
  description: ''
})

const rules = {
  name: [{ required: true, message: '请输入主机组名称', trigger: 'blur' }]
}

const getList = async () => {
  loading.value = true
  try {
    const res = await getHostGroups(queryParams)
    hostGroupList.value = res.data.list
    total.value = res.data.pagination.total
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

const handleQuery = () => {
  queryParams.page = 1
  getList()
}

const resetQuery = () => {
  queryParams.name = ''
  queryParams.status = undefined
  handleQuery()
}

const handleAdd = () => {
  dialog.title = '添加主机组'
  dialog.visible = true
  form.id = 0
  form.name = ''
  form.description = ''
  // 重置表单校验
  setTimeout(() => {
    formRef.value?.clearValidate()
  }, 0)
}

const handleEdit = (row: HostGroup) => {
  dialog.title = '修改主机组'
  dialog.visible = true
  form.id = row.id
  form.name = row.name
  form.description = row.description
}

const submitForm = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (form.id === 0) {
          await createHostGroup({
            name: form.name,
            description: form.description
          })
          ElMessage.success('添加成功')
        } else {
          await updateHostGroup(form.id, {
            name: form.name,
            description: form.description
          })
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

const handleDelete = (row: HostGroup) => {
  ElMessageBox.confirm('确认要删除该主机组吗？', '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await deleteHostGroup(row.id)
      ElMessage.success('删除成功')
      getList()
    } catch (error) {
      console.error(error)
    }
  })
}

const handleStatusChange = async (row: HostGroup) => {
  const newStatus = row.status === 1 ? 0 : 1
  const actionText = newStatus === 1 ? '启用' : '禁用'
  
  ElMessageBox.confirm(`确认要${actionText}该主机组吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await updateHostGroupStatus(row.id, newStatus)
      ElMessage.success(`${actionText}成功`)
      getList()
    } catch (error) {
      console.error(error)
    }
  })
}

const formatDate = (dateStr: string) => {
  if (!dateStr) return ''
  return new Date(dateStr).toLocaleString()
}

onMounted(() => {
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
