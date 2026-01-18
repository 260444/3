# API 测试脚本使用说明

## 概述

本项目提供了两个版本的 API 测试脚本，用于测试所有后端 API 接口：

1. **Python 版本** (`test_api.py`) - 推荐使用
2. **PowerShell 版本** (`test_api.ps1`) - Windows 原生支持

## 测试覆盖范围

脚本会测试以下 API 接口：

### 公开接口（无需认证）
- ✓ 获取验证码
- ✓ 用户注册
- ✓ 用户登录

### 用户管理（需要认证）
- ✓ 获取用户列表
- ✓ 获取用户信息
- ✓ 更新用户信息
- ✓ 更新用户状态
- ✓ 修改密码
- ✓ 删除用户

### 角色管理（需要认证）
- ✓ 创建角色
- ✓ 获取角色列表
- ✓ 获取角色详情
- ✓ 更新角色
- ✓ 删除角色

### 菜单管理（需要认证）
- ✓ 创建菜单
- ✓ 获取菜单树
- ✓ 获取所有菜单
- ✓ 获取菜单详情
- ✓ 更新菜单
- ✓ 删除菜单

### 操作日志（需要认证）
- ✓ 获取操作日志列表

## 前置条件

### 1. 启动后端服务

确保后端服务已启动并运行在 `http://localhost:8080`

```bash
# 方式 1: 直接运行
go run main.go

# 方式 2: 构建后运行
go build -o backend main.go
./backend
```

### 2. 确保数据库已初始化

确保 MySQL 数据库已创建并初始化相关表结构。

### 3. 安装依赖

#### Python 版本

```bash
# 安装 Python 3.6+
python --version

# 安装 requests 库
pip install requests
```

#### PowerShell 版本

无需额外安装，PowerShell 5.1+ 已内置所需功能。

## 使用方法

### Python 版本（推荐）

```bash
# 运行测试脚本
python test_api.py
```

### PowerShell 版本

```powershell
# 方式 1: 直接运行
.\test_api.ps1

# 方式 2: 如果遇到执行策略限制
powershell -ExecutionPolicy Bypass -File test_api.ps1
```

## 测试流程

脚本会按照以下顺序执行测试：

1. **检查服务器状态** - 验证后端服务是否正常运行
2. **公开接口测试** - 测试无需认证的接口
3. **认证接口测试** - 使用注册的用户登录后测试需要认证的接口
4. **清理测试数据** - 删除测试过程中创建的数据（用户、角色、菜单）

## 测试数据说明

脚本会自动创建以下测试数据：

- **测试用户**: `testuser_<timestamp>`
  - 密码: `Test123456`
  - 邮箱: `testuser_<timestamp>@example.com`
  - 昵称: 测试用户

- **测试角色**: `测试角色_<timestamp>`
  - 描述: 这是一个测试角色

- **测试菜单**: `test_menu_<timestamp>`
  - 标题: 测试菜单
  - 路径: /test

所有测试数据在测试完成后会被自动清理。

## 输出说明

### 成功输出示例

```
✓ 获取验证码 - PASS
  状态码: 200
  响应: {
    "message": "验证码生成成功",
    "data": {
      "id": "xxx",
      "image": "data:image/png;base64,..."
    }
  }
```

### 失败输出示例

```
✗ 用户登录 - FAIL
  状态码: 401
  响应: {
    "error": "用户名或密码错误"
  }
```

### 测试结果汇总

```
============================================================
测试结果汇总
============================================================
总计: 23 个测试
通过: 20 个
失败: 3 个
成功率: 87.0%
============================================================
```

## 配置说明

### 修改服务器地址

如果后端服务运行在其他地址或端口，可以修改脚本中的配置：

#### Python 版本

```python
# 修改 test_api.py 中的配置
BASE_URL = "http://your-server:port"
```

#### PowerShell 版本

```powershell
# 修改 test_api.ps1 中的配置
$BASE_URL = "http://your-server:port"
```

## 常见问题

### 1. 无法连接到服务器

**错误信息**: `✗ 无法连接到服务器 http://localhost:8080`

**解决方案**:
- 确认后端服务已启动
- 检查端口是否正确
- 检查防火墙设置

### 2. Python 版本提示缺少 requests 库

**错误信息**: `ModuleNotFoundError: No module named 'requests'`

**解决方案**:
```bash
pip install requests
```

### 3. PowerShell 版本提示执行策略限制

**错误信息**: `无法加载文件 test_api.ps1，因为在此系统上禁止运行脚本`

**解决方案**:
```powershell
powershell -ExecutionPolicy Bypass -File test_api.ps1
```

### 4. 测试失败率高

**可能原因**:
- 数据库未正确初始化
- 数据库连接失败
- JWT 密钥配置错误
- Redis 连接失败（可选）

**解决方案**:
- 检查 `config/config.yaml` 配置
- 查看后端日志 `logs/app.log`
- 确认数据库表结构正确

### 5. 用户名已存在

**错误信息**: 用户注册时提示用户名已存在

**解决方案**:
- 脚本使用时间戳生成唯一用户名，理论上不会重复
- 如果遇到此问题，可以手动删除数据库中的测试用户
- 或者修改脚本中的 `test_username` 生成逻辑

## 扩展测试

### 添加新的测试用例

#### Python 版本

```python
def test_your_new_api():
    """测试新的 API"""
    try:
        data = {
            "param1": "value1",
            "param2": "value2"
        }
        response = make_request("POST", "/your-endpoint", data=data, headers=get_auth_headers())
        if response.status_code == 200:
            print_result("新 API 测试", "PASS", response)
            return True
        else:
            print_result("新 API 测试", "FAIL", response)
            return False
    except Exception as e:
        print_result("新 API 测试", "FAIL", error=str(e))
        return False

# 在 main() 函数的 tests 列表中添加
tests = [
    # ... 其他测试
    ("新 API 测试", test_your_new_api),
]
```

#### PowerShell 版本

```powershell
function Test-YourNewApi {
    try {
        $data = @{
            param1 = "value1"
            param2 = "value2"
        }
        $result = Make-Request -method "POST" -endpoint "/your-endpoint" -data $data -headers (Get-AuthHeaders)
        if ($result.StatusCode -lt 400) {
            Print-Result "新 API 测试" "PASS"
            return $true
        } else {
            Print-Result "新 API 测试" "FAIL"
            return $false
        }
    } catch {
        Print-Result "新 API 测试" "FAIL" -error $_.Exception.Message
        return $false
    }
}

# 在 Main 函数的 $tests 数组中添加
$tests = @(
    # ... 其他测试
    @{ Name = "新 API 测试"; Test = { Test-YourNewApi } }
)
```

## 注意事项

1. **测试数据清理**: 脚本会自动清理测试数据，但如果测试中途失败，可能需要手动清理
2. **并发测试**: 脚本是顺序执行的，不支持并发测试
3. **性能测试**: 此脚本仅用于功能测试，不适用于性能测试
4. **生产环境**: 请勿在生产环境运行此测试脚本
5. **数据安全**: 测试脚本会创建和删除数据，请确保在测试环境中运行

## 技术支持

如遇到问题，请检查：
1. 后端服务日志: `logs/app.log`
2. 数据库连接状态
3. 网络连接状态
4. 配置文件: `config/config.yaml`

## 版本信息

- **当前版本**: 1.0.0
- **支持的后端版本**: 1.0.0+
- **最后更新**: 2026-01-17