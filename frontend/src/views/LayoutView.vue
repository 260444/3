<template>
  <el-container class="layout-container">
    <!-- 侧边栏 -->
    <el-aside width="200px" class="aside">
      <div class="logo">后台管理系统</div>
      <el-menu
        :default-active="$route.path"
        :router="true"
        class="menu"
        :unique-opened="true"
      >
        <template v-for="menu in userStore.menus" :key="menu.id">
          <!-- 有子菜单的情况 -->
          <el-sub-menu v-if="menu.children && menu.children.length > 0" :index="menu.path">
            <template #title>
              <el-icon v-if="menu.icon">
                <component :is="menu.icon" />
              </el-icon>
              <span>{{ menu.title }}</span>
            </template>
            <el-menu-item
              v-for="child in menu.children"
              :key="child.id"
              :index="child.path"
              :route="{ name: child.name }"
            >
              <el-icon v-if="child.icon">
                <component :is="child.icon" />
              </el-icon>
              <span>{{ child.title }}</span>
            </el-menu-item>
          </el-sub-menu>
          <!-- 没有子菜单的情况 -->
          <el-menu-item v-else :index="menu.path" :route="{ name: menu.name }">
            <el-icon v-if="menu.icon">
              <component :is="menu.icon" />
            </el-icon>
            <span>{{ menu.title }}</span>
          </el-menu-item>
        </template>
      </el-menu>
    </el-aside>

    <!-- 主内容区域 -->
    <el-container>
      <!-- 顶部导航 -->
      <el-header class="header">
        <div class="header-left">
          <el-icon class="trigger" @click="toggleCollapse">
            <component :is="isCollapse ? 'Expand' : 'Fold'" />
          </el-icon>
        </div>
        <div class="header-right">
          <el-dropdown v-if="userInfo">
            <span class="el-dropdown-link">
              {{ userInfo.nickname || userInfo.username || '管理员' }}
              <el-icon class="el-icon--right">
                <ArrowDown />
              </el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="showProfile">个人资料</el-dropdown-item>
                <el-dropdown-item @click="changePassword">修改密码</el-dropdown-item>
                <el-dropdown-item @click="handleLogout" divided>退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
          <span v-else class="user-name">管理员</span>
        </div>
      </el-header>

      <!-- 页面内容 -->
      <el-main class="main">
        <router-view />
      </el-main>
    </el-container>

    <!-- 修改密码对话框 -->
    <el-dialog
      v-model="passwordDialogVisible"
      title="修改密码"
      width="500px"
      :close-on-click-modal="false"
    >
      <el-form
        ref="passwordFormRef"
        :model="passwordForm"
        :rules="passwordRules"
        label-width="100px"
      >
        <el-form-item label="旧密码" prop="oldPassword">
          <el-input
            v-model="passwordForm.oldPassword"
            type="password"
            placeholder="请输入旧密码"
            show-password
          />
        </el-form-item>
        <el-form-item label="新密码" prop="newPassword">
          <el-input
            v-model="passwordForm.newPassword"
            type="password"
            placeholder="请输入新密码"
            show-password
          />
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input
            v-model="passwordForm.confirmPassword"
            type="password"
            placeholder="请再次输入新密码"
            show-password
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="handlePasswordCancel">取消</el-button>
          <el-button type="primary" @click="handlePasswordSubmit">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </el-container>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  House,
  User,
  Avatar,
  Menu,
  Document,
  Lock,
  Fold,
  Expand,
  ArrowDown
} from '@element-plus/icons-vue'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()

const isCollapse = ref(false)

const userInfo = computed(() => userStore.userInfo)

const toggleCollapse = () => {
  isCollapse.value = !isCollapse.value
}

const handleLogout = () => {
  ElMessageBox.confirm('确定要退出登录吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    userStore.logout()
    router.push('/login')
    ElMessage.success('已退出登录')
  })
}

const showProfile = () => {
  // 显示个人资料弹窗
  console.log('显示个人资料')
}

const passwordDialogVisible = ref(false)
const passwordForm = ref({
  oldPassword: '',
  newPassword: '',
  confirmPassword: ''
})
const passwordFormRef = ref()
const passwordRules = {
  oldPassword: [
    { required: true, message: '请输入旧密码', trigger: 'blur' },
    { min: 6, max: 20, message: '密码长度在 6 到 20 个字符', trigger: 'blur' }
  ],
  newPassword: [
    { required: true, message: '请输入新密码', trigger: 'blur' },
    { min: 6, max: 20, message: '密码长度在 6 到 20 个字符', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请再次输入新密码', trigger: 'blur' },
    {
      validator: (rule: any, value: string, callback: any) => {
        if (value !== passwordForm.value.newPassword) {
          callback(new Error('两次输入的密码不一致'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
}

const changePassword = () => {
  passwordDialogVisible.value = true
  passwordForm.value = {
    oldPassword: '',
    newPassword: '',
    confirmPassword: ''
  }
}

const handlePasswordSubmit = async () => {
  await passwordFormRef.value.validate()
  
  try {
    const { changePassword: changePasswordApi } = await import('@/api/user')
    await changePasswordApi(passwordForm.value.oldPassword, passwordForm.value.newPassword)
    
    ElMessage.success('密码修改成功，请重新登录')
    passwordDialogVisible.value = false
    userStore.logout()
    router.push('/login')
  } catch (error: any) {
    ElMessage.error(error.message || '密码修改失败')
  }
}

const handlePasswordCancel = () => {
  passwordDialogVisible.value = false
  passwordForm.value = {
    oldPassword: '',
    newPassword: '',
    confirmPassword: ''
  }
}
</script>

<style scoped>
.layout-container {
  height: 100vh;
}

.aside {
  background-color: #191a23;
  color: #ffffff;
}

.logo {
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  font-weight: bold;
  background-color: #1f2029;
  color: #ffffff;
}

.menu {
  border: none;
  background-color: #191a23;
}

.menu :deep(.el-menu-item) {
  color: #e5e7eb;
  background-color: #191a23;
}

.menu :deep(.el-menu-item:hover) {
  background-color: #2c2e3a !important;
  color: #ffffff !important;
}

.menu :deep(.el-menu-item.is-active) {
  background-color: #409EFF !important;
  color: #ffffff !important;
}

.menu :deep(.el-sub-menu__title) {
  color: #e5e7eb;
  background-color: #191a23;
}

.menu :deep(.el-sub-menu__title:hover) {
  background-color: #2c2e3a !important;
  color: #ffffff !important;
}

.menu :deep(.el-sub-menu.is-active > .el-sub-menu__title) {
  color: #409EFF !important;
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  background-color: #fff;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
}

.header-left {
  display: flex;
  align-items: center;
}

.trigger {
  cursor: pointer;
  margin-right: 20px;
  font-size: 18px;
}

.header-right {
  display: flex;
  align-items: center;
}

.el-dropdown-link {
  cursor: pointer;
  display: flex;
  align-items: center;
}

.user-name {
  display: flex;
  align-items: center;
}

.main {
  background-color: #f0f2f5;
  padding: 20px;
}
</style>