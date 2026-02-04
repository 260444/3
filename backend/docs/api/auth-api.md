# 认证接口文档

## 基本信息

- **基础URL**: `http://localhost:8080/api/v1`
- **认证方式**: JWT Token
- **数据格式**: JSON
- **字符编码**: UTF-8

## 通用响应格式

### 成功响应

```json
{
  "success": true,
  "message": "操作成功",
  "data": { ... }
}
```

### 错误响应

```json
{
  "success": false,
  "message": "错误信息",
  "error": "错误信息"
}
```

### 业务错误响应

```json
{
  "success": false,
  "message": "业务错误描述",
  "error": "业务错误描述",
  "code": 40001
}
```

### 列表响应

```json
{
  "success": true,
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
  "success": true,
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
  "success": false,
  "message": "用户名或密码错误",
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
  "success": true,
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
  "success": true,
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
  "success": false,
  "message": "验证码生成失败",
  "error": "验证码生成失败"
}
```

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

## 常见业务错误说明

| 错误信息 | 说明 |
|----------|------|
| 用户名已存在 | 注册时用户名重复 |
| 邮箱已存在 | 注册时邮箱重复 |
| 用户名或密码错误 | 登录失败 |
| Token无效或已过期 | 认证失败 |
| 权限不足 | 无权限访问 |
| 资源不存在 | 记录不存在 |
| 数据验证失败 | 参数验证失败 |
| 验证码生成失败 | 验证码服务异常 |
| 用户未认证 | 未提供有效的JWT token |

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