# 主机管理API文档

## 概述

主机管理模块提供对服务器主机的增删改查、分组管理和监控功能。包括主机信息管理、主机组管理以及主机监控指标等功能。

## 数据模型

### Host（主机模型）

```go
type Host struct {
    ID               uint           `gorm:"primaryKey" json:"id"`
    Hostname         string         `gorm:"size:100;not null;uniqueIndex" json:"hostname"`                    // 主机名
    IPAddress        string         `gorm:"size:45;not null;uniqueIndex" json:"ip_address"`                   // IP地址
    Port             uint16         `gorm:"default:22" json:"port"`                                           // SSH端口
    Username         string         `gorm:"size:50;not null" json:"username"`                                 // 登录用户名
    Password         string         `gorm:"size:255;not null" json:"password"`                                // 加密后的密码
    OSType           string         `gorm:"size:20;default:'linux'" json:"os_type"`                           // 操作系统类型
    CPUCores         *uint16        `json:"cpu_cores"`                                                        // CPU核心数
    MemoryGB         *uint16        `json:"memory_gb"`                                                        // 内存大小(GB)
    DiskSpaceGB      *uint32        `json:"disk_space_gb"`                                                    // 磁盘空间(GB)
    GroupID          uint           `gorm:"not null" json:"group_id"`                                         // 所属主机组ID
    Status           int8           `gorm:"default:1" json:"status"`                                          // 主机状态: 1-在线, 0-离线, -1-故障
    MonitoringEnable int8           `gorm:"default:1" json:"monitoring_enabled"`                              // 监控是否启用
    LastHeartbeat    *time.Time     `json:"last_heartbeat"`                                                   // 最后心跳时间
    Description      string         `gorm:"size:500" json:"description"`                                      // 主机描述
    CreatedBy        *uint          `json:"created_by"`                                                       // 创建人用户ID
    UpdatedBy        *uint          `json:"updated_by"`                                                       // 更新人用户ID
    CreatedAt        time.Time      `json:"created_at"`
    UpdatedAt        time.Time      `json:"updated_at"`
    DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
    Group            *HostGroup     `gorm:"foreignKey:GroupID" json:"group"`                                  // 关联的主机组
}
```

### HostGroup（主机组模型）

```go
type HostGroup struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    Name        string         `gorm:"size:100;not null;uniqueIndex" json:"name"`    // 主机组名称
    Description string         `gorm:"size:500" json:"description"`                  // 描述信息
    Status      int8           `gorm:"default:1" json:"status"`                      // 状态: 1-启用, 0-禁用
    CreatedBy   *uint          `json:"created_by"`                                   // 创建人用户ID
    UpdatedBy   *uint          `json:"updated_by"`                                   // 更新人用户ID
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
    Hosts       []Host         `gorm:"foreignKey:GroupID" json:"hosts"`              // 关联的主机
}
```

### HostMetric（主机监控指标模型）

```go
type HostMetric struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    HostID      uint      `gorm:"not null" json:"host_id"`                      // 主机ID
    MetricType  string    `gorm:"size:30;not null" json:"metric_type"`          // 指标类型: cpu,memory,disk,network
    MetricName  string    `gorm:"size:50;not null" json:"metric_name"`          // 指标名称
    MetricValue float64   `gorm:"type:decimal(10,2);not null" json:"metric_value"` // 指标值
    Unit        string    `gorm:"size:20" json:"unit"`                          // 单位
    RecordedAt  time.Time `gorm:"not null" json:"recorded_at"`                  // 记录时间
    Host        *Host     `gorm:"foreignKey:HostID" json:"host"`                // 关联的主机
}
```

## API接口列表

### 主机管理接口

#### 1. 创建主机
- **URL**: `POST /api/v1/hosts`
- **请求参数**:
```json
{
    "hostname": "web-server-01",
    "ip_address": "192.168.1.100",
    "port": 22,
    "username": "root",
    "password": "encrypted_password",
    "os_type": "linux",
    "cpu_cores": 8,
    "memory_gb": 16,
    "disk_space_gb": 500,
    "group_id": 1,
    "description": "Web服务器01"
}
```
- **成功响应**:
```json
{
    "success": true,
    "message": "主机创建成功",
    "data": {
        "id": 1,
        "hostname": "web-server-01",
        "ip_address": "192.168.1.100",
        "port": 22,
        "username": "root",
        "os_type": "linux",
        "cpu_cores": 8,
        "memory_gb": 16,
        "disk_space_gb": 500,
        "group_id": 1,
        "status": 1,
        "monitoring_enabled": 1,
        "description": "Web服务器01",
        "created_at": "2026-02-04T14:30:00Z",
        "updated_at": "2026-02-04T14:30:00Z"
    }
}
```

- **错误响应**:
```json
{
    "success": false,
    "message": "参数验证失败",
    "error": "hostname: 主机名不能为空"
}
```

#### 2. 获取主机列表
- **URL**: `GET /api/v1/hosts`
- **查询参数**:
  - `page`: 页码，默认1
  - `page_size`: 每页数量，默认10
  - `hostname`: 主机名模糊搜索
  - `ip_address`: IP地址模糊搜索
  - `group_id`: 主机组ID筛选
  - `status`: 状态筛选(1-在线, 0-离线, -1-故障)
  - `os_type`: 操作系统类型筛选
- **成功响应**:
```json
{
    "success": true,
    "message": "获取主机列表成功",
    "data": {
        "list": [
            {
                "id": 1,
                "hostname": "web-server-01",
                "ip_address": "192.168.1.100",
                "port": 22,
                "username": "root",
                "os_type": "linux",
                "cpu_cores": 8,
                "memory_gb": 16,
                "disk_space_gb": 500,
                "group_id": 1,
                "status": 1,
                "monitoring_enabled": 1,
                "last_heartbeat": "2026-02-04T14:25:00Z",
                "description": "Web服务器01",
                "created_at": "2026-02-04T14:30:00Z",
                "group": {
                    "id": 1,
                    "name": "Web服务器组",
                    "description": "Web服务器主机组"
                }
            }
        ],
        "pagination": {
            "page": 1,
            "page_size": 10,
            "total": 25
        }
    }
}
```

- **错误响应**:
```json
{
    "success": false,
    "message": "服务器内部错误",
    "error": "数据库查询失败"
}
```

#### 3. 获取主机详情
- **URL**: `GET /api/v1/hosts/{id}`
- **路径参数**:
  - `id`: 主机ID
- **成功响应**:
```json
{
    "success": true,
    "message": "获取主机详情成功",
    "data": {
        "id": 1,
        "hostname": "web-server-01",
        "ip_address": "192.168.1.100",
        "port": 22,
        "username": "root",
        "os_type": "linux",
        "cpu_cores": 8,
        "memory_gb": 16,
        "disk_space_gb": 500,
        "group_id": 1,
        "status": 1,
        "monitoring_enabled": 1,
        "last_heartbeat": "2026-02-04T14:25:00Z",
        "description": "Web服务器01",
        "created_at": "2026-02-04T14:30:00Z",
        "updated_at": "2026-02-04T14:30:00Z",
        "group": {
            "id": 1,
            "name": "Web服务器组",
            "description": "Web服务器主机组"
        }
    }
}
```

- **错误响应**:
```json
{
    "success": false,
    "message": "资源不存在",
    "error": "主机不存在"
}
```

#### 4. 更新主机信息
- **URL**: `PUT /api/v1/hosts/{id}`
- **路径参数**:
  - `id`: 主机ID
- **请求参数**:
```json
{
    "hostname": "web-server-01",
    "ip_address": "192.168.1.100",
    "port": 22,
    "username": "root",
    "password": "new_encrypted_password",
    "os_type": "linux",
    "cpu_cores": 16,
    "memory_gb": 32,
    "disk_space_gb": 1000,
    "group_id": 1,
    "description": "升级后的Web服务器01"
}
```
- **成功响应**:
```json
{
    "success": true,
    "message": "主机更新成功",
    "data": {
        "id": 1,
        "hostname": "web-server-01",
        "ip_address": "192.168.1.100",
        "port": 22,
        "username": "root",
        "os_type": "linux",
        "cpu_cores": 16,
        "memory_gb": 32,
        "disk_space_gb": 1000,
        "group_id": 1,
        "status": 1,
        "monitoring_enabled": 1,
        "description": "升级后的Web服务器01",
        "updated_at": "2026-02-04T15:00:00Z"
    }
}
```

- **错误响应**:
```json
{
    "success": false,
    "message": "权限不足",
    "error": "无权修改该主机"
}
```

#### 5. 删除主机
- **URL**: `DELETE /api/v1/hosts/{id}`
- **路径参数**:
  - `id`: 主机ID
- **成功响应**:
```json
{
    "success": true,
    "message": "主机删除成功"
}
```

- **错误响应**:
```json
{
    "success": false,
    "message": "资源不存在",
    "error": "主机不存在"
}
```

#### 6. 批量删除主机
- **URL**: `DELETE /api/v1/hosts/batch`
- **请求参数**:
```json
{
    "ids": [1, 2, 3]
}
```
- **成功响应**:
```json
{
    "success": true,
    "message": "批量删除主机成功",
    "data": {
        "deleted_count": 3
    }
}
```

- **错误响应**:
```json
{
    "success": false,
    "message": "参数错误",
    "error": "ids参数不能为空"
}
```

#### 7. 更新主机状态
- **URL**: `PUT /api/v1/hosts/{id}/status`
- **路径参数**:
  - `id`: 主机ID
- **请求参数**:
```json
{
    "status": 2
}
```
- **成功响应**:
```json
{
    "success": true,
    "message": "主机状态更新成功",
    "data": {
        "id": 1,
        "status": 2
    }
}
```
- **备注**:
- 状态值：1-在线, 2-离线, 3-故障

- **错误响应**:
```json
{
    "success": false,
    "message": "参数验证失败",
    "error": "status: 状态值无效"
}
```

#### 8. 更新监控状态
- **URL**: `PUT /api/v1/hosts/{id}/monitoring`
- **路径参数**:
  - `id`: 主机ID
- **请求参数**:
```json
{
    "monitoring_enabled": 2
}

```
- **响应**:
```json
{
    "message": "监控状态更新成功",
    "data": {
        "id": 1,
        "monitoring_enabled": 2
    }
}
```
- **备注**:
- 状态值：1-开启, 2-关闭


### 主机组管理接口

#### 9. 创建主机组
- **URL**: `POST /api/v1/host-groups`
- **请求参数**:
```json
{
    "name": "数据库服务器组",
    "description": "用于存放数据库服务器的主机组"
}
```
- **成功响应**:
```json
{
    "success": true,
    "message": "主机组创建成功",
    "data": {
        "id": 1,
        "name": "数据库服务器组",
        "description": "用于存放数据库服务器的主机组",
        "status": 1,
        "created_at": "2026-02-04T14:30:00Z"
    }
}
```

- **错误响应**:
```json
{
    "success": false,
    "message": "参数验证失败",
    "error": "name: 主机组名称已存在"
}
```

#### 10. 获取主机组列表
- **URL**: `GET /api/v1/host-groups`
- **查询参数**:
  - `page`: 页码，默认1
  - `page_size`: 每页数量，默认10
  - `name`: 主机组名称模糊搜索
  - `status`: 状态筛选
- **成功响应**:
```json
{
    "success": true,
    "message": "获取主机组列表成功",
    "data": {
        "list": [
            {
                "id": 1,
                "name": "Web服务器组",
                "description": "Web服务器主机组",
                "status": 1,
                "host_count": 5,
                "created_at": "2026-02-04T14:30:00Z"
            },
            {
                "id": 2,
                "name": "数据库服务器组",
                "description": "数据库服务器主机组",
                "status": 1,
                "host_count": 3,
                "created_at": "2026-02-04T14:35:00Z"
            }
        ],
        "pagination": {
            "page": 1,
            "page_size": 10,
            "total": 2
        }
    }
}
```

- **错误响应**:
```json
{
    "success": false,
    "message": "服务器内部错误",
    "error": "数据库查询失败"
}
```

#### 11. 获取主机组详情
- **URL**: `GET /api/v1/host-groups/{id}`
- **路径参数**:
  - `id`: 主机组ID
- **成功响应**:
```json
{
    "success": true,
    "message": "获取主机组详情成功",
    "data": {
        "id": 1,
        "name": "Web服务器组",
        "description": "Web服务器主机组",
        "status": 1,
        "created_at": "2026-02-04T14:30:00Z",
        "hosts": [
            {
                "id": 1,
                "hostname": "web-server-01",
                "ip_address": "192.168.1.100",
                "status": 1
            }
        ]
    }
}
```

- **错误响应**:
```json
{
    "success": false,
    "message": "资源不存在",
    "error": "主机组不存在"
}
```

#### 12. 更新主机组
- **URL**: `PUT /api/v1/host-groups/{id}`
- **路径参数**:
  - `id`: 主机组ID
- **请求参数**:
```json
{
    "name": "Web应用服务器组",
    "description": "Web应用服务器主机组"
}
```
- **成功响应**:
```json
{
    "success": true,
    "message": "主机组更新成功",
    "data": {
        "id": 1,
        "name": "Web应用服务器组",
        "description": "Web应用服务器主机组",
        "updated_at": "2026-02-04T15:00:00Z"
    }
}
```

- **错误响应**:
```json
{
    "success": false,
    "message": "权限不足",
    "error": "无权修改该主机组"
}
```

#### 13. 删除主机组
- **URL**: `DELETE /api/v1/host-groups/{id}`
- **路径参数**:
  - `id`: 主机组ID
- **成功响应**:
```json
{
    "success": true,
    "message": "主机组删除成功"
}
```

- **错误响应**:
```json
{
    "success": false,
    "message": "资源不存在",
    "error": "主机组不存在"
}
```

#### 14. 更新主机组状态
- **URL**: `PUT /api/v1/host-groups/{id}/status`
- **路径参数**:
  - `id`: 主机组ID
- **请求参数**:
```json
{
    "status": 1
}
```
- **响应**:
```json
{
    "message": "主机组状态更新成功",
    "data": {
        "id": 1,
        "status": 1
    }
}
```

- **备注**:
- 状态值：1-开启, 2-关闭

### 主机监控指标接口

#### 15. 上报主机指标
- **URL**: `POST /api/v1/host-metrics`
- **请求参数**:
```json
{
    "host_id": 1,
    "metrics": [
        {
            "metric_type": "cpu",
            "metric_name": "cpu_usage",
            "metric_value": 45.5,
            "unit": "%"
        },
        {
            "metric_type": "memory",
            "metric_name": "memory_usage",
            "metric_value": 68.2,
            "unit": "%"
        }
    ]
}
```
- **成功响应**:
```json
{
    "success": true,
    "message": "指标上报成功",
    "data": {
        "inserted_count": 2
    }
}
```

- **错误响应**:
```json
{
    "success": false,
    "message": "参数验证失败",
    "error": "host_id: 主机不存在"
}
```

#### 16. 获取主机指标历史
- **URL**: `GET /api/v1/host-metrics/history`
- **查询参数**:
  - `host_id`: 主机ID（必填）
  - `metric_type`: 指标类型(cpu/memory/disk/network)
  - `metric_name`: 指标名称
  - `start_time`: 开始时间
  - `end_time`: 结束时间
  - `page`: 页码，默认1
  - `page_size`: 每页数量，默认50
- **成功响应**:
```json
{
    "success": true,
    "message": "获取指标历史成功",
    "data": {
        "list": [
            {
                "id": 1,
                "host_id": 1,
                "metric_type": "cpu",
                "metric_name": "cpu_usage",
                "metric_value": 45.5,
                "unit": "%",
                "recorded_at": "2026-02-04T14:30:00Z"
            }
        ],
        "pagination": {
            "page": 1,
            "page_size": 50,
            "total": 100
        }
    }
}
```

- **错误响应**:
```json
{
    "success": false,
    "message": "参数错误",
    "error": "host_id参数不能为空"
}
```

#### 17. 获取主机最新指标
- **URL**: `GET /api/v1/host-metrics/latest`
- **查询参数**:
  - `host_id`: 主机ID（必填）
- **成功响应**:
```json
{
    "success": true,
    "message": "获取最新指标成功",
    "data": [
        {
            "metric_type": "cpu",
            "metric_name": "cpu_usage",
            "metric_value": 45.5,
            "unit": "%",
            "recorded_at": "2026-02-04T14:30:00Z"
        },
        {
            "metric_type": "memory",
            "metric_name": "memory_usage",
            "metric_value": 68.2,
            "unit": "%",
            "recorded_at": "2026-02-04T14:30:00Z"
        }
    ]
}
```

- **错误响应**:
```json
{
    "success": false,
    "message": "资源不存在",
    "error": "主机不存在"
}
```

#### 18. 获取主机统计信息
- **URL**: `GET /api/v1/hosts/statistics`
- **查询参数**:
  - `host_id`: 主机ID（必填）
  - `metric_type`: 指标类型(cpu/memory/disk/network) (必填)
  - `metric_name`: 指标名称 (必填)
  - `start_time`: 开始时间
  - `end_time`: 结束时间
  - `page`: 页码，默认1
  - `page_size`: 每页数量，默认50
- **成功响应**:
```json
{
    "success": true,
    "message": "获取主机统计信息成功",
    "data": {
        "total_hosts": 25,
        "online_hosts": 20,
        "offline_hosts": 3,
        "fault_hosts": 2,
        "enabled_monitoring": 22,
        "disabled_monitoring": 3,
        "by_os_type": {
            "linux": 20,
            "windows": 5
        },
        "by_group": [
            {
                "group_id": 1,
                "group_name": "Web服务器组",
                "host_count": 10
            }
        ]
    }
}
```

- **错误响应**:
```json
{
    "success": false,
    "message": "服务器内部错误",
    "error": "统计查询失败"
}
```

## 错误响应格式

所有API接口的错误响应格式统一如下：

```json
{
    "success": false,
    "message": "错误信息描述",
    "error": "详细的错误信息"
}
```

常见错误码对应的消息：
- `400`: 参数错误/数据验证失败
- `401`: 未授权访问
- `403`: 权限不足
- `404`: 资源不存在
- `500`: 服务器内部错误


## 注意事项

1. 所有涉及密码的操作都需要对密码进行加密处理
2. 主机名和IP地址在系统中必须唯一
3. 删除操作使用软删除，数据不会真正从数据库中删除
4. 监控指标数据量较大，建议定期清理历史数据
5. 主机组删除前需要确保组内没有主机
6. 接口调用需要携带有效的JWT Token进行身份验证
7. 所有API响应都遵循统一的标准格式，包含`success`、`message`和`data`/`error`字段
8. 成功响应的`success`字段为`true`，错误响应的`success`字段为`false`