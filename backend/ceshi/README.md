# API 测试脚本

本目录包含后端 API 的测试脚本，支持 Python 和 Go 两种语言。

## 文件说明

- `test_api.py` - Python 版本的 API 测试脚本
- `test_api_README.md` - Python 测试脚本使用说明
- `test_api.go` - Go 版本的 API 测试脚本（最新）
- `test_api_README_GO.md` - Go 测试脚本使用说明
- `README.md` - 本文件

## 快速开始

### 1. 启动后端服务

```bash
# 在项目根目录执行
go run main.go
```

### 2. 导入测试数据

```bash
# 使用 MySQL 客户端导入测试数据
mysql -u root -p admin_system < docs/test_data.sql
```

### 3. 运行测试

#### Go 版本（推荐）

```bash
cd ceshi
go run test_api.go
```

#### Python 版本

```bash
cd ceshi
python test_api.py
```

## 测试覆盖范围

脚本测试以下 API 接口：

### 公开接口
- ✓ 获取验证码
- ✓ 用户登录
- ✓ 退出登录

### 用户管理
- ✓ 创建用户
- ✓ 获取用户列表
- ✓ 获取用户详情
- ✓ 更新用户信息
- ✓ 更新用户状态
- ✓ 修改密码

### 角色管理
- ✓ 创建角色
- ✓ 获取角色列表
- ✓ 获取角色详情
- ✓ 更新角色

### 菜单管理
- ✓ 创建菜单
- ✓ 获取菜单树
- ✓ 获取所有菜单
- ✓ 获取菜单详情
- ✓ 更新菜单

### 权限管理
- ✓ 为角色分配菜单权限
- ✓ 获取角色菜单权限
- ✓ 添加 Casbin 策略
- ✓ 获取角色 Casbin 策略
- ✓ 获取所有 Casbin 策略
- ✓ 移除 Casbin 策略
- ✓ 移除角色菜单权限

### 操作日志
- ✓ 获取操作日志列表

## 默认测试账号

- **用户名**: `admin`
- **密码**: `123456`

## 注意事项

1. 运行测试前请确保已导入测试数据
2. 确保后端服务正在运行
3. 测试脚本会创建测试数据，但默认不会自动清理
4. 如需修改服务器地址，请编辑测试脚本中的配置

## 更多信息

- Go 版本详细说明: 查看 `test_api_README_GO.md`
- Python 版本详细说明: 查看 `test_api_README.md`
- API 接口文档: 查看 `docs/API文档.md`