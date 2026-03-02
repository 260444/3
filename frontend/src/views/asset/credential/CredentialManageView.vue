<template>
  <div class="credential-manage-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>凭据管理</span>
          <div>
            <el-button type="danger" :disabled="selection.length === 0" @click="handleBatchDelete">批量删除</el-button>
            <el-button type="primary" @click="handleAdd">新建凭据</el-button>
          </div>
        </div>
      </template>
      
      <el-form :inline="true" :model="queryParams" class="demo-form-inline">
        <el-form-item label="凭据名称">
          <el-input v-model="queryParams.name" placeholder="请输入凭据名称" clearable @keyup.enter="handleQuery" />
        </el-form-item>
        <el-form-item label="用户名">
          <el-input v-model="queryParams.username" placeholder="请输入用户名" clearable @keyup.enter="handleQuery" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleQuery">查询</el-button>
          <el-button @click="resetQuery">重置</el-button>
        </el-form-item>
      </el-form>
      
      <el-table 
        v-loading="loading" 
        :data="credentialList" 
        style="width: 100%"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="name" label="凭据名称" min-width="150" />
        <el-table-column prop="username" label="用户名" width="120" />
        <el-table-column label="密码" width="150">
          <template #default="scope">
            <span v-if="scope.row.password && scope.row.password.length > 0">●●●●●●</span>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述" min-width="200" />
        <el-table-column prop="created_at" label="创建时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.created_at) }}
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
        <el-form-item label="凭据名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入凭据名称" />
        </el-form-item>
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" placeholder="请输入用户名" />
        </el-form-item>
        <el-form-item v-if="dialog.type === 'add' || form.password !== undefined" :label="dialog.type === 'add' ? '密码' : '新密码'" :prop="dialog.type === 'add' ? 'password' : ''" :rules="dialog.type === 'add' ? rules.password : []">
          <el-input v-model="form.password" type="password" show-password :placeholder="dialog.type === 'add' ? '请输入密码' : '留空则不修改密码'" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="form.description" type="textarea" :rows="4" placeholder="请输入描述信息" />
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
    
        <el-dialog title="凭据详情" v-model="detailVisible" width="600px">
    
          <el-descriptions :column="2" border>
    
            <el-descriptions-item label="凭据名称">{{ currentCredential?.name }}</el-descriptions-item>
    
            <el-descriptions-item label="用户名">{{ currentCredential?.username }}</el-descriptions-item>
    
            <el-descriptions-item label="密码">
          <span v-if="currentCredential?.password && currentCredential?.password.length > 0">●●●●●●</span>
          <span v-else>-</span>
        </el-descriptions-item>
    
            <el-descriptions-item label="描述" :span="2">{{ currentCredential?.description }}</el-descriptions-item>
    
            <el-descriptions-item label="创建时间">{{ formatDate(currentCredential?.created_at) }}</el-descriptions-item>
    
            <el-descriptions-item label="更新时间">{{ formatDate(currentCredential?.updated_at) }}</el-descriptions-item>
    
          </el-descriptions>
    
        </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox, type FormInstance } from 'element-plus'
import { 
  getCredentials, 
  createCredential, 
  updateCredential, 
  deleteCredential,
  batchDeleteCredentials,
  type Credential, 
  type CredentialQuery,
  type CreateCredentialReq,
  type UpdateCredentialReq
} from '@/api/credential'

const loading = ref(false)
const total = ref(0)
const credentialList = ref<Credential[]>([])
const selection = ref<Credential[]>([])
const formRef = ref<FormInstance>()
const detailVisible = ref(false)
const currentCredential = ref<Credential | null>(null)

const queryParams = reactive<CredentialQuery>({
  page: 1,
  page_size: 10,
  name: '',
  username: ''
})

const dialog = reactive({
  visible: false,
  title: '',
  type: 'add' as 'add' | 'edit'
})

const form = reactive<CreateCredentialReq | UpdateCredentialReq>({
  name: '',
  username: '',
  password: '',
  description: ''
})

const rules = {
  name: [{ required: true, message: '请输入凭据名称', trigger: 'blur' }],
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度至少6位', trigger: 'blur' }
  ]
}

const getList = async () => {
  loading.value = true
  try {
    const res = await getCredentials(queryParams)
    credentialList.value = res.data.list
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
  queryParams.username = ''
  handleQuery()
}

const handleSelectionChange = (val: Credential[]) => {
  selection.value = val
}

const handleAdd = () => {
  dialog.title = '添加凭据'
  dialog.type = 'add'
  dialog.visible = true
  form.id = 0 // 新增时设置ID为0
  form.name = ''
  form.username = ''
  form.password = ''
  form.description = ''
  
  setTimeout(() => {
    formRef.value?.clearValidate()
  }, 0)
}

const handleEdit = (row: Credential) => {
  dialog.title = '修改凭据'
  dialog.type = 'edit'
  dialog.visible = true
  form.id = row.id // 确保设置凭据ID
  form.name = row.name
  form.username = row.username
  form.password = '' // 编辑时不预设密码，如果需要修改则输入新密码，留空表示不修改
  form.description = row.description
  
  setTimeout(() => {
    formRef.value?.clearValidate()
  }, 0)
}

const submitForm = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (dialog.type === 'add') {
          if (!form.password) {
            ElMessage.error('创建凭据时密码不能为空')
            return
          }
          await createCredential(form as CreateCredentialReq)
          ElMessage.success('添加成功')
        } else {
          // 编辑模式下，只有当密码不为空时才更新密码
          const updateData = { ...form } as Partial<UpdateCredentialReq>;
          if (!updateData.password || updateData.password === '') {
            // 如果密码为空，从更新数据中移除password字段，表示不更新密码
            delete updateData.password;
          }
          await updateCredential(form.id as number, updateData as UpdateCredentialReq)
          ElMessage.success('修改成功')
        }
        dialog.visible = false
        getList()
      } catch (error) {
        console.error(error)
        ElMessage.error('操作失败')
      }
    }
  })
}

const handleDelete = (row: Credential) => {
  ElMessageBox.confirm('确认要删除该凭据吗？删除后将无法恢复。', '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await deleteCredential(row.id)
      ElMessage.success('删除成功')
      getList()
    } catch (error) {
      console.error(error)
      ElMessage.error('删除失败')
    }
  })
}

const handleBatchDelete = () => {
  if (selection.value.length === 0) {
    ElMessage.warning('请选择要删除的凭据')
    return
  }
  
  ElMessageBox.confirm(`确认要批量删除选中的 ${selection.value.length} 个凭据吗？删除后将无法恢复。`, '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const ids = selection.value.map(item => item.id)
      const res = await batchDeleteCredentials(ids)
      ElMessage.success(`批量删除成功，共删除 ${res.data.deleted_count} 个凭据`)
      getList()
    } catch (error) {
      console.error(error)
      ElMessage.error('批量删除失败')
    }
  })
}

const handleDetail = (row: Credential) => {
  currentCredential.value = row
  detailVisible.value = true
}

const formatDate = (dateStr: string | undefined) => {
  if (!dateStr) return '-'
  return new Date(dateStr).toLocaleString('zh-CN')
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
.dialog-footer {
  text-align: right;
}
</style>