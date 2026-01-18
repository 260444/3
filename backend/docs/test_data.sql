-- =====================================================
-- 后端测试数据 SQL
-- 生成日期: 2026-01-18
-- =====================================================



-- =====================================================
-- 1. 角色数据
-- =====================================================
INSERT INTO roles (id, name, description, status, created_at, updated_at) VALUES
(1, '超级管理员', '拥有系统所有权限', 1, NOW(), NOW()),
(2, '管理员', '拥有大部分管理权限', 1, NOW(), NOW()),
(3, '普通用户', '拥有基本查看权限', 1, NOW(), NOW()),
(4, '访客', '仅拥有查看权限', 1, NOW(), NOW()),
(5, '测试角色', '用于测试的角色', 1, NOW(), NOW());

-- =====================================================
-- 2. 菜单数据
-- =====================================================
-- 一级菜单
INSERT INTO menus (id, name, title, path, component, icon, sort, is_hidden, is_link, status, created_at, updated_at) VALUES
(1, 'dashboard', '控制台', '/dashboard', 'views/dashboard/index', 'dashboard', 1, 0, 0, 1, NOW(), NOW()),
(2, 'system', '系统管理', '/system', 'Layout', 'setting', 100, 0, 0, 1, NOW(), NOW()),
(3, 'monitor', '系统监控', '/monitor', 'Layout', 'monitor', 200, 0, 0, 1, NOW(), NOW()),
(4, 'tools', '工具箱', '/tools', 'Layout', 'tool', 300, 0, 0, 1, NOW(), NOW());

-- 二级菜单 - 系统管理
INSERT INTO menus (id, name, title, path, component, icon, sort, is_hidden, is_link, parent_id, status, created_at, updated_at) VALUES
(11, 'user', '用户管理', '/system/user', 'views/system/user/index', 'user', 1, 0, 0, 2, 1, NOW(), NOW()),
(12, 'role', '角色管理', '/system/role', 'views/system/role/index', 'user-group', 2, 0, 0, 2, 1, NOW(), NOW()),
(13, 'menu', '菜单管理', '/system/menu', 'views/system/menu/index', 'menu', 3, 0, 0, 2, 1, NOW(), NOW()),
(14, 'log', '操作日志', '/system/log', 'views/system/log/index', 'document', 4, 0, 0, 2, 1, NOW(), NOW());

-- 二级菜单 - 系统监控
INSERT INTO menus (id, name, title, path, component, icon, sort, is_hidden, is_link, parent_id, status, created_at, updated_at) VALUES
(21, 'server', '服务监控', '/monitor/server', 'views/monitor/server/index', 'server', 1, 0, 0, 3, 1, NOW(), NOW()),
(22, 'cache', '缓存监控', '/monitor/cache', 'views/monitor/cache/index', 'database', 2, 0, 0, 3, 1, NOW(), NOW()),
(23, 'online', '在线用户', '/monitor/online', 'views/monitor/online/index', 'peoples', 3, 0, 0, 3, 1, NOW(), NOW());

-- 二级菜单 - 工具箱
INSERT INTO menus (id, name, title, path, component, icon, sort, is_hidden, is_link, parent_id, status, created_at, updated_at) VALUES
(31, 'generator', '代码生成', '/tools/generator', 'views/tools/generator/index', 'code', 1, 0, 0, 4, 1, NOW(), NOW()),
(32, 'swagger', '系统接口', '/tools/swagger', 'views/tools/swagger/index', 'link', 2, 0, 0, 4, 1, NOW(), NOW());

-- =====================================================
-- 3. 角色菜单关联数据
-- =====================================================
-- 超级管理员拥有所有菜单权限
INSERT INTO role_menus (role_id, menu_id) VALUES
(1, 1), (1, 2), (1, 3), (1, 4),
(1, 11), (1, 12), (1, 13), (1, 14),
(1, 21), (1, 22), (1, 23),
(1, 31), (1, 32);

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
-- 4. 用户数据
-- 注意：密码为 '123456' 的 bcrypt 加密值
-- =====================================================
INSERT INTO users (id, username, password, email, phone, nickname, avatar, status, role_id, created_at, updated_at) VALUES
-- 超级管理员
(1, 'admin', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'admin@example.com', '13800138000', '超级管理员', '', 1, 1, NOW(), NOW()),
-- 管理员
(2, 'manager', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'manager@example.com', '13800138001', '管理员', '', 1, 2, NOW(), NOW()),
-- 普通用户
(3, 'user', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'user@example.com', '13800138002', '普通用户', '', 1, 3, NOW(), NOW()),
-- 访客
(4, 'guest', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'guest@example.com', '13800138003', '访客', '', 1, 4, NOW(), NOW()),
-- 测试用户
(5, 'test', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'test@example.com', '13800138004', '测试用户', '', 1, 5, NOW(), NOW()),
-- 更多测试用户
(6, 'zhangsan', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'zhangsan@example.com', '13800138005', '张三', '', 1, 3, NOW(), NOW()),
(7, 'lisi', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'lisi@example.com', '13800138006', '李四', '', 1, 3, NOW(), NOW()),
(8, 'wangwu', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'wangwu@example.com', '13800138007', '王五', '', 1, 3, NOW(), NOW()),
-- 禁用用户
(9, 'disabled', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'disabled@example.com', '13800138008', '禁用用户', '', 0, 3, NOW(), NOW()),
-- 另一个管理员
(10, 'admin2', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'admin2@example.com', '13800138009', '管理员2', '', 1, 2, NOW(), NOW());

-- =====================================================
-- 5. 操作日志数据
-- =====================================================
INSERT INTO operation_logs (operation, user_id, username, ip, user_agent, status, request_method, request_path, request_body, response_body, response_time, created_at, updated_at) VALUES
('用户登录', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/login', '{"username":"admin","password":"***"}', '{"message":"登录成功"}', 156, NOW(), NOW()),
('用户登录', 2, 'manager', '192.168.1.101', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/login', '{"username":"manager","password":"***"}', '{"message":"登录成功"}', 142, NOW(), NOW()),
('查询用户列表', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/users', '', '{"data":[...]}', 89, NOW(), NOW()),
('创建用户', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/users', '{"username":"newuser","email":"new@example.com"}', '{"message":"创建成功"}', 234, NOW(), NOW()),
('更新用户信息', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'PUT', '/api/v1/users/6', '{"nickname":"张三更新"}', '{"message":"更新成功"}', 178, NOW(), NOW()),
('删除用户', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'DELETE', '/api/v1/users/99', '', '{"message":"删除成功"}', 145, NOW(), NOW()),
('查询角色列表', 2, 'manager', '192.168.1.101', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/roles', '', '{"data":[...]}', 76, NOW(), NOW()),
('创建角色', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/roles', '{"name":"新角色","description":"测试角色"}', '{"message":"创建成功"}', 198, NOW(), NOW()),
('分配角色权限', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'PUT', '/api/v1/roles/5/menus', '{"menu_ids":[1,2,11]}', '{"message":"分配成功"}', 267, NOW(), NOW()),
('查询菜单树', 3, 'user', '192.168.1.102', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/menus', '', '{"data":[...]}', 112, NOW(), NOW()),
('创建菜单', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/menus', '{"name":"test","title":"测试菜单"}', '{"message":"创建成功"}', 189, NOW(), NOW()),
('修改密码', 2, 'manager', '192.168.1.101', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'PUT', '/api/v1/users/change-password', '{"old_password":"***","new_password":"***"}', '{"message":"修改成功"}', 156, NOW(), NOW()),
('查询操作日志', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/operation-logs', '', '{"data":[...]}', 134, NOW(), NOW()),
('删除操作日志', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'DELETE', '/api/v1/operation-logs/1', '', '{"message":"删除成功"}', 98, NOW(), NOW()),
('用户登录失败', 9, 'disabled', '192.168.1.109', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 0, 'POST', '/api/v1/login', '{"username":"disabled","password":"***"}', '{"error":"用户已被禁用"}', 67, NOW(), NOW()),
('查询用户详情', 3, 'user', '192.168.1.102', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/users/3', '', '{"data":{...}}', 45, NOW(), NOW()),
('获取验证码', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/captcha', '', '{"data":"base64image"}', 23, NOW(), NOW()),
('用户注册', 3, 'user', '192.168.1.102', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/register', '{"username":"register_user","email":"register@example.com"}', '{"message":"注册成功"}', 289, NOW(), NOW()),
('更新用户状态', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'PUT', '/api/v1/users/9/status', '{"status":0}', '{"message":"状态更新成功"}', 134, NOW(), NOW());

-- =====================================================
-- 数据统计
-- =====================================================
-- 角色数量: 5
-- 菜单数量: 15 (4个一级菜单 + 11个二级菜单)
-- 用户数量: 10 (1个超级管理员、2个管理员、6个普通用户、1个禁用用户)
-- 操作日志数量: 20
-- =====================================================

-- 查询验证
-- SELECT '角色数量' as 类型, COUNT(*) as 数量 FROM roles
-- UNION ALL
-- SELECT '菜单数量', COUNT(*) FROM menus
-- UNION ALL
-- SELECT '用户数量', COUNT(*) FROM users
-- UNION ALL
-- SELECT '正常用户', COUNT(*) FROM users WHERE status = 1
-- UNION ALL
-- SELECT '禁用用户', COUNT(*) FROM users WHERE status = 0
-- UNION ALL
-- SELECT '操作日志', COUNT(*) FROM operation_logs;