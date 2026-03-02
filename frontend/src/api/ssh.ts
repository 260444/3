import request from './request'

export interface TestSSHRequest {
  host_id: number
  credential_id: number
}

export function testSSHConnection(data: TestSSHRequest) {
  return request({
    url: '/ssh/test',
    method: 'post',
    data
  })
}