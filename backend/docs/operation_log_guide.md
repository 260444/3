# 操作日志使用指南

## 功能概述

操作日志系统用于记录用户的操作行为，提供完整的操作审计功能，包括：
- 自动记录用户操作
- 多维度查询和筛选
- 统计分析功能
- 敏感信息脱敏处理
- 批量管理功能

## 核心特性

### 1. 自动记录机制
系统通过中间件自动记录以下操作：
- **创建操作** (POST): 用户创建、角色创建等
- **更新操作** (PUT/PATCH): 信息修改、状态变更等  
- **删除操作** (DELETE): 数据删除操作
- **敏感查询** (GET): 个人资料查询等

### 2. 记录内容
每条操作日志包含以下信息：

| 字段 | 说明 | 示例 |
|------|------|------|
| operation | 操作描述 | "用户管理创建" |
| module | 模块分类 | "用户管理" |
| level | 日志级别 | 1(普通)/2(重要)/3(危险) |
| user_id | 操作用户ID | 1 |
| username | 操作用户名 | "admin" |
| ip | 客户端IP | "192.168.1.100" |
| user_agent | 用户代理 | "Mozilla/5.0..." |
| status | 操作状态 | 1(成功)/0(失败) |
| request_method | 请求方法 | "POST" |
| request_path | 请求路径 | "/api/v1/users" |
| request_body | 请求体 | "{...}" |
| response_body | 响应体 | "{...}" |
| response_time | 响应时间 | 156 |
| refer_id | 关联业务ID | 123 |

### 3. 敏感信息保护
自动对以下敏感字段进行脱敏处理：
- password → ***
- old_password → ***
- new_password → ***
- token → ***

## API 接口

### 1. 查询操作日志列表
```
GET /api/v1/operation-logs
```

**查询参数:**
- `page`: 页码 (默认: 1)
- `page_size`: 每页数量 (默认: 10, 最大: 100)
- `username`: 操作用户名筛选
- `module`: 模块筛选 (用户管理/角色管理/菜单管理/权限管理/认证管理/日志管理)
- `status`: 状态筛选 (1:成功, 0:失败)
- `start_time`: 开始时间 (格式: 2026-01-01)
- `end_time`: 结束时间 (格式: 2026-12-31)

**响应示例:**
```json
{
  "message": "获取成功",
  "data": {
    "list": [
      {
        "id": 1,
        "operation": "用户管理创建",
        "module": "用户管理",
        "level": 2,
        "user_id": 1,
        "username": "admin",
        "ip": "192.168.1.100",
        "user_agent": "Mozilla/5.0...",
        "status": 1,
        "request_method": "POST",
        "request_path": "/api/v1/users",
        "request_body": "{\"username\":\"test\",\"password\":\"***\"}",
        "response_body": "{\"message\":\"创建成功\"}",
        "response_time": 156,
        "refer_id": null,
        "created_at": "2026-01-18T19:07:49Z"
      }
    ],
    "total": 100,
    "page": 1,
    "page_size": 10
  }
}
```

### 2. 删除单条操作日志
```
DELETE /api/v1/operation-logs/{id}
```

### 3. 批量删除操作日志
```
POST /api/v1/operation-logs/batch-delete
```

**请求体:**
```json
{
  "ids": [1, 2, 3, 4, 5]
}
```

### 4. 获取操作日志统计信息
```
GET /api/v1/operation-logs/stats
```

**响应示例:**
```json
{
  "message": "获取统计信息成功",
  "data": {
    "total_count": 1000,
    "success_count": 950,
    "failure_count": 50,
    "user_count": 15,
    "module_stats": [
      {"module": "用户管理", "count": 300},
      {"module": "角色管理", "count": 200},
      {"module": "菜单管理", "count": 150}
    ],
    "recent_activity": [
      {"date": "2026-02-03", "count": 50},
      {"date": "2026-02-02", "count": 45},
      {"date": "2026-02-01", "count": 40}
    ]
  }
}
```

## 使用场景

### 1. 安全审计
- 追踪敏感操作（删除、权限变更等）
- 监控异常登录行为
- 审计数据变更历史

### 2. 问题排查
- 查看操作失败原因
- 分析系统性能瓶颈
- 还原操作现场

### 3. 合规要求
- 满足数据保护法规要求
- 提供操作凭证
- 支持内部审查

## 配置说明

### 中间件配置
在 `main.go` 中启用操作日志中间件：

```go
// 在需要记录操作的路由组中使用
protected.Use(middleware.OperationLogMiddleware(operationLogService))
```

### 日志级别说明
- **级别1 (普通)**: 一般查询操作
- **级别2 (重要)**: 创建、更新操作
- **级别3 (危险)**: 删除操作、敏感操作

### 不记录的操作
- GET 查询操作（除非是敏感查询）
- 健康检查接口
- 系统监控接口

## 最佳实践

1. **定期清理**: 建议保留3-6个月的操作日志
2. **权限控制**: 操作日志查看应限制权限
3. **性能考虑**: 大量日志时注意分页查询
4. **备份策略**: 重要操作日志建议备份存储
5. **监控告警**: 对异常操作设置告警机制

## 常见问题

### Q: 如何排除某些接口不记录日志？
A: 在 `middleware/operation_log_middleware.go` 的 `shouldLogOperation` 函数中添加排除规则。

### Q: 敏感信息如何完全避免记录？
A: 可以在中间件中对特定路径的请求体直接置空处理。

### Q: 日志量太大怎么办？
A: 可以调整记录策略，只记录重要操作，或者实施日志轮转和归档。