<template>
  <div class="permission-resource">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>权限资源管理</span>
        </div>
      </template>

      <!-- 搜索和操作区域 -->
      <div class="search-bar">
        <el-form :inline="true" :model="searchForm">
          <el-form-item label="路径">
            <el-input v-model="searchForm.path" placeholder="请输入路径" clearable />
          </el-form-item>
          <el-form-item label="方法" style="min-width: 220px;">
            <el-select v-model="searchForm.method" placeholder="请选择方法" clearable>
              <el-option label="GET" value="GET" />
              <el-option label="POST" value="POST" />
              <el-option label="PUT" value="PUT" />
              <el-option label="DELETE" value="DELETE" />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleSearch">查询</el-button>
            <el-button @click="handleReset">重置</el-button>
            <el-button type="primary" @click="handleAdd">新增权限</el-button>
          </el-form-item>
        </el-form>
      </div>

      <!-- 权限列表 -->
      <el-table :data="permissionList" border style="width: 100%" v-loading="loading">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="path" label="路径" />
        <el-table-column prop="method" label="方法" width="100">
          <template #default="{ row }">
            <el-tag :type="getMethodType(row.method)">{{ row.method }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '正常' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button type="primary" size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button 
              :type="row.status === 1 ? 'warning' : 'success'" 
              size="small" 
              @click="handleToggleStatus(row)"
            >
              {{ row.status === 1 ? '禁用' : '启用' }}
            </el-button>
            <el-button type="danger" size="small" @click="handleDelete(row)">删除</el-button>
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

    <!-- 权限编辑对话框 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form :model="permissionForm" :rules="permissionRules" ref="permissionFormRef" label-width="100px">
        <el-form-item label="路径" prop="path">
          <el-input v-model="permissionForm.path" placeholder="请输入权限路径，如：/api/v1/users" />
        </el-form-item>
        <el-form-item label="方法" prop="method">
          <el-select v-model="permissionForm.method" placeholder="请选择请求方法">
            <el-option label="GET" value="GET" />
            <el-option label="POST" value="POST" />
            <el-option label="PUT" value="PUT" />
            <el-option label="DELETE" value="DELETE" />
            <el-option label="PATCH" value="PATCH" />
            <el-option label="OPTIONS" value="OPTIONS" />
          </el-select>
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input 
            v-model="permissionForm.description" 
            type="textarea" 
            placeholder="请输入权限描述" 
            :rows="3"
          />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-switch
            v-model="permissionForm.status"
            :active-value="1"
            :inactive-value="0"
            active-text="正常"
            inactive-text="禁用"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleConfirm">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import * as permissionApi from '@/api/permission'

// 权限列表
const permissionList = ref<any[]>([])
// 总数
const total = ref(0)
// 当前页
const currentPage = ref(1)
// 每页大小
const pageSize = ref(10)
// 加载状态
const loading = ref(false)

// 搜索表单
const searchForm = ref({
  path: '',
  method: ''
})

// 对话框相关
const dialogVisible = ref(false)
const dialogTitle = ref('新增权限')
const permissionFormRef = ref<FormInstance>()
const permissionForm = ref({
  id: 0,
  path: '',
  method: '',
  description: '',
  status: 1
})

// 表单验证规则
const permissionRules = ref<FormRules>({
  path: [
    { required: true, message: '请输入权限路径', trigger: 'blur' },
    { min: 1, max: 255, message: '路径长度在1到255个字符之间', trigger: 'blur' }
  ],
  method: [
    { required: true, message: '请选择请求方法', trigger: 'change' }
  ]
})

// 获取权限列表
const fetchPermissions = async () => {
  loading.value = true
  try {
    const params = {
      page: currentPage.value,
      page_size: pageSize.value,
      path: searchForm.value.path || undefined,
      method: searchForm.value.method || undefined
    }
    const response: any = await permissionApi.getPermissions(params)
    permissionList.value = response.data.list
    total.value = response.data.total
  } catch (error) {
    ElMessage.error('获取权限列表失败')
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  currentPage.value = 1
  fetchPermissions()
}

// 重置
const handleReset = () => {
  searchForm.value = {
    path: '',
    method: ''
  }
  currentPage.value = 1
  fetchPermissions()
}

// 分页大小改变
const handleSizeChange = (size: number) => {
  pageSize.value = size
  fetchPermissions()
}

// 当前页改变
const handleCurrentChange = (page: number) => {
  currentPage.value = page
  fetchPermissions()
}

// 新增权限
const handleAdd = () => {
  dialogTitle.value = '新增权限'
  permissionForm.value = {
    id: 0,
    path: '',
    method: '',
    description: '',
    status: 1
  }
  dialogVisible.value = true
}

// 编辑权限
const handleEdit = (row: any) => {
  dialogTitle.value = '编辑权限'
  permissionForm.value = { ...row }
  dialogVisible.value = true
}

// 删除权限
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm(`确定要删除权限 "${row.path}" 吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    
    await permissionApi.deletePermission(row.id)
    ElMessage.success('删除成功')
    fetchPermissions()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

// 切换状态
const handleToggleStatus = async (row: any) => {
  try {
    const newStatus = row.status === 1 ? 0 : 1
    await permissionApi.updatePermissionStatus(row.id, newStatus)
    ElMessage.success(newStatus === 1 ? '启用成功' : '禁用成功')
    fetchPermissions()
  } catch (error) {
    ElMessage.error('更新状态失败')
  }
}

// 确认操作（新增/编辑）
const handleConfirm = async () => {
  try {
    await permissionFormRef.value?.validate()
    
    if (permissionForm.value.id > 0) {
      // 编辑
      await permissionApi.updatePermission(permissionForm.value.id, {
        path: permissionForm.value.path,
        method: permissionForm.value.method,
        description: permissionForm.value.description,
        status: permissionForm.value.status
      })
      ElMessage.success('更新成功')
    } else {
      // 新增
      await permissionApi.createPermission({
        path: permissionForm.value.path,
        method: permissionForm.value.method,
        description: permissionForm.value.description,
        status: permissionForm.value.status
      })
      ElMessage.success('创建成功')
    }
    
    dialogVisible.value = false
    fetchPermissions()
  } catch (error) {
    ElMessage.error('操作失败')
  }
}

// 获取方法类型
const getMethodType = (method: string) => {
  const typeMap: any = {
    GET: 'success',
    POST: 'primary',
    PUT: 'warning',
    DELETE: 'danger',
    PATCH: 'info',
    OPTIONS: 'info'
  }
  return typeMap[method] || 'info'
}

onMounted(() => {
  fetchPermissions()
})
</script>

<style scoped>
.permission-resource {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.search-bar {
  margin-bottom: 20px;
}

.pagination {
  margin-top: 20px;
  text-align: right;
}

.dialog-footer {
  text-align: right;
}
</style>