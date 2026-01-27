<template>
  <div class="user-manage-container">
    <el-card>
      <!-- 搜索和操作区域 -->
      <div class="header-section">
        <el-form :inline="true" :model="searchForm">
          <el-form-item label="用户名">
            <el-input v-model="searchForm.username" placeholder="请输入用户名" />
          </el-form-item>
          <el-form-item label="邮箱">
            <el-input v-model="searchForm.email" placeholder="请输入邮箱" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleSearch">搜索</el-button>
            <el-button @click="resetSearch">重置</el-button>
          </el-form-item>
        </el-form>
        <el-button type="primary" @click="handleAdd">新增用户</el-button>
      </div>

      <!-- 用户列表 -->
      <el-table :data="userList" stripe style="width: 100%" v-loading="loading">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="username" label="用户名" />
        <el-table-column prop="nickname" label="昵称" />
        <el-table-column prop="email" label="邮箱" />
        <el-table-column prop="phone" label="手机号" />
        <el-table-column label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '正常' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="last_login_at" label="最后登录时间" />
        <el-table-column label="操作" width="420">
          <template #default="{ row }">
            <el-button size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button size="small" type="warning" @click="handleResetPassword(row)">重置密码</el-button>
            <el-button size="small" type="primary" @click="handleAssignRoles(row)">分配角色</el-button>
            <el-button 
              size="small" 
              :type="row.status === 1 ? 'danger' : 'success'" 
              @click="toggleStatus(row)"
            >
              {{ row.status === 1 ? '禁用' : '启用' }}
            </el-button>
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

    <!-- 用户编辑弹窗 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form :model="userForm" :rules="userRules" ref="userFormRef" label-width="80px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="userForm.username" :disabled="!!userForm.id" />
        </el-form-item>
        <el-form-item label="昵称" prop="nickname">
          <el-input v-model="userForm.nickname" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="userForm.email" />
        </el-form-item>
        <el-form-item label="手机号" prop="phone">
          <el-input v-model="userForm.phone" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="userForm.status">
            <el-radio :label="1">正常</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item v-if="!userForm.id" label="密码" prop="password">
          <el-input v-model="userForm.password" type="password" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmUser">确定</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 重置密码弹窗 -->
    <el-dialog
      v-model="resetPasswordVisible"
      title="重置密码"
      width="500px"
      :close-on-click-modal="false"
    >
      <el-form
        ref="resetPasswordFormRef"
        :model="resetPasswordForm"
        :rules="resetPasswordRules"
        label-width="100px"
      >
        <el-form-item label="用户名">
          <el-input v-model="resetPasswordForm.username" disabled />
        </el-form-item>
        <el-form-item label="新密码" prop="newPassword">
          <el-input
            v-model="resetPasswordForm.newPassword"
            type="password"
            placeholder="请输入新密码"
            show-password
          />
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input
            v-model="resetPasswordForm.confirmPassword"
            type="password"
            placeholder="请再次输入新密码"
            show-password
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetPasswordVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmResetPassword">确定</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 分配角色弹窗 -->
    <el-dialog v-model="roleDialogVisible" :title="`为【${currentUsername}】分配角色`" width="500px">
      <el-radio-group v-model="selectedRoleIdent">
        <el-radio 
          v-for="role in roleList" 
          :key="role.id" 
          :label="role.ident"
        >
          {{ role.name }} ({{ role.ident }})
        </el-radio>
      </el-radio-group>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="roleDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmAssignRoles">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getUsers, createUser, updateUser, deleteUser, updateUserStatus, resetPassword, getUserRoles, assignRole, removeRole } from '@/api/user'
import { getRoles } from '@/api/role'

// 搜索表单
const searchForm = reactive({
  username: '',
  email: ''
})

// 分页信息
const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 0
})

// 用户列表
const userList = ref<any[]>([])
const loading = ref(false)

// 弹窗相关
const dialogVisible = ref(false)
const dialogTitle = ref('')
const userFormRef = ref()
const userForm = ref<any>({
  id: undefined,
  username: '',
  nickname: '',
  email: '',
  phone: '',
  status: 1,
  password: ''
})

// 角色列表
const roleList = ref<any[]>([])

// 分配角色相关
const roleDialogVisible = ref(false)
const currentUserId = ref(0)
const currentUsername = ref('')
const selectedRoleIdent = ref<string>('')
const originalRoleIdent = ref<string>('')

// 重置密码弹窗相关
const resetPasswordVisible = ref(false)
const resetPasswordFormRef = ref()
const resetPasswordForm = ref({
  userId: 0,
  username: '',
  newPassword: '',
  confirmPassword: ''
})

// 表单验证规则
const userRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 20, message: '长度在 3 到 20 个字符', trigger: 'blur' }
  ],
  nickname: [
    { required: true, message: '请输入昵称', trigger: 'blur' }
  ],
  email: [
    { required: true, message: '请输入邮箱地址', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱地址', trigger: 'blur' }
  ],
  password: [
    { required: !userForm.value.id, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }
  ]
}

const resetPasswordRules = {
  newPassword: [
    { required: true, message: '请输入新密码', trigger: 'blur' },
    { min: 6, max: 20, message: '密码长度在 6 到 20 个字符', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请再次输入新密码', trigger: 'blur' },
    {
      validator: (rule: any, value: string, callback: any) => {
        if (value !== resetPasswordForm.value.newPassword) {
          callback(new Error('两次输入的密码不一致'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
}

// 获取用户列表
const getUserList = async () => {
  loading.value = true
  try {
    const params = {
      page: pagination.page,
      page_size: pagination.page_size,
      ...searchForm
    }
    const response = await getUsers(params)
    userList.value = response.data.list
    pagination.total = response.data.total
  } catch (error) {
    console.error('获取用户列表失败:', error)
    ElMessage.error('获取用户列表失败')
  } finally {
    loading.value = false
  }
}

// 获取角色列表
const getRoleList = async () => {
  try {
    const response = await getRoles({ page: 1, page_size: 100 })
    roleList.value = response.data.list
  } catch (error) {
    console.error('获取角色列表失败:', error)
    ElMessage.error('获取角色列表失败')
  }
}

// 搜索
const handleSearch = () => {
  pagination.page = 1
  getUserList()
}

// 重置搜索
const resetSearch = () => {
  searchForm.username = ''
  searchForm.email = ''
  pagination.page = 1
  getUserList()
}

// 处理分页大小变化
const handleSizeChange = (size: number) => {
  pagination.page_size = size
  pagination.page = 1
  getUserList()
}

// 处理当前页变化
const handleCurrentChange = (page: number) => {
  pagination.page = page
  getUserList()
}

// 处理添加用户
const handleAdd = () => {
  userForm.value = {
    id: undefined,
    username: '',
    nickname: '',
    email: '',
    phone: '',
    status: 1,
    password: ''
  }
  dialogTitle.value = '新增用户'
  dialogVisible.value = true
}

// 处理编辑用户
const handleEdit = (row: any) => {
  userForm.value = { ...row }
  userForm.value.password = '' // 编辑时不显示原密码
  dialogTitle.value = '编辑用户'
  dialogVisible.value = true
}

// 确认用户操作（新增或编辑）
const confirmUser = async () => {
  try {
    await userFormRef.value.validate()
    
    if (userForm.value.id) {
      // 编辑用户
      await updateUser(userForm.value.id, userForm.value)
      ElMessage.success('用户更新成功')
    } else {
      // 新增用户
      await createUser(userForm.value)
      ElMessage.success('用户创建成功')
    }
    
    dialogVisible.value = false
    getUserList()
  } catch (error) {
    console.error('操作用户失败:', error)
    ElMessage.error('操作失败')
  }
}

// 处理删除用户
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm(`确定要删除用户 "${row.username}" 吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    
    await deleteUser(row.id)
    ElMessage.success('删除成功')
    getUserList()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除用户失败:', error)
      ElMessage.error('删除失败')
    }
  }
}

// 处理分配角色
const handleAssignRoles = async (row: any) => {
  currentUserId.value = row.id
  currentUsername.value = row.username
  roleDialogVisible.value = true
  
  try {
    const response = await getUserRoles(row.id)
    const roles = response.data || []
    // 既然改为单选，我们只取第一个角色（如果存在），或者空
    selectedRoleIdent.value = roles.length > 0 ? roles[0] : ''
    originalRoleIdent.value = selectedRoleIdent.value
  } catch (error) {
    console.error('获取用户角色失败:', error)
    ElMessage.error('获取用户角色失败')
  }
}

// 确认分配角色
const confirmAssignRoles = async () => {
  try {
    // 如果没有变化，直接关闭
    if (selectedRoleIdent.value === originalRoleIdent.value) {
      roleDialogVisible.value = false
      return
    }

    // 如果原先有角色，先移除
    if (originalRoleIdent.value) {
      await removeRole(currentUserId.value, originalRoleIdent.value)
    }

    // 如果新选了角色，则添加
    if (selectedRoleIdent.value) {
      await assignRole(currentUserId.value, selectedRoleIdent.value)
    }

    ElMessage.success('角色分配成功')
    roleDialogVisible.value = false
  } catch (error) {
    console.error('分配角色失败:', error)
    ElMessage.error('分配角色失败')
  }
}

// 切换用户状态
const toggleStatus = async (row: any) => {
  try {
    const newStatus = row.status === 1 ? 0 : 1
    await updateUserStatus(row.id, newStatus)
    ElMessage.success(newStatus === 1 ? '用户已启用' : '用户已禁用')
    getUserList()
  } catch (error) {
    console.error('切换状态失败:', error)
    ElMessage.error('操作失败')
  }
}

// 处理重置密码
const handleResetPassword = (row: any) => {
  resetPasswordForm.value = {
    userId: row.id,
    username: row.username,
    newPassword: '',
    confirmPassword: ''
  }
  resetPasswordVisible.value = true
}

// 确认重置密码
const confirmResetPassword = async () => {
  try {
    await resetPasswordFormRef.value.validate()
    
    await resetPassword(resetPasswordForm.value.userId, resetPasswordForm.value.newPassword)
    
    ElMessage.success('密码重置成功')
    resetPasswordVisible.value = false
  } catch (error) {
    console.error('重置密码失败:', error)
    ElMessage.error('重置密码失败')
  }
}

onMounted(() => {
  getUserList()
  getRoleList()
})
</script>

<style scoped>
.user-manage-container {
  background-color: #fff;
  padding: 20px;
  border-radius: 4px;
}

.header-section {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
  align-items: flex-start;
}

.header-section .el-form {
  flex: 1;
  margin-right: 20px;
}

.header-section .el-form-item {
  margin-right: 10px;
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>