# API 测试脚本

这是一个用于测试后台管理系统 API 的 Go 脚本集合，基于最新的API接口文档实现。

## 项目结构

```
ceshi/
├── main.go              # 主程序入口
├── config.go            # 配置常量
├── models.go            # 数据模型定义
├── utils.go             # 工具函数
├── auth_func.go         # 认证相关测试（验证码、登录、登出）
├── user_func.go         # 用户管理测试
├── role_func.go         # 角色管理测试
├── menu_func.go         # 菜单管理测试
├── permission_func.go   # 权限管理测试
├── role_menu_func.go    # 角色菜单关联测试
├── log_func.go          # 操作日志测试
└── readme.md            # 说明文档
```

## 使用方法

### 1. 确保后端服务运行
首先启动后台管理系统：
```bash
cd ../backend
go run main.go
```

### 2. 运行测试脚本
```bash
cd ceshi
go run .
```

或者编译后运行：
```bash
go build -o test.exe
./test.exe
```

## 测试功能全覆盖

### 🔐 认证测试
- ✅ 验证码生成测试
- ✅ 用户登录测试
- ✅ 用户登出测试

### 👤 用户管理测试
- ✅ 创建用户
- ✅ 获取用户列表
- ✅ 获取用户详情
- ✅ 获取当前用户信息
- ✅ 更新用户信息
- ✅ 更新用户状态
- ✅ 修改密码
- ✅ 重置密码
- ✅ 获取用户角色列表
- ✅ 为用户分配角色
- ✅ 移除用户角色
- ✅ 删除用户

### 👥 角色管理测试
- ✅ 创建角色
- ✅ 获取角色列表
- ✅ 获取角色详情
- ✅ 更新角色
- ✅ 删除角色

### 📋 菜单管理测试
- ✅ 创建菜单
- ✅ 获取用户菜单
- ✅ 获取所有菜单
- ✅ 更新菜单
- ✅ 删除菜单

### 🔐 权限管理测试
- ✅ 创建权限资源
- ✅ 获取权限资源列表
- ✅ 获取所有权限资源
- ✅ 获取权限资源详情
- ✅ 更新权限资源
- ✅ 更新权限状态
- ✅ 删除权限资源

### 🔑 角色菜单关联测试
- ✅ 为角色分配菜单权限
- ✅ 获取角色菜单权限
- ✅ 移除角色菜单权限

### ⚖️ 权限策略管理测试
- ✅ 为角色添加Casbin策略
- ✅ 获取角色Casbin策略
- ✅ 移除角色Casbin策略

### 📝 操作日志测试
- ✅ 获取操作日志列表
- ✅ 删除操作日志

## 配置说明

在 `config.go` 中可以修改测试相关的配置：

```go
const (
    TestUsername = "testuser"     // 测试用户名
    TestPassword = "testpass123"  // 测试密码
    TestEmail    = "test@example.com" // 测试邮箱
    TestPhone    = "13800138000"  // 测试手机号
    
    AdminUsername = "admin"       // 管理员用户名
    AdminPassword = "123456"      // 管理员密码
)
```

## 测试结果说明

测试结果会在控制台显示：
- ✓ 表示测试通过
- ✗ 表示测试失败

最后会显示测试汇总：
```
==========================================================
测试结果汇总
==========================================================
✓ 登录测试: 登录成功
✓ 创建用户: 用户创建成功
✗ 获取不存在的用户: 获取失败，状态码: 404
==========================================================
总计: 25 | 通过: 23 | 失败: 2
🎉 所有测试通过!
```

## 开发说明

### 添加新的测试函数

1. 在对应的 `_func.go` 文件中添加测试函数
2. 函数名以 `test` 开头
3. 在 `main.go` 的 `runTests()` 函数中调用新函数
4. 使用 `addTestResult()` 记录测试结果

### 测试函数模板

```go
func testYourFunction() {
    fmt.Println("\n--- 测试你的功能 ---")
    
    // 准备测试数据
    testData := map[string]interface{}{
        "field": "value",
    }
    
    // 发送请求
    resp, body, err := sendAuthenticatedPost(BaseURL+"/your-endpoint", testData)
    
    if err != nil {
        addTestResult("你的测试", false, fmt.Sprintf("请求失败: %v", err))
        return
    }
    defer resp.Body.Close()
    
    // 验证结果
    if resp.StatusCode == http.StatusOK {
        addTestResult("你的测试", true, "测试成功")
    } else {
        addTestResult("你的测试", false, fmt.Sprintf("测试失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
    }
}
```

## 工具函数说明

### HTTP请求函数
- `sendGet(url)` - 发送GET请求
- `sendPost(url, data)` - 发送POST请求
- `sendAuthenticatedGet(url)` - 发送带认证的GET请求
- `sendAuthenticatedPost(url, data)` - 发送带认证的POST请求
- `sendAuthenticatedPut(url, data)` - 发送带认证的PUT请求
- `sendAuthenticatedDelete(url)` - 发送带认证的DELETE请求

### 辅助函数
- `prettyPrintJSON(data)` - 格式化JSON输出
- `addTestResult(name, passed, message)` - 记录测试结果

## 注意事项

1. **服务依赖**: 确保后端服务正常运行且端口配置正确（默认8080）
2. **认证要求**: 大部分测试需要管理员权限，请确保使用正确的认证信息
3. **测试数据**: 测试会在系统中产生真实记录，建议在测试环境中运行
4. **接口版本**: 测试脚本针对 `/api/v1` 版本的API设计
5. **响应格式**: 脚本已适配最新的统一响应格式（success/message/data/error）

## 常见问题

### Q: 测试显示连接失败
A: 请确认后端服务已启动，并检查端口配置是否正确

### Q: 认证相关测试失败
A: 请检查管理员账户凭据是否正确，确保用户存在且密码正确

### Q: 权限相关测试失败
A: 确认测试用户具有相应权限，或检查Casbin策略配置

### Q: 如何只运行特定测试
A: 可以注释掉 `main.go` 中 `runTests()` 函数里不想运行的测试调用

## 依赖关系

测试脚本通过 Go modules 管理依赖，会自动替换为本地 backend 模块：
```go
replace backend => ../backend
```

## 更新日志

### v2.0 (2026-02-04)
- ✨ 基于最新API文档重新设计测试脚本
- ✨ 实现全部接口的测试覆盖
- ✨ 统一响应格式处理
- ✨ 改进错误处理和测试报告
- ✨ 添加详细的测试分类和说明

### v1.0 (历史版本)
- 基础测试功能实现