# 操作日志接口文档

## 基本信息

- **基础URL**: `http://localhost:8080/api/v1`
- **认证方式**: JWT Token
- **数据格式**: JSON
- **字符编码**: UTF-8

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