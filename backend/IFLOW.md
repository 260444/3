# iFlow 项目上下文文件

## 项目概述

这是一个基于 Go 语言开发的企业级后台管理系统后端，采用前后端分离架构。系统提供了完整的用户管理、角色管理、权限管理、菜单管理和操作日志功能，使用 RBAC（基于角色的访问控制）模型进行权限管理。

### 核心技术栈

- **语言**: Go 1.24.12
- **Web 框架**: Gin
- **ORM**: GORM
- **数据库**: MySQL
- **缓存**: Redis
- **权限管理**: Casbin
- **日志**: Zap
- **配置管理**: Viper
- **认证**: JWT (golang-jwt/jwt/v5)
- **验证码**: base64Captcha

### 项目架构

采用经典的分层架构设计：

```
├── api/              # API 层
│   ├── handler/      # HTTP 处理器（处理请求和响应）
│   ├── middleware/   # 中间件（JWT 认证、CORS、日志）
│   └── router/       # 路由定义
├── config/           # 配置管理
├── internal/         # 内部包，不可被外部引用
│   ├── model/        # 数据模型定义
│   ├── repository/   # 数据访问层
│   └── service/      # 业务逻辑层
├── pkg/              # 可被外部引用的公共包
│   ├── cache/        # 缓存相关
│   ├── casbin/       # 权限管理
│   ├── database/     # 数据库初始化
│   ├── logger/       # 日志配置
│   ├── redis/        # Redis 配置
│   └── utils/        # 工具函数（JWT、验证码）
├── docs/             # 项目文档
├── logs/             # 日志文件
├── main.go           # 应用入口
├── go.mod            # Go 模块定义
└── go.sum            # 依赖锁定文件
```

## 构建和运行

### 前置条件

- Go 1.24.0 或更高版本
- MySQL 5.7 或更高版本
- Redis 5.0 或更高版本（可选）

### 配置文件

配置文件位于 `config/config.yaml`，包含以下配置项：

```yaml
server:
  port: ":8080"           # 服务器端口
  mode: "debug"           # 运行模式: debug/release
  read_timeout: 60        # 读取超时（秒）
  write_timeout: 60       # 写入超时（秒）

database:
  host: "localhost"       # 数据库主机
  port: 3306              # 数据库端口
  user: "root"            # 数据库用户
  password: "123456"      # 数据库密码
  dbname: "admin_system"  # 数据库名称
  charset: "utf8mb4"      # 字符集

redis:
  addr: "localhost:6379"  # Redis 地址
  password: ""            # Redis 密码
  db: 0                   # Redis 数据库

jwt:
  secret: "your-secret-key"  # JWT 密钥
  timeout: 7200              # Token 过期时间（秒）
```

### 构建命令

```bash
# 安装依赖
go mod tidy

# 构建
go build -o backend main.go

# 运行
./backend

# 或直接运行
go run main.go
```

### 开发模式

```bash
# 直接运行开发服务器
go run main.go
```

### 生产模式

```bash
# 构建
go build -ldflags="-s -w" -o backend main.go

# 使用 systemd 启动（需要创建服务文件）
sudo systemctl start golang-backend
```

## 开发规范

### 命名约定

- **包名**: 小写字母，简洁明了（如 `config`, `model`, `service`）
- **结构体名**: 驼峰命名，导出结构体首字母大写
- **接口名**: 通常以 `er` 结尾（如 `UserService`）
- **函数名**: 导出函数首字母大写，私有函数首字母小写
- **变量名**: 驼峰命名，首字母小写

### 代码格式

- 使用 `go fmt` 自动格式化代码
- 每行代码长度不超过 120 个字符

### 分层架构职责

- **Handler 层**: 处理 HTTP 请求和响应，参数验证，调用 Service 层
- **Service 层**: 业务逻辑处理，事务管理，调用 Repository 层
- **Repository 层**: 数据访问操作，数据库 CRUD
- **Model 层**: 数据结构定义，GORM 标签
- **中间件层**: 请求处理前后的中间件（认证、日志、CORS）

### 数据库规范

- **表名**: 使用复数形式（如 `users`, `roles`, `menus`）
- **字段名**: 小写蛇形命名（如 `user_name`, `created_at`）
- **主键**: 使用 `uint` 类型，`gorm:"primaryKey"`
- **时间戳**: `time.Time` 类型，GORM 自动处理
- **软删除**: 使用 `gorm.DeletedAt` 类型

### API 设计规范

- 使用 RESTful 风格
- 路由使用名词复数形式：`/api/v1/users`, `/api/v1/roles`
- 使用 HTTP 动词表示操作：GET、POST、PUT、DELETE
- 响应格式统一：
  ```json
  {
    "message": "操作成功",
    "data": { ... }
  }
  ```
- 错误响应：
  ```json
  {
    "error": "错误信息"
  }
  ```

### 错误处理

- 错误信息应清晰明了
- 使用 `errors.New()` 或 `fmt.Errorf()` 创建错误
- 统一错误处理格式

### 日志规范

使用 Zap 日志库，输出 JSON 格式日志：

```go
logger.Logger.Info("用户登录",
    zap.String("username", username),
    zap.String("ip", ip),
)
```

日志级别：
- DEBUG: 调试信息
- INFO: 一般信息
- WARN: 警告信息
- ERROR: 错误信息
- FATAL: 致命错误

### Git 提交规范

- `feat`: 新功能
- `fix`: 修复 bug
- `docs`: 文档更新
- `style`: 代码格式调整
- `refactor`: 代码重构
- `test`: 测试相关
- `chore`: 构建过程或辅助工具的变动

示例：
```
feat(user): 添加用户注册功能

- 实现用户注册接口
- 添加密码加密逻辑
- 完善错误处理机制
```

## 核心功能模块

### 1. 用户管理

- 用户注册、登录
- 用户信息 CRUD
- 用户状态管理（启用/禁用）
- 密码修改
- 用户与角色关联

### 2. 角色管理

- 角色创建、编辑、删除
- 角色权限分配
- 角色与菜单关联（多对多）

### 3. 菜单管理

- 菜单树形结构
- 菜单 CRUD
- 菜单层级管理
- 菜单与角色关联
- 支持隐藏菜单、外部链接

### 4. 权限管理

- 基于 Casbin 的 RBAC 权限控制
- JWT 认证
- 中间件权限验证
- 细粒度权限控制

### 5. 操作日志

- 记录用户操作
- 日志查询和删除
- 记录请求和响应信息

### 6. 验证码

- 图形验证码生成
- 登录验证码验证

## 数据模型

### User（用户模型）

```go
type User struct {
    ID           uint       `gorm:"primaryKey" json:"id"`
    Username     string     `gorm:"uniqueIndex;size:50;not null" json:"username"`
    Password     string     `gorm:"size:255;not null" json:"password"`
    Email        string     `gorm:"uniqueIndex;size:100" json:"email"`
    Phone        string     `gorm:"uniqueIndex;size:20" json:"phone"`
    Nickname     string     `gorm:"size:50" json:"nickname"`
    Avatar       string     `gorm:"size:255" json:"avatar"`
    Status       int        `gorm:"default:1" json:"status"`
    RoleID       uint       `json:"role_id"`
    Role         Role       `gorm:"foreignKey:RoleID" json:"role"`
    // ... 其他字段
}
```

### Role（角色模型）

```go
type Role struct {
    ID          uint   `gorm:"primaryKey" json:"id"`
    Name        string `gorm:"uniqueIndex;size:50;not null" json:"name"`
    Description string `gorm:"size:255" json:"description"`
    Status      int    `gorm:"default:1" json:"status"`
    Users       []User `gorm:"foreignKey:RoleID" json:"users"`
    Menus       []Menu `gorm:"many2many:role_menus;" json:"menus"`
    // ... 其他字段
}
```

### Menu（菜单模型）

```go
type Menu struct {
    ID        uint    `gorm:"primaryKey" json:"id"`
    Name      string  `gorm:"size:50;not null" json:"name"`
    Title     string  `gorm:"size:50;not null" json:"title"`
    Path      string  `gorm:"size:100" json:"path"`
    Component string  `gorm:"size:100" json:"component"`
    ParentID  *uint   `json:"parent_id"`
    Parent    *Menu   `gorm:"foreignKey:ParentID" json:"parent"`
    Children  []Menu  `gorm:"foreignKey:ParentID" json:"children"`
    Icon      string  `gorm:"size:50" json:"icon"`
    Sort      int     `gorm:"default:0" json:"sort"`
    IsHidden  bool    `gorm:"default:false" json:"is_hidden"`
    Status    int     `gorm:"default:1" json:"status"`
    // ... 其他字段
}
```

## API 路由

### 公开路由（无需认证）

- `POST /api/v1/login` - 用户登录
- `POST /api/v1/register` - 用户注册
- `GET /api/v1/captcha` - 获取验证码

### 受保护路由（需要 JWT 认证）

#### 用户管理

- `GET /api/v1/users` - 获取用户列表
- `GET /api/v1/users/:id` - 获取用户详情
- `PUT /api/v1/users/:id` - 更新用户
- `PUT /api/v1/users/:id/status` - 更新用户状态
- `DELETE /api/v1/users/:id` - 删除用户
- `PUT /api/v1/users/change-password` - 修改密码

#### 角色管理

- `POST /api/v1/roles` - 创建角色
- `GET /api/v1/roles` - 获取角色列表
- `GET /api/v1/roles/:id` - 获取角色详情
- `PUT /api/v1/roles/:id` - 更新角色
- `DELETE /api/v1/roles/:id` - 删除角色

#### 菜单管理

- `POST /api/v1/menus` - 创建菜单
- `GET /api/v1/menus` - 获取菜单树
- `GET /api/v1/menus/all` - 获取所有菜单
- `GET /api/v1/menus/:id` - 获取菜单详情
- `PUT /api/v1/menus/:id` - 更新菜单
- `DELETE /api/v1/menus/:id` - 删除菜单

#### 操作日志

- `GET /api/v1/operation-logs` - 获取操作日志列表
- `DELETE /api/v1/operation-logs/:id` - 删除操作日志

## 安全规范

### 输入验证

- 对所有用户输入进行验证
- 防止 SQL 注入（使用 GORM 参数化查询）
- 防止 XSS 攻击

### 认证和授权

- 使用 JWT 进行身份认证
- 实现细粒度权限控制（Casbin）
- 定期更新密码策略
- 密码使用 bcrypt 加密存储

### 数据安全

- 敏感数据加密存储
- 使用 HTTPS 传输（生产环境）
- 实施访问日志记录

## 测试

### 单元测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./internal/service/

# 查看测试覆盖率
go test -cover ./...
```

### API 测试

使用 Postman 或类似工具测试 API 接口。参考 `docs/API文档.md` 获取详细的接口说明。

## 部署

详细部署说明请参考 `docs/部署文档.md`。

### 生产环境检查清单

- [ ] 修改配置文件中的敏感信息（数据库密码、JWT 密钥等）
- [ ] 设置 Gin 模式为 `release`
- [ ] 配置 HTTPS
- [ ] 设置防火墙规则
- [ ] 配置日志轮转
- [ ] 设置数据库备份策略
- [ ] 配置监控和告警

## 文档

项目文档位于 `docs/` 目录：

- `使用说明.md` - 系统使用指南
- `开发规范.md` - 编码规范和最佳实践
- `部署文档.md` - 部署和运维指南
- `系统架构.md` - 系统架构设计
- `API文档.md` - API 接口文档

## 常见问题

### 服务无法启动

1. 检查配置文件是否正确
2. 确认数据库和 Redis 是否正常运行
3. 检查端口是否被占用
4. 查看日志文件 `logs/app.log`

### 数据库连接失败

1. 确认 MySQL 服务是否运行
2. 检查数据库配置是否正确
3. 确认数据库用户权限
4. 检查防火墙设置

### Redis 连接失败

- Redis 是可选的，连接失败不会影响系统运行
- 检查 Redis 配置是否正确
- 确认 Redis 服务是否运行

## 项目状态

- **当前版本**: 1.0.0
- **Go 版本**: 1.24.12
- **开发状态**: 活跃开发中

## 联系方式

如有问题或建议，请联系开发团队。

---

**注意**: 请勿在代码中硬编码敏感信息（如密码、API 密钥等）。所有敏感配置应通过配置文件或环境变量管理。