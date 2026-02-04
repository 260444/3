// api/request.ts
import axios, { AxiosRequestConfig, AxiosResponse } from 'axios'

// 创建axios实例
const request = axios.create({
  baseURL: '/api/v1', // 通过vite.config.ts中的代理配置
  timeout: 10000
})

// 请求拦截器
request.interceptors.request.use(
  (config: AxiosRequestConfig) => {
    // 从localStorage获取token
    const token = localStorage.getItem('token')
    if (token) {
      if (!config.headers) {
        config.headers = {}
      }
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    console.error('请求错误:', error)
    return Promise.reject(error)
  }
)

// 响应拦截器
request.interceptors.response.use(
  (response: AxiosResponse) => {
    // 新的统一响应格式: { success: boolean, message: string, data: any, error: string }
    const res = response.data;
    
    // 如果success为false，视为业务错误
    if (res.success === false) {
      return Promise.reject(new Error(res.error || res.message || '请求失败'));
    }
    
    // 成功响应，返回整个响应对象
    return res;
  },
  (error) => {
        if (error.response?.status === 401) {
          // 如果是401未授权，跳转到登录页
          localStorage.removeItem('token')
          window.location.href = '/login'
        } else if (error.response?.status === 403) {
          // 如果是403权限不足，跳转到无权限页面
          window.location.href = '/no-permission'
        }
    // 后端错误格式: { error: string }
    const errorMsg = error.response?.data?.error || error.message || '请求失败'
    console.error('响应错误:', errorMsg)
    return Promise.reject(new Error(errorMsg))
  }
)

export default request