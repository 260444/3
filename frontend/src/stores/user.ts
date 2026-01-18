// stores/user.ts
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { login, getUserInfo, logout as logoutApi } from '@/api/user'

export const useUserStore = defineStore('user', () => {
  const userInfo = ref<any>(null)
  const token = ref<string | null>(localStorage.getItem('token'))
  const isAuthenticated = ref<boolean>(!!token.value)
  const permissions = ref<any[]>([]) // 用户权限列表
  const menus = ref<any[]>([]) // 用户菜单列表

  const loginAction = async (username: string, password: string) => {
    try {
      const response = await login(username, password)
      // 后端返回格式: { message: string, data: { user: {...}, token: string } }
      token.value = response.data.token
      userInfo.value = response.data.user
      isAuthenticated.value = true
      localStorage.setItem('token', token.value)
      return { success: true, data: response.data }
    } catch (error: any) {
      return { success: false, error: error.message }
    }
  }

  const logout = async () => {
    try {
      await logoutApi()
    } catch (error) {
      console.error('退出登录请求失败', error)
    } finally {
      token.value = null
      userInfo.value = null
      isAuthenticated.value = false
      localStorage.removeItem('token')
    }
  }

  const checkLoginStatus = () => {
    const storedToken = localStorage.getItem('token')
    if (storedToken) {
      token.value = storedToken
      isAuthenticated.value = true
      // 可以选择获取用户信息
      // fetchUserInfo()
    }
  }

  const fetchUserInfo = async () => {
    try {
      const response = await getUserInfo()
      // 后端返回格式: { message: string, data: {...} }
      userInfo.value = response.data
      return { success: true, data: response.data }
    } catch (error) {
      console.error('获取用户信息失败', error)
      return { success: false, error }
    }
  }

  // 设置用户权限
  const setPermissions = (perms: any[]) => {
    permissions.value = perms
  }

  // 设置用户菜单
  const setMenus = (menuList: any[]) => {
    menus.value = menuList
  }

  // 检查是否有权限
  const hasPermission = (permission: string) => {
    return permissions.value.includes(permission)
  }

  // 检查是否有任意权限
  const hasAnyPermission = (perms: string[]) => {
    return perms.some(perm => permissions.value.includes(perm))
  }

  // 检查是否有所有权限
  const hasAllPermissions = (perms: string[]) => {
    return perms.every(perm => permissions.value.includes(perm))
  }

  return {
    userInfo,
    token,
    isAuthenticated,
    permissions,
    menus,
    loginAction,
    logout,
    checkLoginStatus,
    fetchUserInfo,
    setPermissions,
    setMenus,
    hasPermission,
    hasAnyPermission,
    hasAllPermissions
  }
})