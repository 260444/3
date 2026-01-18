// api/role.ts
import request from './request'

// 创建角色
export const createRole = (data: any) => {
  return request.post('/roles', data)
}

// 获取角色列表
export const getRoles = (params: { page?: number; page_size?: number }) => {
  return request.get('/roles', { params })
}

// 获取角色详情
export const getRole = (id: number) => {
  return request.get(`/roles/${id}`)
}

// 更新角色
export const updateRole = (id: number, data: any) => {
  return request.put(`/roles/${id}`, data)
}

// 删除角色
export const deleteRole = (id: number) => {
  return request.delete(`/roles/${id}`)
}