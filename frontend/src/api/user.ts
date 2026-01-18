// api/user.ts
import request from './request'

// 用户登录
export const login = (username: string, password: string) => {
  return request.post('/login', { username, password })
}

// 用户注册
export const register = (data: { username: string; password: string; email?: string; nickname?: string }) => {
  return request.post('/register', data)
}

// 获取用户信息
export const getUserInfo = () => {
  return request.get('/users/profile')
}

// 获取用户列表
export const getUsers = (params: { page?: number; page_size?: number }) => {
  return request.get('/users', { params })
}

// 创建用户
export const createUser = (data: any) => {
  return request.post('/users', data)
}

// 更新用户信息
export const updateUser = (id: number, data: any) => {
  return request.put(`/users/${id}`, data)
}

// 删除用户
export const deleteUser = (id: number) => {
  return request.delete(`/users/${id}`)
}

// 更新用户状态
export const updateUserStatus = (id: number, status: number) => {
  return request.put(`/users/${id}/status`, { status })
}

// 修改密码
export const changePassword = (oldPassword: string, newPassword: string) => {
  return request.put('/users/change-password', { old_password: oldPassword, new_password: newPassword })
}

// 重置用户密码（管理员功能）
export const resetPassword = (id: number, newPassword: string) => {
  return request.put(`/users/${id}/reset-password`, { new_password: newPassword })
}

// 退出登录
export const logout = () => {
  return request.post('/logout')
}