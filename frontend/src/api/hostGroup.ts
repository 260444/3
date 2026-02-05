import request from './request'

export interface HostGroup {
  id: number
  name: string
  description: string
  status: number
  host_count?: number
  created_at: string
  updated_at?: string
}

export interface HostGroupQuery {
  page: number
  page_size: number
  name?: string
  status?: number
}

export interface CreateHostGroupReq {
  name: string
  description: string
}

export interface UpdateHostGroupReq {
  name: string
  description: string
}

export function getHostGroups(params: HostGroupQuery) {
  return request({
    url: '/host-groups',
    method: 'get',
    params
  })
}

export function getHostGroup(id: number) {
  return request({
    url: `/host-groups/${id}`,
    method: 'get'
  })
}

export function createHostGroup(data: CreateHostGroupReq) {
  return request({
    url: '/host-groups',
    method: 'post',
    data
  })
}

export function updateHostGroup(id: number, data: UpdateHostGroupReq) {
  return request({
    url: `/host-groups/${id}`,
    method: 'put',
    data
  })
}

export function deleteHostGroup(id: number) {
  return request({
    url: `/host-groups/${id}`,
    method: 'delete'
  })
}

export function updateHostGroupStatus(id: number, status: number) {
  return request({
    url: `/host-groups/${id}/status`,
    method: 'put',
    data: { status }
  })
}
