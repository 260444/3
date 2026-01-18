// api/menu.ts
import request from './request'

// 创建菜单
export const createMenu = (data: any) => {
  return request.post('/menus', data)
}

// 获取菜单树（用于菜单管理）
export const getMenuTree = (parentId?: number) => {
  const params: any = {}
  if (parentId !== undefined) {
    params.parent_id = parentId
  }
  return request.get('/menus/tree', { params })
}

// 获取所有菜单
export const getAllMenus = () => {
  return request.get('/menus/all')
}

// 获取菜单详情
export const getMenu = (id: number) => {
  return request.get(`/menus/${id}`)
}

// 更新菜单
export const updateMenu = (id: number, data: any) => {
  return request.put(`/menus/${id}`, data)
}

// 删除菜单
export const deleteMenu = (id: number) => {
  return request.delete(`/menus/${id}`)
}

// 获取当前用户的菜单（用于动态路由）
export const getUserMenus = () => {
  return request.get('/menus')
}