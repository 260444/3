# iFlow 项目上下文文件

## 项目概述

这是一个全栈企业级后台管理系统，包含前后端分离架构的完整解决方案。系统提供了用户管理、角色管理、菜单管理、权限分配、操作日志以及资产管理系统等核心功能，使用 RBAC（基于角色的访问控制）模型进行权限管理。

**版本要求**：
- Go 1.24.12+
- Node.js 16+ (推荐 LTS 版本)
- MySQL 8.0+ (兼容 5.7+)
- Redis 5.0+ (推荐 7.2+)

### 技术栈

**后端技术栈**：
- **语言**: Go 1.24.12
- **Web 框架**: Gin v1.11.0
- **ORM**: GORM v1.31.1
- **数据库**: MySQL 8.0
- **缓存**: Redis 7.2
- **权限管理**: Casbin v3.8.1
- **日志**: Zap v1.27.1 + Lumberjack v2.2.1
- **配置管理**: Viper v1.21.0
- **认证**: JWT (golang-jwt/jwt/v5 v5.2.2)
- **验证码**: base64Captcha v1.3.8

**前端技术栈**：
- **框架**: Vue 3.4.21 (Composition API)
- **语言**: TypeScript 5.2.2
- **构建工具**: Vite 5.2.0
- **UI 组件库**: Element Plus 2.6.1
- **图标库**: @element-plus/icons-vue 2.3.1
- **路由**: Vue Router 4.3.0
- **状态管理**: Pinia 2.1.7
- **HTTP 客户端**: Axios 1.6.8

### 项目架构

项目采用前后端分离架构，包含以下目录结构：

```
├── backend/          # 后端项目
│   ├── api/          # API 层
│   │   ├── handler/  # HTTP 处理器
│   │   ├── middleware/ # 中间件
│   │   └── router/   # 路由定义
│   ├── config/       # 配置管理
│   ├── internal/     # 内部包
│   │   ├── model/    # 数据模型
│   │   ├── repository/ # 数据访问层
│   │   └── service/  # 业务逻辑层
│   ├── pkg/          # 可被外部引用的公共包
│   ├── sql/          # SQL脚本
│   ├── docs/         # 项目文档
│   ├── logs/         # 日志文件
│   ├── main.go       # 应用入口
│   └── go.mod        # Go 模块定义
├── frontend/         # 前端项目
│   ├── src/
│   │   ├── api/      # API 接口层
│   │   ├── components/ # 公共组件
│   │   ├── router/   # 路由配置
│   │   ├── stores/   # 状态管理
│   │   ├── utils/    # 工具函数
│   │   ├── views/    # 页面组件
│   │   ├── App.vue   # 根组件
│   │   └── main.ts   # 应用入口
│   ├── index.html
│   ├── package.json
│   ├── tsconfig.json # TypeScript 配置
│   └── vite.config.ts # Vite 配置
└── AGENTS.md         # 项目上下文文件
```

## 构建和运行

### 后端

#### 前置条件

**方式一：本地环境**
- Go 1.24.12 或更高版本
- MySQL 5.7 或更高版本（推荐 MySQL 8.0）
- Redis 5.0 或更高版本（推荐 Redis 7.2）

**方式二：Docker 环境**
- Docker 和 Docker Compose
- 项目提供 `backend/docker/docker-compose.yaml` 配置文件

#### 配置文件

配置文件位于 `backend/config/config.yaml`，包含以下配置项：

```yaml
server:
  port: ":8080"           # 服务器端口
  mode: "debug"           # 运行模式: debug/release
  read_timeout: 60        # 读取超时（秒）
  write_timeout: 60       # 写入超时（秒）
  log_file: "./logs/app.log" # 日志文件路径
  max_size: 128           # 日志文件最大大小(MB)
  max_backups: 30         # 保留旧文件的最大个数
  max_age: 7              # 保留旧文件的最大天数
  compress: true          # 是否压缩旧文件
  is_console: true        # 是否输出到控制台

database:
  host: "localhost"       # 数据库主机
  port: 13306             # 数据库端口（Docker 部署使用 13306，本地 MySQL 可使用 3306）
  user: "root"            # 数据库用户
  password: "123456"      # 数据库密码
  dbname: "admin_system"  # 数据库名称
  charset: "utf8mb4"      # 字符集
  log_file: "./logs/gorm.log" # GORM日志文件
  log_level: "info"       # GORM日志级别 (silent/error/warn/info)

redis:
  addr: "localhost:6379"  # Redis 地址
  password: "123456"      # Redis 密码 (更新为当前配置)
  db: 0                   # Redis 数据库

jwt:
  secret: "your-secret-key"  # JWT 密钥
  timeout: 2               # Token 过期时间（秒）(更新为当前配置)
```

#### 构建命令

```bash
# 后端
cd backend
go mod tidy

# 使用 Docker 启动依赖服务（MySQL + Redis）
cd docker
docker-compose up -d
cd ..

# 构建
go build -o backend main.go

# 运行
./backend

# 或直接运行
go run main.go
```

### 前端

#### 前置条件

- Node.js 16+ (推荐使用 LTS 版本)

#### 构建命令

```bash
# 前端
cd frontend
npm install

# 开发模式
npm run dev

# 生产构建
npm run build

# 预览生产构建
npm run preview
```

### 数据库初始化

项目提供了多个 SQL 脚本用于数据库初始化：

- **`backend/docker/init.sql`** - Docker 部署时的初始化脚本
- **`backend/sql/hosts.sql`** - 主机相关表结构
- **`backend/sql/credentials.sql`** - 凭据相关表结构
- **`backend/sql/add_host_permissions.sql`** - 添加主机权限
- **`backend/sql/test_data.sql`** - 测试数据
- **`backend/sql/generate_test_data.sql`** - 生成测试数据的脚本
- **`backend/sql/quick_test_data.sql`** - 快速测试数据

**初始化步骤**：

1. 使用 Docker Compose 启动 MySQL 时会自动执行 `init.sql`
2. 如需手动初始化，可执行相关 SQL 脚本：
   ```bash
   mysql -h localhost -P 13306 -u root -p admin_system < backend/sql/hosts.sql
   mysql -h localhost -P 13306 -u root -p admin_system < backend/sql/credentials.sql
   mysql -h localhost -P 13306 -u root -p admin_system < backend/sql/add_host_permissions.sql
   ```

## 核心功能模块

### 1. 用户管理

- 用户注册、登录
- 用户信息 CRUD
- 用户状态管理（启用/禁用）
- 密码修改
- 用户与角色关联
- 用户角色分配与移除

### 2. 角色管理

- 角色创建、编辑、删除
- 角色标识符管理
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
- 权限资源管理
- 权限策略管理（添加、移除、查询）

### 5. 操作日志

- 记录用户操作
- 日志查询和删除
- 记录请求和响应信息

### 6. 验证码

- 图形验证码生成
- 登录验证码验证

### 7. 资产管理

#### 7.1 主机管理

- 主机信息管理（主机名、IP地址、端口等）
- 主机组管理
- 主机监控指标管理
- 主机状态管理（在线/离线/故障）
- 主机监控状态管理

#### 7.2 凭据管理

- 凭据信息管理（用户名和密码集中管理）
- 凭据与主机的多对多关联
- 凭据的CRUD操作（创建、查询、更新、删除）
- 凭据安全存储和使用
- 支持多个主机共享相同凭据
- 完整的前端管理界面（CredentialManageView.vue）

## 关键API接口

### 用户角色分配

**为用户分配角色**:
- `POST /users-roles/:username`
- 请求体：`{ "role_id": 1 }`
- 为指定用户分配角色

**移除用户角色**:
- `DELETE /users-roles/:username`
- 请求体：`{ "role_id": 1 }`
- 移除用户的角色

**获取用户角色**:
- `GET /users-roles/:username`
- 获取指定用户的角色列表

### 菜单权限管理

**获取角色的菜单权限**:
- `GET /roles/:id/menus`
- 响应格式：`{"data":[{"p_id":1,"m_id":null},{"p_id":2,"m_id":[11,12,14]}]}`
- `p_id` 代表父菜单ID，`m_id` 代表该父菜单下的子菜单ID数组

**分配菜单权限**:
- `POST /roles/:id/menus`
- 请求体：`{ "menu_ids": [1, 2, 3] }`
- 为角色分配指定的菜单权限

**移除菜单权限**:
- `DELETE /roles/:id/menus`
- 请求体：`{ "menu_ids": [2, 3] }`
- 移除角色的指定菜单权限

### 权限资源管理

**创建权限资源**:
- `POST /permissions`
- 请求体：`{ "path": "/api/users", "method": "GET", "description": "获取用户列表", "status": 1 }`
- 创建新的权限资源

**获取权限资源列表**:
- `GET /permissions`
- 查询参数：`page`, `page_size`, `path`, `method`
- 分页获取权限资源列表

**获取所有权限资源**:
- `GET /permissions/all`
- 查询参数：`path`, `method`
- 获取所有权限资源（不分页）

**获取权限资源详情**:
- `GET /permissions/:id`
- 获取指定权限资源的详细信息

**更新权限资源**:
- `PUT /permissions/:id`
- 请求体：`{ "path": "/api/users", "method": "POST", "description": "创建用户" }`
- 更新权限资源信息

**更新权限状态**:
- `PUT /permissions/:id/status`
- 请求体：`{ "status": 0 }`
- 更新权限资源的状态（启用/禁用）

**删除权限资源**:
- `DELETE /permissions/:id`
- 删除指定的权限资源

### 权限策略管理

**为角色添加策略**:
- `POST /roles/:id/policies`
- 请求体：`{ "path": "/api/users", "method": "GET" }`
- 为角色添加权限策略

**移除角色策略**:
- `DELETE /roles/:id/policies`
- 请求体：`{ "path": "/api/users", "method": "GET" }`
- 移除角色的权限策略

**获取角色策略**:
- `GET /roles/:id/policies`
- 获取角色的所有权限策略

### 资产管理API

#### 主机管理API

**创建主机**:
- `POST /api/v1/hosts`
- 请求体：`{ "hostname": "server01", "ip_address": "192.168.1.100", "port": 22, "os_type": "linux", "group_id": 1, "credential_ids": [1, 2] }`
- 创建新的主机记录，支持关联凭据ID列表

**获取主机列表**:
- `GET /api/v1/hosts`
- 查询参数：`page`, `page_size`, `hostname`, `ip_address`, `group_id`, `status`, `os_type`
- 分页获取主机列表

**获取主机详情**:
- `GET /api/v1/hosts/{id}`
- 获取指定主机的详细信息

**更新主机信息**:
- `PUT /api/v1/hosts/{id}`
- 更新主机的基本信息，包括凭据关联

**删除主机**:
- `DELETE /api/v1/hosts/{id}`
- 删除指定的主机记录

**批量删除主机**:
- `DELETE /api/v1/hosts/batch`
- 批量删除多个主机记录

**更新主机状态**:
- `PUT /api/v1/hosts/{id}/status`
- 更新主机的状态

**更新主机监控状态**:
- `PUT /api/v1/hosts/{id}/monitoring`
- 更新主机的监控启用状态

**获取主机最新监控指标**:
- `GET /api/v1/host-metrics/latest`
- 查询参数：`host_id`
- 获取指定主机的最新监控指标

#### 凭据管理API

**创建凭据**:
- `POST /api/v1/credentials`
- 请求体：`{ "name": "生产服务器凭据", "username": "root", "password": "encrypted_password", "description": "生产环境服务器通用凭据" }`
- 创建新的凭据记录

**获取凭据列表**:
- `GET /api/v1/credentials`
- 查询参数：`page`, `page_size`, `name`, `username`
- 分页获取凭据列表

**获取凭据详情**:
- `GET /api/v1/credentials/{id}`
- 获取指定凭据的详细信息

**更新凭据信息**:
- `PUT /api/v1/credentials/{id}`
- 更新凭据的基本信息

**删除凭据**:
- `DELETE /api/v1/credentials/{id}`
- 删除指定的凭据记录

**批量删除凭据**:
- `DELETE /api/v1/credentials/batch`
- 批量删除多个凭据记录

**获取主机关联的凭据**:
- `GET /api/v1/credentials/host`
- 查询参数：`host_id`
- 获取指定主机关联的凭据信息

## 开发规范

### 命名约定

- **包名**: 小写字母，简洁明了（如 `config`, `model`, `service`）
- **结构体名**: 驼峰命名，导出结构体首字母大写
- **接口名**: 通常以 `er` 结尾（如 `UserService`）
- **函数名**: 导出函数首字母大写，私有函数首字母小写
- **变量名**: 驼峰命名，首字母小写

### 代码格式

- 使用 `go fmt` 自动格式化 Go 代码
- 使用 `npm run lint` 检查前端代码格式
- 每行代码长度不超过 120 个字符

### 分层架构职责

**后端**:
- **Handler 层**: 处理 HTTP 请求和响应，参数验证，调用 Service 层
- **Service 层**: 业务逻辑处理，事务管理，调用 Repository 层
- **Repository 层**: 数据访问操作，数据库 CRUD
- **Model 层**: 数据结构定义，GORM 标签
- **中间件层**: 请求处理前后的中间件（认证、日志、CORS）

**前端**:
- **API 层**: 与后端 API 交互的封装
- **Router 层**: 路由配置和守卫
- **Store 层**: 状态管理（Pinia）
- **View 层**: 页面组件和业务逻辑
- **Component 层**: 可复用的 UI 组件

### 数据库规范

- **表名**: 使用复数形式（如 `users`, `roles`, `menus`, `credentials`）
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

### 前端路由配置

- `/login` - 登录页面
- `/no-permission` - 无权限访问页面
- `/` - 主布局（包含子路由）
  - `/dashboard` - 首页
  - `/users` - 用户管理
  - `/roles` - 角色管理
  - `/menus` - 菜单管理
  - `/operation-logs` - 操作日志
  - `/assets/hosts` - 主机管理
  - `/assets/groups` - 主机组管理
  - `/assets/credentials` - 凭据管理

## 前端菜单权限处理

前端在处理菜单权限时有特殊的逻辑：

1. **API响应格式处理**: `GET /roles/:id/menus` 接口返回格式为 `{"data":[{"p_id":1,"m_id":null},{"p_id":2,"m_id":[11,12,14]}]}`
   - `p_id` 代表父菜单ID
   - `m_id` 代表该父菜单下的子菜单ID数组或null
   - 当收到此格式响应时，前端会将 `p_id` 和 `m_id` 中的所有ID都作为选中的菜单ID

2. **菜单选中逻辑**:
   - 当 `m_id` 为数组时，选中所有子菜单ID
   - 当 `m_id` 为null时，仍然选中 `p_id` 代表的父菜单
   - 使用 `check-strictly="true"` 模式，使父菜单和子菜单独立选择，不相互影响

3. **权限提交逻辑**:
   - 比较当前选中的菜单与角色已有的权限
   - 计算需要新增和移除的权限差异
   - 只对发生变化的权限进行操作

## 数据模型

### Permission（权限模型）

```go
type Permission struct {
    BaseModel
    Path        string `json:"path" gorm:"type:varchar(255);not null;comment:请求路径"`
    Method      string `json:"method" gorm:"type:varchar(10);not null;comment:请求方法"`
    Description string `json:"description" gorm:"type:varchar(255);comment:权限描述"`
    Status      int8   `json:"status" gorm:"type:tinyint;default:1;comment:请求路径"`
}

func (Permission) TableName() string {
    return "permission"
}
```

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

### Credential（凭据模型）

```go
type Credential struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    Name        string         `gorm:"size:100;not null;uniqueIndex" json:"name"`    // 凭据名称
    Username    string         `gorm:"size:50;not null" json:"username"`             // 登录用户名
    Password    string         `gorm:"size:255;not null" json:"password"`            // 加密后的密码
    Description string         `gorm:"size:500" json:"description"`                  // 凭据描述
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
    Hosts       []Host         `gorm:"many2many:host_credentials" json:"hosts"`      // 关联的主机
}

func (Credential) TableName() string {
    return "credentials"
}
```

### Host（主机模型）

```go
type Host struct {
    ID               uint           `gorm:"primaryKey" json:"id"`
    Hostname         string         `gorm:"size:100;not null;uniqueIndex" json:"hostname"`                 // 主机名
    IPAddress        string         `gorm:"size:45;not null;uniqueIndex" json:"ip_address"`                // IP地址
    Port             uint16         `gorm:"default:22" json:"port"`                                        // SSH端口
    OSType           string         `gorm:"size:20;default:'linux'" json:"os_type"`                        // 操作系统类型
    CPUCores         *uint16        `json:"cpu_cores"`                                                     // CPU核心数
    MemoryGB         *uint16        `json:"memory_gb"`                                                     // 内存大小(GB)
    DiskSpaceGB      *uint32        `json:"disk_space_gb"`                                                 // 磁盘空间(GB)
    GroupID          uint           `gorm:"not null" json:"group_id"`                                      // 所属主机组ID
    Status           int8           `gorm:"default:1" json:"status"`                                       // 主机状态: 1-在线, 0-离线, -1-故障
    MonitoringEnable int8           `gorm:"column:monitoring_enabled;default:1" json:"monitoring_enabled"` // 监控是否启用
    LastHeartbeat    *time.Time     `json:"last_heartbeat"`                                                // 最后心跳时间
    Description      string         `gorm:"size:500" json:"description"`                                   // 主机描述
    CreatedBy        *uint          `json:"created_by"`                                                    // 创建人用户ID
    UpdatedBy        *uint          `json:"updated_by"`                                                    // 更新人用户ID
    CreatedAt        time.Time      `json:"created_at"`
    UpdatedAt        time.Time      `json:"updated_at"`
    DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
    Group            *HostGroup     `gorm:"foreignKey:GroupID" json:"group"`                               // 关联的主机组
    Credentials      []Credential   `gorm:"many2many:host_credentials" json:"credentials"`                 // 关联的凭据
}
```

### HostGroup（主机组模型）

```go
type HostGroup struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    Name        string         `gorm:"size:100;not null;uniqueIndex" json:"name"` // 主机组名称
    Description string         `gorm:"size:500" json:"description"`               // 描述信息
    Status      int8           `gorm:"default:1" json:"status"`                   // 状态: 1-启用, 0-禁用
    CreatedBy   *uint          `json:"created_by"`                                // 创建人用户ID
    UpdatedBy   *uint          `json:"updated_by"`                                // 更新人用户ID
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
    Hosts       []Host         `gorm:"foreignKey:GroupID" json:"hosts"` // 关联的主机
}
```

### HostMetric（主机监控指标模型）

```go
type HostMetric struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    HostID      uint      `gorm:"not null" json:"host_id"`                         // 主机ID
    MetricType  string    `gorm:"size:30;not null" json:"metric_type"`             // 指标类型: cpu,memory,disk,network
    MetricName  string    `gorm:"size:50;not null" json:"metric_name"`             // 指标名称
    MetricValue float64   `gorm:"type:decimal(10,2);not null" json:"metric_value"` // 指标值
    Unit        string    `gorm:"size:20" json:"unit"`                             // 单位
    RecordedAt  time.Time `gorm:"not null" json:"recorded_at"`                     // 记录时间
    Host        *Host     `gorm:"foreignKey:HostID" json:"host"`                   // 关联的主机
}

func (HostMetric) TableName() string {
    return "host_metrics"
}
```

## 错误处理

- 错误信息应清晰明了
- 使用 `errors.New()` 或 `fmt.Errorf()` 创建错误
- 统一错误处理格式

## 日志规范

后端使用 Zap 日志库，输出 JSON 格式日志：

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

## Git 提交规范

- `feat`: 新功能
- `fix`: 修复 bug
- `docs`: 文档更新
- `style`: 代码格式调整
- `refactor`: 代码重构
- `test`: 测试相关
- `chore`: 构建过程或辅助工具的变动

示例：
```
feat(asset): 添加资产管理模块 - 实现主机管理和凭据管理功能
- 添加主机信息管理（主机名、IP地址、端口等）
- 实现主机组管理和主机监控指标管理
- 添加凭据信息集中安全管理功能
- 实现凭据与主机的多对多关联管理
- 完善资产管理系统API接口和数据库表结构
- 更新项目文档中的技术架构和功能说明
```

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
- 凭据信息集中安全管理

## 测试

### 单元测试

```bash
# 后端测试
cd backend
go test ./...

# 前端测试
cd frontend
npm run test
```

### API 测试

使用 Postman 或类似工具测试 API 接口。参考 `backend/docs/API文档.md` 获取详细的接口说明。

## 部署

### 开发环境

- 后端运行在 `http://localhost:8080`
- 前端运行在 `http://localhost:3000`，通过代理转发 API 请求到后端

### Docker 部署

项目提供 Docker Compose 配置文件，可以快速启动开发环境：

```bash
# 启动服务（MySQL + Redis）
cd backend/docker
docker-compose up -d

# 查看服务状态
docker-compose ps

# 停止服务
docker-compose down
```

Docker 服务配置：
- **MySQL**: 端口 13306，密码 123456
- **Redis**: 端口 6379，密码 123456

### 生产环境

- 构建前端项目并部署到静态服务器
- 后端服务独立部署
- 配置反向代理（如 Nginx）统一处理请求

### 生产环境检查清单

- [ ] 修改配置文件中的敏感信息（数据库密码、JWT 密钥等）
- [ ] 设置 Gin 模式为 `release`
- [ ] 配置 HTTPS
- [ ] 设置防火墙规则
- [ ] 配置日志轮转
- [ ] 设置数据库备份策略
- [ ] 配置监控和告警

## 项目状态

- **当前版本**: 1.0.0
- **开发状态**: 功能完整，持续优化中
- **核心功能**: 已全部实现并可用
  - 用户管理 ✅
  - 角色管理 ✅
  - 菜单管理 ✅
  - 权限管理 ✅
  - 操作日志 ✅
  - 资产管理（主机、主机组、凭据）✅

## 重要注意事项

1. **完整实现**: 资产管理模块（主机管理、主机组管理、凭据管理）前后端已完整实现，包括 API 接口、业务逻辑、数据模型和前端界面。

2. **权限控制**: 系统采用基于 Casbin 的 RBAC 权限模型，实现细粒度的权限控制，包括菜单权限和接口权限。

3. **数据安全**: 凭据信息采用加密存储，确保敏感数据的安全性。

## 联系方式

如有问题或建议，请联系开发团队。

---

**注意**: 请勿在代码中硬编码敏感信息（如密码、API 密钥等）。所有敏感配置应通过配置文件或环境变量管理。