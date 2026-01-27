# 权限管理接口文档

## 基本信息

- **基础URL**: `http://localhost:8080/api/v1`
- **认证方式**: JWT Token
- **数据格式**: JSON
- **字符编码**: UTF-8

---

## 权限资源管理接口

### 1. 创建权限

**接口地址**: `POST /permissions`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| path | string | 是 | 请求路径 |
| method | string | 是 | 请求方法（GET、POST、PUT、DELETE等） |
| description | string | 否 | 权限描述 |
| status | int8 | 否 | 状态(1:正常, 0:禁用)，默认1 |

**请求示例**:
```json
{
  "path": "/api/v1/users",
  "method": "GET",
  "description": "获取用户列表权限",
  "status": 1
}
```

**成功响应**:
```json
{
  "message": "权限创建成功",
  "data": {
    "id": 1,
    "path": "/api/v1/users",
    "method": "GET",
    "description": "获取用户列表权限",
    "status": 1,
    "created_at": "2026-01-18T10:00:00Z"
  }
}
```

---

### 2. 获取权限列表

**接口地址**: `GET /permissions`

**请求头**:
```
Authorization: Bearer <token>
```

**请求参数**:

| 参数名 | 类型 | 必填 | 默认值 | 说明 |
|--------|------|------|--------|------|
| page | int | 否 | 1 | 页码 |
| page_size | int | 否 | 10 | 每页数量(最大100) |

**请求示例**:
```
GET /permissions?page=1&page_size=10
```

**成功响应**:
```json
{
  "message": "获取成功",
  "data": {
    "list": [
      {
        "id": 1,
        "path": "/api/v1/users",
        "method": "GET",
        "description": "获取用户列表权限",
        "status": 1,
        "created_at": "2026-01-18T10:00:00Z"
      }
    ],
    "total": 100,
    "page": 1,
    "page_size": 10
  }
}
```

---

### 3. 获取权限详情

**接口地址**: `GET /permissions/:id`

**请求头**:
```
Authorization: Bearer <token>
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 权限ID |

**请求示例**:
```
GET /permissions/1
```

**成功响应**:
```json
{
  "message": "获取成功",
  "data": {
    "id": 1,
    "path": "/api/v1/users",
    "method": "GET",
    "description": "获取用户列表权限",
    "status": 1,
    "created_at": "2026-01-18T10:00:00Z"
  }
}
```

---

### 4. 更新权限

**接口地址**: `PUT /permissions/:id`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 权限ID |

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| path | string | 否 | 请求路径 |
| method | string | 否 | 请求方法 |
| description | string | 否 | 权限描述 |
| status | int8 | 否 | 状态(1:正常, 0:禁用) |

**请求示例**:
```json
{
  "path": "/api/v1/users",
  "method": "POST",
  "description": "创建用户权限"
}
```

**成功响应**:
```json
{
  "message": "更新成功"
}
```

---

### 5. 更新权限状态

**接口地址**: `PUT /permissions/:id/status`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/x-www-form-urlencoded
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 权限ID |

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| status | int8 | 是 | 状态(1:正常, 0:禁用) |

**请求示例**:
```
status=0
```

**成功响应**:
```json
{
  "message": "更新状态成功"
}
```

---

### 6. 删除权限

**接口地址**: `DELETE /permissions/:id`

**请求头**:
```
Authorization: Bearer <token>
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 权限ID |

**请求示例**:
```
DELETE /permissions/1
```

**成功响应**:
```json
{
  "message": "删除成功"
}
```

---

## 数据模型

### Permission（权限模型）

| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | uint | 权限ID |
| path | string | 请求路径 |
| method | string | 请求方法 |
| description | string | 权限描述 |
| status | int8 | 状态(1:正常, 0:禁用) |
| created_at | time.Time | 创建时间 |
| updated_at | time.Time | 更新时间 |

---

## HTTP 状态码说明

| 状态码 | 说明 |
|--------|------|
| 200 | 请求成功 |
| 400 | 请求参数错误 |
| 401 | 未授权（Token无效或过期） |
| 403 | 禁止访问（无权限） |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

---

## 错误码说明

| 错误码 | 说明 |
|--------|------|
| 用户名已存在 | 注册时用户名重复 |
| 邮箱已存在 | 注册时邮箱重复 |
| 用户名或密码错误 | 登录失败 |
| Token无效或已过期 | 认证失败 |
| 权限不足 | 无权限访问 |
| 资源不存在 | 记录不存在 |
| 数据验证失败 | 参数验证失败 |

---

## 注意事项

1. **认证**: 除公开接口外，所有接口都需要在请求头中携带有效的 JWT Token
2. **分页**: 列表接口支持分页，默认每页10条，最大100条
3. **软删除**: 用户、角色、菜单支持软删除，删除后不会物理删除数据
4. **时间格式**: 所有时间字段使用 ISO 8601 格式（如：`2026-01-18T10:00:00Z`）
5. **密码安全**: 密码使用 bcrypt 加密存储，传输时建议使用 HTTPS
6. **权限控制**: 系统基于 RBAC 模型进行权限控制，通过角色分配权限
7. **操作日志**: 系统自动记录用户操作，包括请求和响应信息