# 用户管理接口文档

## 基本信息

- **基础URL**: `http://localhost:8080/api/v1`
- **认证方式**: JWT Token
- **数据格式**: JSON
- **字符编码**: UTF-8

---

## 用户管理接口

### 1. 创建用户（需要认证）

**接口地址**: `POST /users`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| username | string | 是 | 用户名，唯一 |
| password | string | 是 | 密码 |
| email | string | 否 | 邮箱，唯一 |
| nickname | string | 否 | 昵称 |
| role_id | uint | 否 | 角色ID |

**请求示例**:
```json
{
  "username": "testuser",
  "password": "123456",
  "email": "test@example.com",
  "nickname": "测试用户",
  "role_id": 1
}
```

**成功响应**:
```json
{
  "message": "用户创建成功",
  "data": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "nickname": "测试用户",
    "status": 1,
    "role_id": 1,
    "created_at": "2026-01-18T10:00:00Z"
  }
}
```

**错误响应**:
```json
{
  "error": "用户名已存在"
}
```

---

### 2. 获取用户列表

**接口地址**: `GET /users`

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
GET /users?page=1&page_size=10
```

**成功响应**:
```json
{
  "message": "获取成功",
  "data": {
    "list": [
      {
        "id": 1,
        "username": "testuser",
        "email": "test@example.com",
        "nickname": "测试用户",
        "phone": "13800138000",
        "avatar": "",
        "status": 1,
        "role_id": 1,
        "role": {
          "id": 1,
          "name": "管理员",
          "description": "系统管理员"
        },
        "created_at": "2026-01-18T10:00:00Z",
        "updated_at": "2026-01-18T10:00:00Z"
      }
    ],
    "total": 100,
    "page": 1,
    "page_size": 10
  }
}
```

---

### 3. 获取用户详情

**接口地址**: `GET /users/:id`

**请求头**:
```
Authorization: Bearer <token>
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 用户ID |

**请求示例**:
```
GET /users/1
```

**成功响应**:
```json
{
  "message": "获取成功",
  "data": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "nickname": "测试用户",
    "phone": "13800138000",
    "avatar": "",
    "status": 1,
    "role_id": 1,
    "role": {
      "id": 1,
      "name": "管理员",
      "description": "系统管理员"
    },
    "created_at": "2026-01-18T10:00:00Z",
    "updated_at": "2026-01-18T10:00:00Z"
  }
}
```

---

### 4. 更新用户信息

**接口地址**: `PUT /users/:id`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 用户ID |

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| email | string | 否 | 邮箱 |
| phone | string | 否 | 手机号 |
| nickname | string | 否 | 昵称 |
| avatar | string | 否 | 头像URL |
| role_id | int | 否 | 角色ID |

**请求示例**:
```json
{
  "email": "newemail@example.com",
  "nickname": "新昵称",
  "role_id": 2
}
```

**成功响应**:
```json
{
  "message": "更新成功"
}
```

---

### 5. 更新用户状态

**接口地址**: `PUT /users/:id/status`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/x-www-form-urlencoded
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 用户ID |

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| status | int | 是 | 状态(1:正常, 0:禁用) |

**请求示例**:
```
status=1
```

**成功响应**:
```json
{
  "message": "更新状态成功"
}
```

---

### 6. 删除用户

**接口地址**: `DELETE /users/:id`

**请求头**:
```
Authorization: Bearer <token>
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 用户ID |

**请求示例**:
```
DELETE /users/1
```

**成功响应**:
```json
{
  "message": "删除成功"
}
```

---

### 7. 修改密码

**接口地址**: `PUT /users/change-password`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| old_password | string | 是 | 旧密码 |
| new_password | string | 是 | 新密码 |

**请求示例**:
```json
{
  "old_password": "123456",
  "new_password": "654321"
}
```

**成功响应**:
```json
{
  "message": "密码修改成功"
}
```

**错误响应**:
```json
{
  "error": "旧密码错误"
}
```

---

## 数据模型

### User（用户模型）

| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | uint | 用户ID |
| username | string | 用户名，唯一 |
| password | string | 密码（加密存储） |
| email | string | 邮箱，唯一 |
| phone | string | 手机号 |
| nickname | string | 昵称 |
| avatar | string | 头像URL |
| status | int | 状态(1:正常, 0:禁用) |
| last_login_at | time.Time | 最后登录时间 |
| last_login_ip | string | 最后登录IP |
| role_id | uint | 角色ID |
| role | Role | 关联角色 |
| created_at | time.Time | 创建时间 |
| updated_at | time.Time | 更新时间 |
| deleted_at | gorm.DeletedAt | 删除时间（软删除） |

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