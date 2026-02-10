<template>
  <div class="credential-manage">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>凭据管理</span>
          <el-button type="primary" @click="handleCreate">新建凭据</el-button>
        </div>
      </template>

      <!-- 搜索条件 -->
      <el-form :model="searchForm" label-width="80px" class="search-form">
        <el-row :gutter="20">
          <el-col :span="6">
            <el-form-item label="凭据名称">
              <el-input v-model="searchForm.name" placeholder="请输入凭据名称" clearable />
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item label="用户名">
              <el-input v-model="searchForm.username" placeholder="请输入用户名" clearable />
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item label="创建时间">
              <el-date-picker
                v-model="searchForm.createTime"
                type="daterange"
                range-separator="至"
                start-placeholder="开始日期"
                end-placeholder="结束日期"
                value-format="YYYY-MM-DD"
              />
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item>
              <el-button type="primary" @click="handleSearch">搜索</el-button>
              <el-button @click="handleReset">重置</el-button>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>

      <!-- 表格数据 -->
      <el-table
        :data="tableData"
        style="width: 100%"
        v-loading="loading"
        border
        stripe
      >
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="凭据名称" min-width="150" />
        <el-table-column prop="username" label="用户名" min-width="120" />
        <el-table-column prop="description" label="描述" min-width="200" />
        <el-table-column prop="createdAt" label="创建时间" width="180">
          <template #default="{ row }">
            {{ formatDate(row.createdAt) }}
          </template>
        </el-table-column>
        <el-table-column prop="updatedAt" label="更新时间" width="180">
          <template #default="{ row }">
            {{ formatDate(row.updatedAt) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button size="small" type="danger" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="pagination">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :page-sizes="[10, 20, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>

    <!-- 新建/编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="500px"
      @close="handleDialogClose"
    >
      <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        label-width="100px"
      >
        <el-form-item label="凭据名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入凭据名称" />
        </el-form-item>
        <el-form-item label="用户名" prop="username">
          <el-input v-model="formData.username" placeholder="请输入用户名" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input
            v-model="formData.password"
            type="password"
            placeholder="请输入密码"
            show-password
          />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input
            v-model="formData.description"
            type="textarea"
            placeholder="请输入描述信息"
            :rows="3"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit" :loading="submitLoading">
            确定
          </el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 删除确认对话框 -->
    <el-dialog
      v-model="deleteDialogVisible"
      title="确认删除"
      width="400px"
    >
      <span>确定要删除凭据 "{{ deleteItem?.name }}" 吗？</span>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="deleteDialogVisible = false">取消</el-button>
          <el-button type="danger" @click="confirmDelete" :loading="deleteLoading">
            确定
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import { credentialApi } from '@/api/credential'

// 表格数据
const tableData = ref([])
const loading = ref(false)
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)

// 搜索表单
const searchForm = reactive({
  name: '',
  username: '',
  createTime: []
})

// 对话框相关
const dialogVisible = ref(false)
const dialogTitle = ref('')
const submitLoading = ref(false)
const formRef = ref<FormInstance>()
const deleteDialogVisible = ref(false)
const deleteLoading = ref(false)
const deleteItem = ref<any>(null)

// 表单数据
const formData = reactive({
  id: undefined,
  name: '',
  username: '',
  password: '',
  description: ''
})

// 表单验证规则
const formRules = reactive<FormRules>({
  name: [
    { required: true, message: '请输入凭据名称', trigger: 'blur' },
    { min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur' }
  ],
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, max: 100, message: '长度在 6 到 100 个字符', trigger: 'blur' }
  ]
})

// 格式化日期
const formatDate = (dateString: string) => {
  if (!dateString) return ''
  return new Date(dateString).toLocaleString('zh-CN')
}

// 获取表格数据
const getTableData = async () => {
  loading.value = true
  try {
    const params = {
      page: currentPage.value,
      page_size: pageSize.value,
      name: searchForm.name || undefined,
      username: searchForm.username || undefined,
      start_time: searchForm.createTime?.[0],
      end_time: searchForm.createTime?.[1]
    }
    
    const res = await credentialApi.getList(params)
    tableData.value = res.data?.list || []
    total.value = res.data?.pagination?.total || 0
  } catch (error) {
    console.error('获取凭据列表失败:', error)
    ElMessage.error('获取凭据列表失败')
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  currentPage.value = 1
  getTableData()
}

// 重置搜索
const handleReset = () => {
  searchForm.name = ''
  searchForm.username = ''
  searchForm.createTime = []
  currentPage.value = 1
  getTableData()
}

// 分页变化
const handleSizeChange = (val: number) => {
  pageSize.value = val
  currentPage.value = 1
  getTableData()
}

const handleCurrentChange = (val: number) => {
  currentPage.value = val
  getTableData()
}

// 新建凭据
const handleCreate = () => {
  dialogTitle.value = '新建凭据'
  resetForm()
  dialogVisible.value = true
}

// 编辑凭据
const handleEdit = (row: any) => {
  dialogTitle.value = '编辑凭据'
  Object.assign(formData, {
    id: row.id,
    name: row.name,
    username: row.username,
    password: '', // 编辑时不显示原密码
    description: row.description
  })
  dialogVisible.value = true
}

// 删除凭据
const handleDelete = (row: any) => {
  deleteItem.value = row
  deleteDialogVisible.value = true
}

// 确认删除
const confirmDelete = async () => {
  if (!deleteItem.value) return
  
  deleteLoading.value = true
  try {
    await credentialApi.delete(deleteItem.value.id)
    ElMessage.success('删除成功')
    deleteDialogVisible.value = false
    deleteItem.value = null
    getTableData()
  } catch (error) {
    console.error('删除凭据失败:', error)
    ElMessage.error('删除凭据失败')
  } finally {
    deleteLoading.value = false
  }
}

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    
    submitLoading.value = true
    try {
      if (formData.id) {
        // 编辑
        await credentialApi.update(formData.id, {
          name: formData.name,
          username: formData.username,
          password: formData.password,
          description: formData.description
        })
        ElMessage.success('编辑成功')
      } else {
        // 新建
        await credentialApi.create({
          name: formData.name,
          username: formData.username,
          password: formData.password,
          description: formData.description
        })
        ElMessage.success('创建成功')
      }
      
      dialogVisible.value = false
      getTableData()
    } catch (error) {
      console.error('操作失败:', error)
      ElMessage.error('操作失败')
    } finally {
      submitLoading.value = false
    }
  })
}

// 重置表单
const resetForm = () => {
  Object.assign(formData, {
    id: undefined,
    name: '',
    username: '',
    password: '',
    description: ''
  })
  
  // 清除表单验证
  if (formRef.value) {
    formRef.value.clearValidate()
  }
}

// 关闭对话框
const handleDialogClose = () => {
  resetForm()
}

// 初始化
onMounted(() => {
  getTableData()
})
</script>

<style scoped>
.credential-manage {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.search-form {
  margin-bottom: 20px;
  padding: 20px;
  background-color: #f5f7fa;
  border-radius: 4px;
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}
</style>