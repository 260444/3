# Frontend 项目文档

## 项目概述

这是一个基于 Vue 3 + TypeScript + Vite 构建的后台管理系统前端项目。项目采用现代化的前端技术栈，提供了用户管理、角色管理、菜单管理和操作日志等核心功能。

### 主要技术栈

- **框架**: Vue 3 (Composition API)
- **语言**: TypeScript
- **构建工具**: Vite
- **UI 组件库**: Element Plus
- **路由**: Vue Router 4
- **状态管理**: Pinia
- **HTTP 客户端**: Axios

### 项目架构

```
src/
├── api/              # API 接口层
│   ├── menu.ts      # 菜单相关接口
│   ├── operationLog.ts  # 操作日志接口
│   ├── request.ts   # Axios 实例配置
│   ├── role.ts      # 角色相关接口
│   └── user.ts      # 用户相关接口
├── components/      # 公共组件
├── router/          # 路由配置
│   └── index.ts     # 路由定义和守卫
├── stores/          # 状态管理
│   └── user.ts      # 用户状态 store
├── utils/           # 工具函数
├── views/           # 页面组件
│   ├── DashboardView.vue      # 首页
│   ├── LayoutView.vue         # 布局页面
│   ├── LoginView.vue          # 登录页面
│   ├── OperationLogView.vue   # 操作日志页面
│   ├── menu/
│   │   └── MenuManageView.vue # 菜单管理页面
│   ├── role/
│   │   └── RoleManageView.vue # 角色管理页面
│   └── user/
│       └── UserManageView.vue # 用户管理页面
├── App.vue          # 根组件
├── main.ts          # 应用入口
└── vite-env.d.ts    # Vite 类型声明
```

## 构建和运行

### 开发环境

```bash
npm run dev
```

开发服务器将在 `http://localhost:3000` 启动。

### 生产构建

```bash
npm run build
```

此命令会先执行 TypeScript 类型检查 (`vue-tsc`)，然后执行 Vite 构建打包。

### 预览生产构建

```bash
npm run preview
```

## 开发规范

### 路径别名

项目配置了 `@` 别名指向 `src` 目录，推荐使用别名进行导入：

```typescript
import { useUserStore } from '@/stores/user'
import LoginView from '@/views/LoginView.vue'
```

### API 请求配置

- **Base URL**: `/api/v1`（通过 Vite 代理转发到 `http://localhost:8080`）
- **认证方式**: Bearer Token（从 localStorage 获取）
- **超时时间**: 10000ms

API 请求使用统一的 Axios 实例（`src/api/request.ts`），已配置请求和响应拦截器：
- 请求拦截器自动添加 Authorization header
- 响应拦截器处理 401 未授权错误，自动跳转登录页

### 路由配置

路由采用嵌套结构，主要路由包括：

- `/login` - 登录页面
- `/` - 主布局（包含子路由）
  - `/dashboard` - 首页
  - `/users` - 用户管理
  - `/roles` - 角色管理
  - `/menus` - 菜单管理
  - `/operation-logs` - 操作日志

路由守卫已配置，未登录用户访问受保护路由会自动跳转到登录页。

### 状态管理

使用 Pinia 进行状态管理，当前主要 store：

- `useUserStore` - 用户状态管理（登录、登出、用户信息）

### TypeScript 配置

- 启用严格模式 (`strict: true`)
- 未使用变量检查 (`noUnusedLocals`, `noUnusedParameters`)
- 路径别名配置 (`@/*` 映射到 `./src/*`)

### UI 组件使用

项目使用 Element Plus 组件库，所有图标已在 `main.ts` 中全局注册，可直接使用：

```vue
<el-icon><House /></el-icon>
```

### 样式规范

- 使用 scoped CSS 避免样式污染
- 主色调：Element Plus 默认蓝色 (#409EFF)
- 侧边栏背景色：#304156
- 主内容区背景色：#f0f2f5

## 开发注意事项

1. **API 代理**: 开发时 `/api` 请求会被代理到 `http://localhost:8080`，确保后端服务在该端口运行
2. **Token 管理**: Token 存储在 localStorage 中，登录成功后自动设置
3. **类型安全**: 项目使用 TypeScript，建议为 API 响应、组件 props 等定义类型
4. **图标使用**: Element Plus 图标已全局注册，无需额外导入
5. **路由守卫**: 所有非登录路由都需要认证，路由守卫会自动处理

## 依赖说明

### 生产依赖

- `vue`: ^3.4.21 - Vue 3 框架
- `vue-router`: ^4.3.0 - 路由管理
- `pinia`: ^2.1.7 - 状态管理
- `axios`: ^1.6.8 - HTTP 客户端
- `element-plus`: ^2.6.1 - UI 组件库
- `@element-plus/icons-vue`: ^2.3.1 - Element Plus 图标

### 开发依赖

- `@vitejs/plugin-vue`: ^5.0.4 - Vite Vue 插件
- `typescript`: ^5.2.2 - TypeScript 编译器
- `vue-tsc`: ^1.8.27 - Vue TypeScript 类型检查
- `vite`: ^5.2.0 - 构建工具