import request from './request'

// 权限管理相关接口

// 为角色分配菜单权限
export const assignMenuToRole = (roleId: number, menuIds: number[]) => {
  return request.post(`/roles/${roleId}/menus`, { menu_ids: menuIds })
}

// 获取角色的菜单权限
export const getRoleMenus = (roleId: number) => {
  return request.get(`/roles/${roleId}/menus`)
}

// 移除角色的菜单权限
export const removeMenuFromRole = (roleId: number, menuIds: number[]) => {
  return request.delete(`/roles/${roleId}/menus`, { data: { menu_ids: menuIds } })
}

// 添加Casbin策略
export const addPolicy = (roleId: number, path: string, method: string) => {
  return request.post(`/roles/${roleId}/policies`, { path, method })
}

// 移除Casbin策略
export const removePolicy = (roleId: number, path: string, method: string) => {
  return request.delete(`/roles/${roleId}/policies`, { data: { path, method } })
}

// 获取角色的Casbin策略
export const getPolicies = (roleId: number) => {
  return request.get(`/roles/${roleId}/policies`)
}

// 获取所有Casbin策略
export const getAllPolicies = () => {
  return request.get('/policies')
}

// 权限资源管理相关接口

// 创建权限
export const createPermission = (data: { path: string; method: string; description?: string; status?: number }) => {
  return request.post('/permissions', data)
}

// 获取权限列表
export const getPermissions = (params?: { page?: number; page_size?: number; path?: string; method?: string }) => {
  return request.get('/permissions', { params })
}

// 获取权限详情
export const getPermission = (id: number) => {
  return request.get(`/permissions/${id}`)
}

// 根据路径和方法获取权限详情
export const getPermissionByPathAndMethod = (path: string, method: string) => {
  return request.get('/permissions/detail', { params: { path, method } })
}

// 更新权限
export const updatePermission = (id: number, data: { path?: string; method?: string; description?: string; status?: number }) => {
  return request.put(`/permissions/${id}`, data)
}

// 更新权限状态
export const updatePermissionStatus = (id: number, status: number) => {
  return request.put(`/permissions/${id}/status`, { status })
}

// 删除权限
export const deletePermission = (id: number) => {
  return request.delete(`/permissions/${id}`)
}