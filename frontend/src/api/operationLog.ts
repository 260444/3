// api/operationLog.ts
import request from './request'

// 获取操作日志列表
export const getOperationLogs = (params: { page?: number; page_size?: number }) => {
  return request.get('/operation-logs', { params })
}

// 删除操作日志
export const deleteOperationLog = (id: number) => {
  return request.delete(`/operation-logs/${id}`)
}