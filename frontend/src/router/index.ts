import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { getUserMenus } from '@/api/menu'
import { ElMessage } from 'element-plus'
import LoginView from '@/views/LoginView.vue'
import LayoutView from '@/views/LayoutView.vue'
import DashboardView from '@/views/DashboardView.vue'
import UserManageView from '@/views/user/UserManageView.vue'
import RoleManageView from '@/views/role/RoleManageView.vue'
import MenuManageView from '@/views/menu/MenuManageView.vue'
import OperationLogView from '@/views/OperationLogView.vue'
import PermissionManageView from '@/views/permission/PermissionManageView.vue'

// 静态路由（不需要权限）
const constantRoutes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'login',
    component: LoginView,
    meta: { title: '登录' }
  }
]

// 异步路由（需要权限）
const asyncRoutes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'layout',
    component: LayoutView,
    redirect: '/dashboard',
    meta: { title: '首页' }
  }
]

// 组件映射表：将后端返回的 component 字段映射到实际组件
const componentMap: Record<string, any> = {
  'DashboardView': DashboardView,
  'UserManageView': UserManageView,
  'RoleManageView': RoleManageView,
  'MenuManageView': MenuManageView,
  'OperationLogView': OperationLogView,
  'PermissionManageView': PermissionManageView
}

const router = createRouter({
  history: createWebHistory(),
  routes: constantRoutes
})

// 将后端菜单转换为路由
function generateRoutes(menus: any[]): RouteRecordRaw[] {
  const routes: RouteRecordRaw[] = []

  menus.forEach(menu => {
    const route: any = {
      path: menu.path,
      name: menu.name,
      meta: {
        title: menu.title,
        icon: menu.icon,
        hidden: menu.is_hidden
      }
    }

    // 如果有 component 且不为空字符串，则设置组件
    if (menu.component && menu.component.trim() !== '' && componentMap[menu.component]) {
      route.component = componentMap[menu.component]
      route.meta!.component = menu.component
    }

    // 如果有 redirect，设置重定向
    if (menu.redirect) {
      route.redirect = menu.redirect
    }

    // 如果有子菜单，递归处理
    if (menu.children && menu.children.length > 0) {
      route.children = generateRoutes(menu.children)
    }

    routes.push(route)
  })

  return routes
}

// 添加动态路由
function addDynamicRoutes(menus: any[]) {
  const layoutRoute = router.getRoutes().find(r => r.name === 'layout')
  
  if (!layoutRoute) {
    // 如果 layout 路由不存在，先添加
    router.addRoute(asyncRoutes[0])
  }

  // 生成路由并添加
  const routes = generateRoutes(menus)
  routes.forEach(route => {
    router.addRoute('layout', route)
  })
}

// 重置路由
function resetRouter() {
  const layoutRoute = router.getRoutes().find(r => r.name === 'layout')
  if (layoutRoute) {
    // 移除所有动态添加的路由
    router.removeRoute('layout')
  }
}

// 路由守卫
router.beforeEach(async (to, from, next) => {
  const userStore = useUserStore()
  const token = localStorage.getItem('token')

  if (to.path === '/login') {
    if (token) {
      next('/')
    } else {
      next()
    }
    return
  }

  if (!token) {
    next('/login')
    return
  }

  // 检查是否已经添加了动态路由
  const hasDynamicRoutes = router.getRoutes().some(r => r.name === 'dashboard')
  
  // 如果已经有动态路由，直接放行
  if (hasDynamicRoutes) {
    next()
    return
  }

  try {
    // 获取用户菜单
    const response = await getUserMenus()
    console.log('获取到的菜单数据:', response.data)
    userStore.setMenus(response.data)

    // 添加动态路由
    addDynamicRoutes(response.data)
    console.log('当前路由列表:', router.getRoutes())

    // 重新跳转，确保路由已加载
    next({ ...to, replace: true })
  } catch (error) {
    console.error('获取菜单失败:', error)
    ElMessage.error('获取菜单失败，请重新登录')
    await userStore.logout()
    next('/login')
  }
})

export default router