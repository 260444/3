# API 接口文档

## 基本信息

- **基础URL**: `http://localhost:8080/api/v1`
- **认证方式**: JWT Token
- **数据格式**: JSON
- **字符编码**: UTF-8

## 通用响应格式

### 成功响应

```json
{
  "message": "操作成功",
  "data": { ... }
}
```

### 错误响应

```json
{
  "error": "错误信息"
}
```

### 列表响应

```json
{
  "message": "获取成功",
  "data": {
    "list": [ ... ],
    "total": 100,
    "page": 1,
    "page_size": 10
  }
}
```

## 认证说明

除公开接口外，所有接口都需要在请求头中携带 JWT Token：

```
Authorization: Bearer <token>
```

---

## 公开接口

### 1. 用户登录

**接口地址**: `POST /login`

**请求头**:
```
Content-Type: application/json
```

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| username | string | 是 | 用户名 |
| password | string | 是 | 密码 |

**请求示例**:
```json
{
  "username": "testuser",
  "password": "123456"
}
```

**成功响应**:
```json
{
  "message": "登录成功",
  "data": {
    "user": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "nickname": "测试用户",
      "status": 1,
      "role_id": 1,
      "last_login_at": "2026-01-18T10:00:00Z",
      "last_login_ip": "127.0.0.1"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

**错误响应**:
```json
{
  "error": "用户名或密码错误"
}
```

---

### 2. 退出登录

**接口地址**: `POST /logout`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**请求参数**: 无

**请求示例**:
```bash
POST /api/v1/logout
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**成功响应**:
```json
{
  "message": "退出登录成功",
  "data": {
    "user_id": 1,
    "username": "admin"
  }
}
```

**说明**:
- 此接口需要 JWT 认证
- 由于 JWT 是无状态的，客户端只需删除本地存储的 token 即可实现退出登录
- 如果使用 Redis 存储 token，可以将 token 加入黑名单以实现强制失效

---

### 3. 获取验证码

**接口地址**: `GET /captcha`

**请求头**: 无需特殊请求头

**请求参数**: 无

**成功响应**:
```json
{
  "message": "验证码生成成功",
  "data": {
    "id": "captcha_id_123456",
    "image": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAA..."
  }
}
```

**错误响应**:
```json
{
  "message": "验证码生成失败"
}
```

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

## 角色管理接口

### 1. 创建角色

**接口地址**: `POST /roles`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| name | string | 是 | 角色名称，唯一 |
| description | string | 否 | 角色描述 |
| status | int | 否 | 状态(1:正常, 0:禁用)，默认1 |

**请求示例**:
```json
{
  "name": "编辑",
  "description": "内容编辑角色",
  "status": 1
}
```

**成功响应**:
```json
{
  "message": "角色创建成功",
  "data": {
    "id": 1,
    "name": "编辑",
    "description": "内容编辑角色",
    "status": 1,
    "created_at": "2026-01-18T10:00:00Z"
  }
}
```

---

### 2. 获取角色列表

**接口地址**: `GET /roles`

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
GET /roles?page=1&page_size=10
```

**成功响应**:
```json
{
  "message": "获取成功",
  "data": {
    "list": [
      {
        "id": 1,
        "name": "管理员",
        "description": "系统管理员",
        "status": 1,
        "users": [ ... ],
        "menus": [ ... ],
        "created_at": "2026-01-18T10:00:00Z",
        "updated_at": "2026-01-18T10:00:00Z"
      }
    ],
    "total": 10,
    "page": 1,
    "page_size": 10
  }
}
```

---

### 3. 获取角色详情

**接口地址**: `GET /roles/:id`

**请求头**:
```
Authorization: Bearer <token>
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 角色ID |

**请求示例**:
```
GET /roles/1
```

**成功响应**:
```json
{
  "message": "获取成功",
  "data": {
    "id": 1,
    "name": "管理员",
    "description": "系统管理员",
    "status": 1,
    "users": [
      {
        "id": 1,
        "username": "admin",
        "email": "admin@example.com"
      }
    ],
    "menus": [
      {
        "id": 1,
        "name": "dashboard",
        "title": "仪表盘",
        "path": "/dashboard"
      }
    ],
    "created_at": "2026-01-18T10:00:00Z",
    "updated_at": "2026-01-18T10:00:00Z"
  }
}
```

---

### 4. 更新角色

**接口地址**: `PUT /roles/:id`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 角色ID |

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| name | string | 否 | 角色名称 |
| description | string | 否 | 角色描述 |
| status | int | 否 | 状态(1:正常, 0:禁用) |
| menus | array | 否 | 菜单ID列表 |

**请求示例**:
```json
{
  "name": "超级管理员",
  "description": "拥有所有权限",
  "status": 1
}
```

**成功响应**:
```json
{
  "message": "更新成功"
}
```

---

### 5. 删除角色

**接口地址**: `DELETE /roles/:id`

**请求头**:
```
Authorization: Bearer <token>
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 角色ID |

**请求示例**:
```
DELETE /roles/1
```

**成功响应**:
```json
{
  "message": "删除成功"
}
```

---

## 菜单管理接口

### 1. 创建菜单

**接口地址**: `POST /menus`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| name | string | 是 | 菜单名称 |
| title | string | 是 | 菜单标题 |
| path | string | 否 | 路由路径 |
| component | string | 否 | 组件路径 |
| redirect | string | 否 | 重定向路径 |
| parent_id | int | 否 | 父级菜单ID |
| icon | string | 否 | 图标 |
| sort | int | 否 | 排序，默认0 |
| is_hidden | bool | 否 | 是否隐藏，默认false |
| is_link | bool | 否 | 是否外部链接，默认false |
| link_url | string | 否 | 外部链接地址 |
| status | int | 否 | 状态(1:正常, 0:禁用)，默认1 |

**请求示例**:
```json
{
  "name": "system",
  "title": "系统管理",
  "path": "/system",
  "icon": "setting",
  "sort": 1,
  "status": 1
}
```

**成功响应**:
```json
{
  "message": "菜单创建成功",
  "data": {
    "id": 1,
    "name": "system",
    "title": "系统管理",
    "path": "/system",
    "icon": "setting",
    "sort": 1,
    "status": 1,
    "created_at": "2026-01-18T10:00:00Z"
  }
}
```

---

### 2. 获取菜单树

**接口地址**: `GET /menus`

**请求头**:
```
Authorization: Bearer <token>
```

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| parent_id | int | 否 | 父级菜单ID，不传则获取所有根菜单 |

**请求示例**:
```
GET /menus
```

**成功响应**:
```json
{
  "message": "获取成功",
  "data": [
    {
      "id": 1,
      "name": "system",
      "title": "系统管理",
      "path": "/system",
      "icon": "setting",
      "sort": 1,
      "status": 1,
      "children": [
        {
          "id": 2,
          "name": "user",
          "title": "用户管理",
          "path": "/system/user",
          "component": "system/user/index",
          "parent_id": 1,
          "sort": 1,
          "status": 1
        }
      ],
      "created_at": "2026-01-18T10:00:00Z",
      "updated_at": "2026-01-18T10:00:00Z"
    }
  ]
}
```

---

### 3. 获取所有菜单

**接口地址**: `GET /menus/all`

**请求头**:
```
Authorization: Bearer <token>
```

**请求参数**: 无

**请求示例**:
```
GET /menus/all
```

**成功响应**:
```json
{
  "message": "获取成功",
  "data": [
    {
      "id": 1,
      "name": "system",
      "title": "系统管理",
      "path": "/system",
      "icon": "setting",
      "sort": 1,
      "status": 1,
      "created_at": "2026-01-18T10:00:00Z",
      "updated_at": "2026-01-18T10:00:00Z"
    },
    {
      "id": 2,
      "name": "user",
      "title": "用户管理",
      "path": "/system/user",
      "component": "system/user/index",
      "parent_id": 1,
      "sort": 1,
      "status": 1,
      "created_at": "2026-01-18T10:00:00Z",
      "updated_at": "2026-01-18T10:00:00Z"
    }
  ]
}
```

---

### 4. 获取菜单详情

**接口地址**: `GET /menus/:id`

**请求头**:
```
Authorization: Bearer <token>
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 菜单ID |

**请求示例**:
```
GET /menus/1
```

**成功响应**:
```json
{
  "message": "获取成功",
  "data": {
    "id": 1,
    "name": "system",
    "title": "系统管理",
    "path": "/system",
    "icon": "setting",
    "sort": 1,
    "status": 1,
    "created_at": "2026-01-18T10:00:00Z",
    "updated_at": "2026-01-18T10:00:00Z"
  }
}
```

---

### 5. 更新菜单

**接口地址**: `PUT /menus/:id`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 菜单ID |

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| name | string | 否 | 菜单名称 |
| title | string | 否 | 菜单标题 |
| path | string | 否 | 路由路径 |
| component | string | 否 | 组件路径 |
| redirect | string | 否 | 重定向路径 |
| parent_id | int | 否 | 父级菜单ID |
| icon | string | 否 | 图标 |
| sort | int | 否 | 排序 |
| is_hidden | bool | 否 | 是否隐藏 |
| is_link | bool | 否 | 是否外部链接 |
| link_url | string | 否 | 外部链接地址 |
| status | int | 否 | 状态(1:正常, 0:禁用) |

**请求示例**:
```json
{
  "title": "系统设置",
  "sort": 2
}
```

**成功响应**:
```json
{
  "message": "更新成功"
}
```

---

### 6. 删除菜单

**接口地址**: `DELETE /menus/:id`

**请求头**:
```
Authorization: Bearer <token>
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 菜单ID |

**请求示例**:
```
DELETE /menus/1
```

**成功响应**:
```json
{
  "message": "删除成功"
}
```

---

## 操作日志接口

### 1. 获取操作日志列表

**接口地址**: `GET /operation-logs`

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
GET /operation-logs?page=1&page_size=10
```

**成功响应**:
```json
{
  "message": "获取成功",
  "data": {
    "list": [
      {
        "id": 1,
        "operation": "用户登录",
        "user_id": 1,
        "username": "admin",
        "ip": "127.0.0.1",
        "user_agent": "Mozilla/5.0...",
        "status": 1,
        "request_method": "POST",
        "request_path": "/api/v1/login",
        "request_body": "{\"username\":\"admin\",\"password\":\"***\"}",
        "response_body": "{\"message\":\"登录成功\"}",
        "response_time": 150,
        "created_at": "2026-01-18T10:00:00Z"
      }
    ],
    "total": 1000,
    "page": 1,
    "page_size": 10
  }
}
```

---

### 2. 删除操作日志

**接口地址**: `DELETE /operation-logs/:id`

**请求头**:
```
Authorization: Bearer <token>
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 日志ID |

**请求示例**:
```
DELETE /operation-logs/1
```

**成功响应**:
```json
{
  "message": "删除成功"
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

### Role（角色模型）

| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | uint | 角色ID |
| name | string | 角色名称，唯一 |
| description | string | 角色描述 |
| status | int | 状态(1:正常, 0:禁用) |
| users | []User | 关联用户列表 |
| menus | []Menu | 关联菜单列表 |
| created_at | time.Time | 创建时间 |
| updated_at | time.Time | 更新时间 |
| deleted_at | gorm.DeletedAt | 删除时间（软删除） |

### Menu（菜单模型）

| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | uint | 菜单ID |
| name | string | 菜单名称 |
| title | string | 菜单标题 |
| path | string | 路由路径 |
| component | string | 组件路径 |
| redirect | string | 重定向路径 |
| parent_id | uint | 父级菜单ID |
| parent | Menu | 父级菜单 |
| children | []Menu | 子菜单列表 |
| icon | string | 图标 |
| sort | int | 排序 |
| is_hidden | bool | 是否隐藏 |
| is_link | bool | 是否外部链接 |
| link_url | string | 外部链接地址 |
| status | int | 状态(1:正常, 0:禁用) |
| roles | []Role | 关联角色列表 |
| created_at | time.Time | 创建时间 |
| updated_at | time.Time | 更新时间 |
| deleted_at | gorm.DeletedAt | 删除时间（软删除） |

### OperationLog（操作日志模型）

| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | uint | 日志ID |
| operation | string | 操作描述 |
| user_id | uint | 用户ID |
| username | string | 用户名 |
| ip | string | IP地址 |
| user_agent | string | 用户代理 |
| status | int | 状态(1:成功, 0:失败) |
| request_method | string | 请求方法 |
| request_path | string | 请求路径 |
| request_body | string | 请求体 |
| response_body | string | 响应体 |
| response_time | int64 | 响应时间(毫秒) |
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

---

## 测试建议

推荐使用以下工具进行 API 测试：

- **Postman**: 功能强大的 API 测试工具
- **Insomnia**: 轻量级 API 测试工具
- **curl**: 命令行工具
- **HTTPie**: 友好的命令行 HTTP 客户端

### Postman 测试示例

1. 导入接口集合
2. 设置环境变量（如 `base_url`, `token`）
3. 在 Pre-request Script 中自动设置 Token：
```javascript
// 如果有 token，自动添加到请求头
if (pm.environment.get("token")) {
    pm.request.headers.add({
        key: "Authorization",
        value: "Bearer " + pm.environment.get("token")
    });
}
```

---

## 权限管理接口

### 1. 为角色分配菜单权限

**接口地址**: `POST /roles/:id/menus`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 角色ID |

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| menu_ids | array | 是 | 菜单ID列表 |

**请求示例**:
```json
{
  "menu_ids": [1, 2, 3]
}
```

**成功响应**:
```json
{
  "message": "菜单权限分配成功"
}
```

---

### 2. 获取角色的菜单权限

**接口地址**: `GET /roles/:id/menus`

**请求头**:
```
Authorization: Bearer <token>
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 角色ID |

**请求示例**:
```
GET /roles/1/menus
```

**成功响应**:
```json
{
  "message": "获取成功",
  "data": [
    {
      "id": 1,
      "name": "system",
      "title": "系统管理",
      "path": "/system",
      "icon": "setting",
      "sort": 1,
      "status": 1,
      "created_at": "2026-01-18T10:00:00Z",
      "updated_at": "2026-01-18T10:00:00Z"
    },
    {
      "id": 2,
      "name": "user",
      "title": "用户管理",
      "path": "/system/user",
      "component": "system/user/index",
      "parent_id": 1,
      "sort": 1,
      "status": 1,
      "created_at": "2026-01-18T10:00:00Z",
      "updated_at": "2026-01-18T10:00:00Z"
    }
  ]
}
```

---

### 3. 移除角色的菜单权限

**接口地址**: `DELETE /roles/:id/menus`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 角色ID |

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| menu_ids | array | 是 | 要移除的菜单ID列表 |

**请求示例**:
```json
{
  "menu_ids": [2, 3]
}
```

**成功响应**:
```json
{
  "message": "菜单权限移除成功"
}
```

---

### 4. 添加 Casbin 策略

**接口地址**: `POST /roles/:id/policies`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 角色ID |

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| path | string | 是 | API 路径（如 `/api/v1/users`） |
| method | string | 是 | HTTP 方法（GET、POST、PUT、DELETE） |

**请求示例**:
```json
{
  "path": "/api/v1/users",
  "method": "GET"
}
```

**成功响应**:
```json
{
  "message": "策略添加成功"
}
```

**错误响应**:
```json
{
  "error": "策略已存在"
}
```

---

### 5. 移除 Casbin 策略

**接口地址**: `DELETE /roles/:id/policies`

**请求头**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 角色ID |

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| path | string | 是 | API 路径 |
| method | string | 是 | HTTP 方法 |

**请求示例**:
```json
{
  "path": "/api/v1/users",
  "method": "GET"
}
```

**成功响应**:
```json
{
  "message": "策略移除成功"
}
```

---

### 6. 获取角色的 Casbin 策略

**接口地址**: `GET /roles/:id/policies`

**请求头**:
```
Authorization: Bearer <token>
```

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 角色ID |

**请求示例**:
```
GET /roles/1/policies
```

**成功响应**:
```json
{
  "message": "获取成功",
  "data": [
    {
      "sub": "role_1",
      "obj": "/api/v1/users",
      "act": "GET"
    },
    {
      "sub": "role_1",
      "obj": "/api/v1/roles",
      "act": "GET"
    }
  ]
}
```

---

### 7. 获取所有 Casbin 策略

**接口地址**: `GET /policies`

**请求头**:
```
Authorization: Bearer <token>
```

**请求参数**: 无

**请求示例**:
```
GET /policies
```

**成功响应**:
```json
{
  "message": "获取成功",
  "data": [
    {
      "sub": "role_1",
      "obj": "/api/v1/users",
      "act": "GET"
    },
    {
      "sub": "role_1",
      "obj": "/api/v1/users",
      "act": "POST"
    },
    {
      "sub": "role_2",
      "obj": "/api/v1/menus",
      "act": "GET"
    }
  ]
}
```

---

## 更新日志

### v1.0.0 (2026-01-18)

- 初始版本发布
- 实现用户管理、角色管理、菜单管理、操作日志功能
- 支持 JWT 认证和 RBAC 权限控制
- 提供完整的 RESTful API 接口
- 新增权限管理接口（菜单权限分配、Casbin 策略管理）