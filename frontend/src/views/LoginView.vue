<template>
  <div class="login-container">
    <el-card class="login-card">
      <h2 class="login-title">后台管理系统</h2>
      <el-form 
        :model="loginForm" 
        :rules="loginRules" 
        ref="loginFormRef"
        @keyup.enter="handleLogin"
      >
        <el-form-item prop="username">
          <el-input 
            v-model="loginForm.username" 
            placeholder="请输入用户名"
            prefix-icon="User"
            clearable
          />
        </el-form-item>
        <el-form-item prop="password">
          <el-input 
            v-model="loginForm.password" 
            :type="showPassword ? 'text' : 'password'" 
            placeholder="请输入密码"
            prefix-icon="Lock"
            :show-password="true"
          />
        </el-form-item>
        <el-form-item>
          <div class="login-options">
            <el-checkbox v-model="rememberMe">记住密码</el-checkbox>
            <el-tooltip content="在公共设备上请勿勾选" placement="top" effect="light">
              <el-icon><Warning /></el-icon>
            </el-tooltip>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button 
            type="primary" 
            @click="handleLogin" 
            :loading="loading" 
            style="width: 100%;"
          >
            登录
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { Warning } from '@element-plus/icons-vue'
import { useUserStore } from '@/stores/user'

interface LoginForm {
  username: string
  password: string
}

const router = useRouter()
const userStore = useUserStore()

const loginForm = reactive<LoginForm>({
  username: '',
  password: ''
})

const loginRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }
  ]
}

const loading = ref(false)
const showPassword = ref(false)
const rememberMe = ref(false)

// 从localStorage加载记住的用户名和密码
onMounted(() => {
  const savedUsername = localStorage.getItem('remembered_username')
  const savedPassword = localStorage.getItem('remembered_password')
  const savedRememberMe = localStorage.getItem('remember_me') === 'true'
  
  if (savedUsername) {
    loginForm.username = savedUsername
  }
  if (savedPassword && savedRememberMe) {
    loginForm.password = savedPassword
    rememberMe.value = savedRememberMe
  }
})

const handleLogin = async () => {
  // 简单验证
  if (!loginForm.username || !loginForm.password) {
    ElMessage.error('请输入用户名和密码')
    return
  }

  loading.value = true
  try {
    const result = await userStore.loginAction(loginForm.username, loginForm.password)
    if (result.success) {
      // 如果勾选了记住密码，则保存用户名和密码到localStorage
      if (rememberMe.value) {
        localStorage.setItem('remembered_username', loginForm.username)
        localStorage.setItem('remembered_password', loginForm.password)
        localStorage.setItem('remember_me', 'true')
      } else {
        // 如果没有勾选，则清除保存的信息
        localStorage.removeItem('remembered_username')
        localStorage.removeItem('remembered_password')
        localStorage.removeItem('remember_me')
      }
      
      ElMessage.success('登录成功')
      router.push('/')
    } else {
      ElMessage.error(result.error || '登录失败')
    }
  } catch (error) {
    console.error('登录错误:', error)
    ElMessage.error('登录失败')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f0f2f5;
}

.login-card {
  width: 400px;
  padding: 20px;
}

.login-title {
  text-align: center;
  margin-bottom: 30px;
}

.login-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.el-icon {
  color: #e6a23c;
  cursor: pointer;
}
</style>