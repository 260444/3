-- =====================================================
-- 更新菜单数据以支持动态路由
-- 执行日期: 2026-01-18
-- =====================================================

DELETE FROM menus;
delete from role_menus;

-- 清空现有菜单数据

-- =====================================================
-- 1. 一级菜单
-- =====================================================
INSERT INTO menus (id, name, title, path, component, icon, sort, is_hidden, is_link, status, created_at, updated_at) VALUES
(1, 'dashboard', '控制台', '/dashboard', 'DashboardView', 'House', 1, 0, 0, 1, NOW(), NOW()),
(2, 'system', '系统管理', '/system', '', 'Setting', 100, 0, 0, 1, NOW(), NOW());

-- =====================================================
-- 2. 二级菜单 - 系统管理
-- =====================================================
INSERT INTO menus (id, name, title, path, component, icon, sort, is_hidden, is_link, parent_id, status, created_at, updated_at) VALUES
(11, 'users', '用户管理', '/users', 'UserManageView', 'User', 1, 0, 0, 2, 1, NOW(), NOW()),
(12, 'roles', '角色管理', '/roles', 'RoleManageView', 'Avatar', 2, 0, 0, 2, 1, NOW(), NOW()),
(13, 'menus', '菜单管理', '/menus', 'MenuManageView', 'Menu', 3, 0, 0, 2, 1, NOW(), NOW()),
(14, 'operation-logs', '操作日志', '/operation-logs', 'OperationLogView', 'Document', 4, 0, 0, 2, 1, NOW(), NOW()),
(15, 'permissions', '权限管理', '/permissions', 'PermissionManageView', 'Lock', 5, 0, 0, 2, 1, NOW(), NOW());

-- =====================================================
-- 3. 角色菜单关联数据
-- =====================================================
-- 超级管理员拥有所有菜单权限
INSERT INTO role_menus (role_id, menu_id) VALUES
(1, 1), (1, 2),
(1, 11), (1, 12), (1, 13), (1, 14), (1, 15);

-- 管理员拥有部分菜单权限
INSERT INTO role_menus (role_id, menu_id) VALUES
(2, 1), (2, 2),
(2, 11), (2, 12), (2, 13), (2, 14);

-- 普通用户拥有基本权限
INSERT INTO role_menus (role_id, menu_id) VALUES
(3, 1), (3, 2),
(3, 11);

-- 访客仅有查看权限
INSERT INTO role_menus (role_id, menu_id) VALUES
(4, 1);

-- =====================================================
-- 菜单字段说明
-- =====================================================
-- id: 菜单ID（唯一标识）
-- name: 路由名称（前端路由的 name，必须唯一）
-- title: 菜单标题（显示在侧边栏的文本）
-- path: 路由路径（URL路径，如 /dashboard, /users）
-- component: 组件名称（对应前端组件映射表中的键名）
--   可选值：
--   - DashboardView: 首页
--   - UserManageView: 用户管理
--   - RoleManageView: 角色管理
--   - MenuManageView: 菜单管理
--   - OperationLogView: 操作日志
--   - PermissionManageView: 权限管理
--   - 空字符串: 表示该菜单是父菜单，没有具体页面
-- icon: 图标名称（Element Plus 图标组件名）
--   可选值：House, Setting, User, Avatar, Menu, Document, Lock 等
-- sort: 排序（数字越小越靠前）
-- is_hidden: 是否隐藏（0: 显示, 1: 隐藏）
-- is_link: 是否外部链接（0: 否, 1: 是）
-- parent_id: 父菜单ID（一级菜单为 NULL）
-- status: 状态（1: 正常, 0: 禁用）
-- =====================================================

-- =====================================================
-- 添加新菜单的示例
-- =====================================================
-- 添加一级菜单：
-- INSERT INTO menus (name, title, path, component, icon, sort, is_hidden, is_link, status, created_at, updated_at)
-- VALUES ('monitor', '系统监控', '/monitor', '', 'Monitor', 200, 0, 0, 1, NOW(), NOW());

-- 添加二级菜单：
-- INSERT INTO menus (name, title, path, component, icon, sort, is_hidden, is_link, parent_id, status, created_at, updated_at)
-- VALUES ('server', '服务监控', '/server', 'ServerMonitorView', 'Monitor', 1, 0, 0, (SELECT id FROM menus WHERE name = 'monitor'), 1, NOW(), NOW());

-- 为角色分配菜单权限：
-- INSERT INTO role_menus (role_id, menu_id)
-- VALUES (1, (SELECT id FROM menus WHERE name = 'monitor'));
-- =====================================================