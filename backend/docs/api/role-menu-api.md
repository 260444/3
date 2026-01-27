# 角色菜单关联接口文档

## 基本信息

- **基础URL**: `http://localhost:8080/api/v1`
- **认证方式**: JWT Token
- **数据格式**: JSON
- **字符编码**: UTF-8

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
  "message": "菜单权限分配成功",
  
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
  "message": "菜单权限分配成功",
  "data":[{"p_id":1,"m_id":null},{"p_id":2,"m_id":[11,12,14]}]
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

## 数据模型

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