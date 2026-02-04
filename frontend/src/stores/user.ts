// stores/user.ts
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { login, getUserInfo, logout as logoutApi } from '@/api/user'

export const useUserStore = defineStore('user', () => {
  // 记住密码相关
  const rememberPassword = ref<boolean>(localStorage.getItem('rememberPassword') === 'true')
  const savedUsername = ref<string>(localStorage.getItem('savedUsername') || '')
  const savedPassword = ref<string>(localStorage.getItem('savedPassword') || '')

  const userInfo = ref<any>(null)
  const token = ref<string | null>(localStorage.getItem('token'))
  const isAuthenticated = ref<boolean>(!!token.value)
  const permissions = ref<any[]>([]) // 用户权限列表
  const menus = ref<any[]>([]) // 用户菜单列表

  // 兼容性响应处理函数
  const handleResponse = (response: any) => {
    // 支持新格式 { success: boolean, message: string, data: any }
    // 和旧格式 { message: string, data: any }
    if (response.hasOwnProperty('success')) {
      return {
        success: response.success,
        data: response.data,
        error: response.error || response.message
      };
    } else {
      // 旧格式兼容
      return {
        success: true,
        data: response.data,
        error: null
      };
    }
  };

  const loginAction = async (username: string, password: string) => {
    try {
      const response = await login(username, password);
      const result = handleResponse(response);
      
      if (result.success) {
        token.value = result.data.token;
        userInfo.value = result.data.user;
        isAuthenticated.value = true;
        localStorage.setItem('token', token.value!);
        
        // 保存用户名和密码（如果用户选择记住密码）
        if (rememberPassword.value) {
          localStorage.setItem('rememberPassword', 'true');
          localStorage.setItem('savedUsername', username);
          localStorage.setItem('savedPassword', password);
        } else {
          localStorage.removeItem('rememberPassword');
          localStorage.removeItem('savedUsername');
          localStorage.removeItem('savedPassword');
        }
        
        return { success: true, data: result.data };
      } else {
        return { success: false, error: result.error };
      }
    } catch (error: any) {
      return { success: false, error: error.message };
    }
  }

  const logout = async () => {
    try {
      await logoutApi();
    } catch (error) {
      console.error('退出登录请求失败', error);
    } finally {
      token.value = null;
      userInfo.value = null;
      isAuthenticated.value = false;
      localStorage.removeItem('token');
    }
  }

  const checkLoginStatus = () => {
    const storedToken = localStorage.getItem('token');
    if (storedToken) {
      token.value = storedToken;
      isAuthenticated.value = true;
      // 可以选择获取用户信息
      // fetchUserInfo()
    }
    
    // 检查是否记住密码
    const storedRemember = localStorage.getItem('rememberPassword');
    if (storedRemember) {
      rememberPassword.value = storedRemember === 'true';
      savedUsername.value = localStorage.getItem('savedUsername') || '';
      savedPassword.value = localStorage.getItem('savedPassword') || '';
    }
  }

  const fetchUserInfo = async () => {
    try {
      const response = await getUserInfo();
      const result = handleResponse(response);
      
      if (result.success) {
        userInfo.value = result.data;
        return { success: true, data: result.data };
      } else {
        return { success: false, error: result.error };
      }
    } catch (error) {
      console.error('获取用户信息失败', error);
      return { success: false, error: (error as Error).message };
    }
  }

  // 设置用户权限
  const setPermissions = (perms: any[]) => {
    permissions.value = perms;
  }

  // 设置用户菜单
  const setMenus = (menuList: any[]) => {
    menus.value = menuList;
  }

  // 检查是否有权限
  const hasPermission = (permission: string) => {
    return permissions.value.includes(permission);
  }

  // 检查是否有任意权限
  const hasAnyPermission = (perms: string[]) => {
    return perms.some(perm => permissions.value.includes(perm));
  }

  // 检查是否有所有权限
  const hasAllPermissions = (perms: string[]) => {
    return perms.every(perm => permissions.value.includes(perm));
  }

  return {
    userInfo,
    token,
    isAuthenticated,
    permissions,
    menus,
    rememberPassword,
    savedUsername,
    savedPassword,
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