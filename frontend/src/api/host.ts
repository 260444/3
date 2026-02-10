import request from './request'

export interface Host {
  id: number
  hostname: string
  ip_address: string
  port: number
  os_type: string
  cpu_cores?: number
  memory_gb?: number
  disk_space_gb?: number
  group_id: number
  status: number
  monitoring_enabled: number
  last_heartbeat?: string
  description: string
  created_at: string
  updated_at?: string
  group?: {
    id: number
    name: string
  }
}

export interface HostQuery {
  page: number
  page_size: number
  hostname?: string
  ip_address?: string
  group_id?: number
  status?: number
  os_type?: string
}

export interface CreateHostReq {
  hostname: string
  ip_address: string
  port: number
  os_type: string
  cpu_cores?: number
  memory_gb?: number
  disk_space_gb?: number
  group_id: number | undefined
  description: string
}

export interface UpdateHostReq {
  hostname: string
  ip_address: string
  port: number
  os_type: string
  cpu_cores?: number
  memory_gb?: number
  disk_space_gb?: number
  group_id: number | undefined
  description: string
}

export function getHosts(params: HostQuery) {
  return request({
    url: '/hosts',
    method: 'get',
    params
  })
}

export function getHost(id: number) {
  return request({
    url: `/hosts/${id}`,
    method: 'get'
  })
}

export function createHost(data: CreateHostReq) {
  return request({
    url: '/hosts',
    method: 'post',
    data
  })
}

export function updateHost(id: number, data: UpdateHostReq) {
  return request({
    url: `/hosts/${id}`,
    method: 'put',
    data
  })
}

export function deleteHost(id: number) {
  return request({
    url: `/hosts/${id}`,
    method: 'delete'
  })
}

export function batchDeleteHosts(ids: number[]) {
  return request({
    url: '/hosts/batch',
    method: 'delete',
    data: { ids }
  })
}

export function updateHostStatus(id: number, status: number) {
  return request({
    url: `/hosts/${id}/status`,
    method: 'put',
    data: { status }
  })
}

export function updateHostMonitoring(id: number, monitoring_enabled: number) {
  return request({
    url: `/hosts/${id}/monitoring`,
    method: 'put',
    data: { monitoring_enabled }
  })
}

export function getHostMetricsLatest(host_id: number) {
  return request({
    url: '/host-metrics/latest',
    method: 'get',
    params: { host_id }
  })
}
