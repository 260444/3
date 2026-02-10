import request from './request'

// 凭据相关API
export const credentialApi = {
  // 获取凭据列表
  getList: (params?: any) => {
    return request.get('/credentials', { params })
  },

  // 获取凭据详情
  getById: (id: number) => {
    return request.get(`/credentials/${id}`)
  },

  // 创建凭据
  create: (data: any) => {
    return request.post('/credentials', data)
  },

  // 更新凭据
  update: (id: number, data: any) => {
    return request.put(`/credentials/${id}`, data)
  },

  // 删除凭据
  delete: (id: number) => {
    return request.delete(`/credentials/${id}`)
  },

  // 批量删除凭据
  batchDelete: (ids: number[]) => {
    return request.delete('/credentials/batch', { data: { ids } })
  },

  // 获取主机关联的凭据
  getHostCredentials: (hostId: number) => {
    return request.get('/credentials/host', { params: { host_id: hostId } })
  }
}