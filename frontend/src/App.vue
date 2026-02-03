<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()

onMounted(async () => {
  // 检查用户登录状态
  userStore.checkLoginStatus()
  // 如果已登录，获取用户信息
  if (userStore.isAuthenticated) {
    try {
      await userStore.fetchUserInfo()
    } catch (error) {
      // 如果获取用户信息失败，不进行处理，避免循环请求
      console.error('获取用户信息失败，可能没有权限:', error)
    }
  }
})
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>