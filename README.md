# Go语言后台管理系统

基于Go语言开发的后台管理系统，采用前后端分离的架构设计。系统提供了完整的用户管理、角色管理、权限管理、菜单管理等功能，是一个企业级的后台管理解决方案。

## 技术栈

### 后端技术栈
- 语言: Go 1.24.0
- Web框架: Gin
- 数据库: MySQL + GORM
- 缓存: Redis
- 权限管理: Casbin
- 认证: JWT
- 配置管理: Viper
- 日志: Zap + Lumberjack
- API文档: Swagger
- 验证码: Base64Captcha

### 前端技术栈
- 框架: Vue 3 + TypeScript
- UI组件库: Element Plus
- 构建工具: Vite
- 路由: Vue Router
- 状态管理: Pinia
- HTTP客户端: Axios

## 功能特性

1. **用户管理**
   - 用户注册、登录、注销
   - 用户信息管理
   - 密码修改
   - 用户状态管理

2. **角色管理**
   - 角色创建、编辑、删除
   - 角色权限分配
   - 角色用户关联

3. **权限管理**
   - 基于RBAC的权限控制
   - 细粒度权限管理
   - 动态权限加载

4. **菜单管理**
   - 动态菜单配置
   - 菜单权限控制
   - 多级菜单支持

5. **系统功能**
   - 操作日志记录
   - 系统监控
   - 数据统计

## 快速开始

### 后端启动

1. 安装Go环境（1.24.0或更高版本）
2. 安装MySQL和Redis
3. 克隆项目
4. 进入backend目录，安装依赖：
   ```bash
   go mod tidy
   ```
5. 修改配置文件 `config/config.yaml`，配置数据库连接信息
6. 启动服务：
   ```bash
   go run main.go
   ```

### 前端启动

1. 安装Node.js和npm
2. 进入frontend目录，安装依赖：
   ```bash
   npm install
   ```
3. 启动开发服务器：
   ```bash
   npm run dev
   ```

## 项目结构

```
golang后台管理项目/
├── backend/                 # 后端代码
│   ├── api/                # API层
│   │   ├── handler/        # 处理器
│   │   ├── middleware/     # 中间件
│   │   └── router/         # 路由
│   ├── config/             # 配置管理
│   ├── internal/           # 内部包
│   │   ├── model/          # 数据模型
│   │   ├── repository/     # 数据访问层
│   │   └── service/        # 业务逻辑层
│   ├── pkg/                # 公共包
│   │   ├── cache/          # 缓存
│   │   ├── casbin/         # 权限管理
│   │   ├── database/       # 数据库
│   │   ├── logger/         # 日志
│   │   ├── redis/          # Redis
│   │   └── utils/          # 工具函数
│   └── main.go             # 程序入口
├── frontend/               # 前端代码
│   ├── src/
│   │   ├── api/            # API接口
│   │   ├── components/     # 组件
│   │   ├── views/          # 页面
│   │   ├── router/         # 路由
│   │   └── utils/          # 工具函数
│   └── package.json
└── docs/                   # 项目文档
```

## 环境配置

### 后端配置

配置文件 `backend/config/config.yaml`:

```yaml
server:
  port: ":8080"
  mode: "debug"
  read_timeout: 60
  write_timeout: 60

database:
  host: "localhost"
  port: 3306
  user: "root"
  password: "password"
  dbname: "admin_system"
  charset: "utf8mb4"

redis:
  addr: "localhost:6379"
  password: ""
  db: 0

jwt:
  secret: "your-secret-key"
  timeout: 7200
```

### 前端配置

代理配置在 `frontend/vite.config.ts` 中，将API请求代理到后端服务。

## API接口

### 用户相关
- `POST /api/v1/login` - 用户登录
- `POST /api/v1/register` - 用户注册
- `GET /api/v1/users` - 获取用户列表
- `PUT /api/v1/users/:id` - 更新用户信息
- `DELETE /api/v1/users/:id` - 删除用户
- `PUT /api/v1/users/:id/status` - 更新用户状态
- `PUT /api/v1/users/change-password` - 修改密码

### 角色相关
- `POST /api/v1/roles` - 创建角色
- `GET /api/v1/roles` - 获取角色列表
- `GET /api/v1/roles/:id` - 获取角色详情
- `PUT /api/v1/roles/:id` - 更新角色
- `DELETE /api/v1/roles/:id` - 删除角色

### 菜单相关
- `POST /api/v1/menus` - 创建菜单
- `GET /api/v1/menus` - 获取菜单树
- `GET /api/v1/menus/all` - 获取所有菜单
- `GET /api/v1/menus/:id` - 获取菜单详情
- `PUT /api/v1/menus/:id` - 更新菜单
- `DELETE /api/v1/menus/:id` - 删除菜单

### 操作日志相关
- `GET /api/v1/operation-logs` - 获取操作日志列表
- `DELETE /api/v1/operation-logs/:id` - 删除操作日志

## 部署

### 后端部署

1. 在服务器上安装Go环境
2. 修改配置文件以适配服务器环境
3. 构建二进制文件：
   ```bash
   go build -o backend main.go
   ```
4. 启动服务：
   ```bash
   ./backend
   ```

### 前端部署

1. 构建生产版本：
   ```bash
   npm run build
   ```
2. 部署 `dist` 目录到Web服务器

## 安全说明

- 使用JWT进行身份认证
- 采用bcrypt对密码进行加密存储
- 使用Casbin进行权限控制
- 使用中间件防止常见Web攻击

## 开发说明

1. 严格遵循Go语言编码规范
2. 使用GORM进行数据库操作
3. 统一错误处理机制
4. 日志记录规范

## 许可证

[MIT LICENSE]

## 联系方式

如有问题或建议，请提交Issue或联系开发团队。