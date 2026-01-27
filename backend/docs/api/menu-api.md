# 菜单管理接口文档

## 基本信息

- **基础URL**: `http://localhost:8080/api/v1`
- **认证方式**: JWT Token
- **数据格式**: JSON
- **字符编码**: UTF-8

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
      "created_at": "2026-01-18 21:54:50",
      "updated_at": "2026-01-18 21:54:50",
      "deleted_at": null,
      "name": "dashboard",
      "title": "控制台",
      "path": "/dashboard",
      "component": "DashboardView",
      "redirect": "",
      "parent_id": 0,
      "parent": null,
      "children": [
        {
          "id": 0,
          "created_at": "",
          "updated_at": "",
          "deleted_at": null,
          "name": "",
          "title": "",
          "path": "",
          "component": "",
          "redirect": "",
          "parent_id": null,
          "parent": null,
          "children": null,
          "icon": "",
          "sort": 0,
          "is_hidden": false,
          "is_link": false,
          "link_url": "",
          "status": 0,
          "roles": null
        }
      ],
      "icon": "House",
      "sort": 1,
      "is_hidden": false,
      "is_link": false,
      "link_url": "",
      "status": 1,
      "roles": null
    },
    {
      "id": 2,
      "created_at": "2026-01-18 21:54:50",
      "updated_at": "2026-01-18 21:54:50",
      "deleted_at": null,
      "name": "system",
      "title": "系统管理",
      "path": "/system",
      "component": "",
      "redirect": "",
      "parent_id": 0,
      "parent": null,
      "children": [
        {
          "id": 11,
          "created_at": "2026-01-18 21:54:50",
          "updated_at": "2026-01-18 21:54:50",
          "deleted_at": null,
          "name": "users",
          "title": "用户管理",
          "path": "/users",
          "component": "UserManageView",
          "redirect": "",
          "parent_id": 2,
          "parent": null,
          "children": null,
          "icon": "User",
          "sort": 1,
          "is_hidden": false,
          "is_link": false,
          "link_url": "",
          "status": 1,
          "roles": null
        }
      ],
      "icon": "Setting",
      "sort": 100,
      "is_hidden": false,
      "is_link": false,
      "link_url": "",
      "status": 1,
      "roles": null
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

## 数据模型

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