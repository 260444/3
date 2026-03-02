import request from './request'

export interface Credential {
  id: number
  name: string
  username: string
  password?: string
  description: string
  created_at: string
  updated_at?: string
  created_by?: number
  updated_by?: number
}

export interface CredentialQuery {
  page: number
  page_size: number
  name?: string
  username?: string
}

export interface CreateCredentialReq {
  name: string
  username: string
  password: string
  description: string
}

export interface UpdateCredentialReq {
  name: string
  username: string
  password?: string
  description: string
}

export function getCredentials(params: CredentialQuery) {
  return request({
    url: '/credentials',
    method: 'get',
    params
  })
}

export function getCredential(id: number) {
  return request({
    url: `/credentials/${id}`,
    method: 'get'
  })
}

export function createCredential(data: CreateCredentialReq) {
  return request({
    url: '/credentials',
    method: 'post',
    data
  })
}

export function updateCredential(id: number, data: UpdateCredentialReq) {
  return request({
    url: `/credentials/${id}`,
    method: 'put',
    data
  })
}

export function deleteCredential(id: number) {
  return request({
    url: `/credentials/${id}`,
    method: 'delete'
  })
}

export function batchDeleteCredentials(ids: number[]) {
  return request({
    url: '/credentials/batch',
    method: 'delete',
    data: { ids }
  })
}

export function getCredentialsByHost(host_id: number) {
  return request({
    url: '/credentials/host',
    method: 'get',
    params: { host_id }
  })
}