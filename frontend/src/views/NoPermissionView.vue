<template>
  <div class="no-permission-container">
    <div class="error-content" role="alert" aria-labelledby="error-title" aria-describedby="error-message">
      <div class="error-icon" aria-hidden="true">
        <el-icon><CircleClose /></el-icon>
      </div>
      <h1 id="error-title" class="error-title">无权限访问</h1>
      <p id="error-message" class="error-message">抱歉，您没有权限访问此页面</p>
      <div class="error-actions">
        <el-button type="primary" @click="goToConsole" size="default">返回控制台</el-button>
        <el-button @click="logout" size="default">退出登录</el-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ElIcon, ElButton } from 'element-plus'
import { CircleClose } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()

const goToConsole = () => {
  router.push('/dashboard')
}

const logout = async () => {
  await userStore.logout()
  router.push('/login')
}
</script>

<style scoped>
.no-permission-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f5f5f5;
  padding: 20px;
}

.error-content {
  text-align: center;
  background: white;
  padding: 40px;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  max-width: 400px;
  width: 100%;
}

.error-icon {
  font-size: 64px;
  color: #f56c6c;
  margin-bottom: 20px;
}

.error-title {
  font-size: 24px;
  color: #303133;
  margin-bottom: 10px;
}

.error-message {
  font-size: 16px;
  color: #606266;
  margin-bottom: 30px;
}

.error-actions {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.error-actions .el-button {
  width: 100%;
  margin: 0;
}
</style>