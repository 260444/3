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

// 获取角色的菜单权限
export const getRoleMenus = (roleId: number) => {
  return request.get(`/roles/${roleId}/menus`)
}

// 为角色分配菜单权限
export const assignRoleMenus = (roleId: number, menuIds: number[]) => {
  return request.post(`/roles/${roleId}/menus`, { menu_ids: menuIds })
}

// 移除角色的菜单权限
export const removeRoleMenus = (roleId: number, menuIds: number[]) => {
  return request.delete(`/roles/${roleId}/menus`, { data: { menu_ids: menuIds } })
}