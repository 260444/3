<template>
  <div class="role-manage-container">
    <el-card>
      <!-- 操作区域 -->
      <div class="header-section">
        <el-button type="primary" @click="handleAdd">新增角色</el-button>
      </div>

      <!-- 角色列表 -->
      <el-table :data="roleList" stripe style="width: 100%" v-loading="loading">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="角色名称" />
        <el-table-column prop="description" label="描述" />
        <el-table-column label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '正常' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="300">
          <template #default="{ row }">
            <el-button size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button size="small" type="primary" @click="handlePermission(row)">权限</el-button>
            <el-button size="small" type="danger" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="pagination">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.page_size"
          :page-sizes="[10, 20, 50, 100]"
          :total="pagination.total"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>

    <!-- 角色编辑弹窗 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form :model="roleForm" :rules="roleRules" ref="roleFormRef" label-width="80px">
        <el-form-item label="角色名称" prop="name">
          <el-input v-model="roleForm.name" placeholder="请输入角色名称" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="roleForm.description" type="textarea" placeholder="请输入角色描述" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="roleForm.status">
            <el-radio :label="1">正常</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmRole">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getRoles, createRole, updateRole, deleteRole } from '@/api/role'

const router = useRouter()

// 分页信息
const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 0
})

// 角色列表
const roleList = ref<any[]>([])
const loading = ref(false)

// 弹窗相关
const dialogVisible = ref(false)
const dialogTitle = ref('')
const roleFormRef = ref()
const roleForm = ref<any>({
  id: undefined,
  name: '',
  description: '',
  status: 1
})

// 表单验证规则
const roleRules = {
  name: [
    { required: true, message: '请输入角色名称', trigger: 'blur' },
    { min: 2, max: 20, message: '长度在 2 到 20 个字符', trigger: 'blur' }
  ],
  description: [
    { required: true, message: '请输入角色描述', trigger: 'blur' }
  ]
}

// 获取角色列表
const getRoleList = async () => {
  loading.value = true
  try {
    const params = {
      page: pagination.page,
      page_size: pagination.page_size
    }
    const response = await getRoles(params)
    roleList.value = response.data.list
    pagination.total = response.data.total
  } catch (error) {
    console.error('获取角色列表失败:', error)
    ElMessage.error('获取角色列表失败')
  } finally {
    loading.value = false
  }
}

// 处理分页大小变化
const handleSizeChange = (size: number) => {
  pagination.page_size = size
  pagination.page = 1
  getRoleList()
}

// 处理当前页变化
const handleCurrentChange = (page: number) => {
  pagination.page = page
  getRoleList()
}

// 处理添加角色
const handleAdd = () => {
  roleForm.value = {
    id: undefined,
    name: '',
    description: '',
    status: 1
  }
  dialogTitle.value = '新增角色'
  dialogVisible.value = true
}

// 处理编辑角色
const handleEdit = (row: any) => {
  roleForm.value = { ...row }
  dialogTitle.value = '编辑角色'
  dialogVisible.value = true
}

// 确认角色操作（新增或编辑）
const confirmRole = async () => {
  try {
    await roleFormRef.value.validate()
    
    if (roleForm.value.id) {
      // 编辑角色
      await updateRole(roleForm.value.id, roleForm.value)
      ElMessage.success('角色更新成功')
    } else {
      // 新增角色
      await createRole(roleForm.value)
      ElMessage.success('角色创建成功')
    }
    
    dialogVisible.value = false
    getRoleList()
  } catch (error) {
    console.error('操作角色失败:', error)
    ElMessage.error('操作失败')
  }
}

// 处理删除角色
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm(`确定要删除角色 "${row.name}" 吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })

    await deleteRole(row.id)
    ElMessage.success('删除成功')
    getRoleList()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除角色失败:', error)
      ElMessage.error('删除失败')
    }
  }
}

// 处理权限管理
const handlePermission = (row: any) => {
  router.push({
    name: 'permissions',
    query: { roleId: row.id, roleName: row.name }
  })
}

onMounted(() => {
  getRoleList()
})
</script>

<style scoped>
.role-manage-container {
  background-color: #fff;
  padding: 20px;
  border-radius: 4px;
}

.header-section {
  margin-bottom: 20px;
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>