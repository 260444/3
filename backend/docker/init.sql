/*
 Navicat Premium Dump SQL

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80027 (8.0.27)
 Source Host           : localhost:13306
 Source Schema         : admin_system

 Target Server Type    : MySQL
 Target Server Version : 80027 (8.0.27)
 File Encoding         : 65001

 Date: 29/01/2026 20:41:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype` ASC, `v0` ASC, `v1` ASC, `v2` ASC, `v3` ASC, `v4` ASC, `v5` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 96 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (16, 'g', 'admin', 'admin', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (20, 'g', 'ceshi', 'ceshi', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (19, 'g', 'ceshi2', 'user', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (91, 'g', 'ceshiadmin', 'surperadmin', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (15, 'g', 'superadmin', 'surperadmin', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (56, 'p', 'admin', '/api/v1/captcha', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (55, 'p', 'admin', '/api/v1/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (57, 'p', 'admin', '/api/v1/logout', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (75, 'p', 'admin', '/api/v1/menus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (74, 'p', 'admin', '/api/v1/menus', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (78, 'p', 'admin', '/api/v1/menus/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (77, 'p', 'admin', '/api/v1/menus/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (76, 'p', 'admin', '/api/v1/menus/all', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (86, 'p', 'admin', '/api/v1/permissions', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (85, 'p', 'admin', '/api/v1/permissions', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (90, 'p', 'admin', '/api/v1/permissions/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (87, 'p', 'admin', '/api/v1/permissions/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (88, 'p', 'admin', '/api/v1/permissions/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (89, 'p', 'admin', '/api/v1/permissions/:id/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (95, 'p', 'admin', '/api/v1/permissions/all', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (94, 'p', 'admin', '/api/v1/permissions/all/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (70, 'p', 'admin', '/api/v1/roles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (69, 'p', 'admin', '/api/v1/roles', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (73, 'p', 'admin', '/api/v1/roles/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (71, 'p', 'admin', '/api/v1/roles/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (72, 'p', 'admin', '/api/v1/roles/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (81, 'p', 'admin', '/api/v1/roles/:id/menus', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (80, 'p', 'admin', '/api/v1/roles/:id/menus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (79, 'p', 'admin', '/api/v1/roles/:id/menus', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (83, 'p', 'admin', '/api/v1/roles/:id/policies', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (84, 'p', 'admin', '/api/v1/roles/:id/policies', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (82, 'p', 'admin', '/api/v1/roles/:id/policies', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (59, 'p', 'admin', '/api/v1/users', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (58, 'p', 'admin', '/api/v1/users', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (67, 'p', 'admin', '/api/v1/users-roles/:username', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (68, 'p', 'admin', '/api/v1/users-roles/:username', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (66, 'p', 'admin', '/api/v1/users-roles/:username', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (63, 'p', 'admin', '/api/v1/users/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (60, 'p', 'admin', '/api/v1/users/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (61, 'p', 'admin', '/api/v1/users/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (65, 'p', 'admin', '/api/v1/users/:id/reset-password', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (62, 'p', 'admin', '/api/v1/users/:id/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (64, 'p', 'admin', '/api/v1/users/change-password', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (22, 'p', 'surperadmin', '/api/v1/captcha', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (21, 'p', 'surperadmin', '/api/v1/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (23, 'p', 'surperadmin', '/api/v1/logout', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (39, 'p', 'surperadmin', '/api/v1/menus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (38, 'p', 'surperadmin', '/api/v1/menus', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (42, 'p', 'surperadmin', '/api/v1/menus/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (41, 'p', 'surperadmin', '/api/v1/menus/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (40, 'p', 'surperadmin', '/api/v1/menus/all', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (50, 'p', 'surperadmin', '/api/v1/permissions', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (49, 'p', 'surperadmin', '/api/v1/permissions', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (54, 'p', 'surperadmin', '/api/v1/permissions/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (51, 'p', 'surperadmin', '/api/v1/permissions/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (52, 'p', 'surperadmin', '/api/v1/permissions/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (53, 'p', 'surperadmin', '/api/v1/permissions/:id/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (93, 'p', 'surperadmin', '/api/v1/permissions/all', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (92, 'p', 'surperadmin', '/api/v1/permissions/all/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (34, 'p', 'surperadmin', '/api/v1/roles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (33, 'p', 'surperadmin', '/api/v1/roles', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (37, 'p', 'surperadmin', '/api/v1/roles/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (35, 'p', 'surperadmin', '/api/v1/roles/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (36, 'p', 'surperadmin', '/api/v1/roles/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (45, 'p', 'surperadmin', '/api/v1/roles/:id/menus', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (44, 'p', 'surperadmin', '/api/v1/roles/:id/menus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (43, 'p', 'surperadmin', '/api/v1/roles/:id/menus', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (47, 'p', 'surperadmin', '/api/v1/roles/:id/policies', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (48, 'p', 'surperadmin', '/api/v1/roles/:id/policies', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (46, 'p', 'surperadmin', '/api/v1/roles/:id/policies', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (8, 'p', 'surperadmin', '/api/v1/users', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (9, 'p', 'surperadmin', '/api/v1/users', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (31, 'p', 'surperadmin', '/api/v1/users-roles/:username', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (32, 'p', 'surperadmin', '/api/v1/users-roles/:username', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (30, 'p', 'surperadmin', '/api/v1/users-roles/:username', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (27, 'p', 'surperadmin', '/api/v1/users/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (24, 'p', 'surperadmin', '/api/v1/users/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (25, 'p', 'surperadmin', '/api/v1/users/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (29, 'p', 'surperadmin', '/api/v1/users/:id/reset-password', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (26, 'p', 'surperadmin', '/api/v1/users/:id/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (28, 'p', 'surperadmin', '/api/v1/users/change-password', 'PUT', '', '', '');

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `updated_at` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `redirect` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `parent_id` bigint UNSIGNED NULL DEFAULT 0,
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `sort` bigint NULL DEFAULT 0,
  `is_hidden` tinyint(1) NULL DEFAULT 0,
  `is_link` tinyint(1) NULL DEFAULT 0,
  `link_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `status` bigint NULL DEFAULT 1,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_menus_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_menus_parent`(`parent_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 118 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of menus
-- ----------------------------
INSERT INTO `menus` VALUES (1, '2026-01-18 21:54:50', '2026-01-18 21:54:50', NULL, 'dashboard', '控制台', '/dashboard', 'DashboardView', NULL, 0, 'House', 1, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (2, '2026-01-18 21:54:50', '2026-01-18 21:54:50', NULL, 'system', '系统管理', '/system', '', NULL, 0, 'Setting', 100, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (11, '2026-01-18 21:54:50', '2026-01-18 21:54:50', NULL, 'users', '用户管理', '/users', 'UserManageView', NULL, 2, 'User', 1, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (12, '2026-01-18 21:54:50', '2026-01-18 21:54:50', NULL, 'roles', '角色管理', '/roles', 'RoleManageView', NULL, 2, 'Avatar', 2, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (13, '2026-01-18 21:54:50', '2026-01-18 21:54:50', NULL, 'menus', '菜单管理', '/menus', 'MenuManageView', NULL, 2, 'Menu', 3, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (14, '2026-01-18 21:54:50', '2026-01-18 21:54:50', NULL, 'operation-logs', '操作日志', '/operation-logs', 'OperationLogView', NULL, 2, 'Document', 4, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (15, '2026-01-18 21:54:50', '2026-01-18 21:54:50', NULL, 'permissions', '权限管理', '/permissions', 'PermissionResourceView', NULL, 2, 'Lock', 5, 0, 0, NULL, 1);

-- ----------------------------
-- Table structure for operation_logs
-- ----------------------------
DROP TABLE IF EXISTS `operation_logs`;
CREATE TABLE `operation_logs`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `updated_at` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `operation` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `user_agent` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `status` bigint NULL DEFAULT 1,
  `request_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `request_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `request_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `response_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `response_time` bigint NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_operation_logs_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of operation_logs
-- ----------------------------
INSERT INTO `operation_logs` VALUES (1, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '用户登录', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/login', '{\"username\":\"admin\",\"password\":\"***\"}', '{\"message\":\"登录成功\"}', 156, NULL);
INSERT INTO `operation_logs` VALUES (2, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '用户登录', 2, 'manager', '192.168.1.101', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/login', '{\"username\":\"manager\",\"password\":\"***\"}', '{\"message\":\"登录成功\"}', 142, NULL);
INSERT INTO `operation_logs` VALUES (3, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '查询用户列表', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/users', '', '{\"data\":[...]}', 89, NULL);
INSERT INTO `operation_logs` VALUES (4, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '创建用户', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/users', '{\"username\":\"newuser\",\"email\":\"new@example.com\"}', '{\"message\":\"创建成功\"}', 234, NULL);
INSERT INTO `operation_logs` VALUES (5, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '更新用户信息', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'PUT', '/api/v1/users/6', '{\"nickname\":\"张三更新\"}', '{\"message\":\"更新成功\"}', 178, NULL);
INSERT INTO `operation_logs` VALUES (6, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '删除用户', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'DELETE', '/api/v1/users/99', '', '{\"message\":\"删除成功\"}', 145, NULL);
INSERT INTO `operation_logs` VALUES (7, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '查询角色列表', 2, 'manager', '192.168.1.101', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/roles', '', '{\"data\":[...]}', 76, NULL);
INSERT INTO `operation_logs` VALUES (8, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '创建角色', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/roles', '{\"name\":\"新角色\",\"description\":\"测试角色\"}', '{\"message\":\"创建成功\"}', 198, NULL);
INSERT INTO `operation_logs` VALUES (9, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '分配角色权限', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'PUT', '/api/v1/roles/5/menus', '{\"menu_ids\":[1,2,11]}', '{\"message\":\"分配成功\"}', 267, NULL);
INSERT INTO `operation_logs` VALUES (10, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '查询菜单树', 3, 'user', '192.168.1.102', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/menus', '', '{\"data\":[...]}', 112, NULL);
INSERT INTO `operation_logs` VALUES (11, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '创建菜单', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/menus', '{\"name\":\"test\",\"title\":\"测试菜单\"}', '{\"message\":\"创建成功\"}', 189, NULL);
INSERT INTO `operation_logs` VALUES (12, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '修改密码', 2, 'manager', '192.168.1.101', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'PUT', '/api/v1/users/change-password', '{\"old_password\":\"***\",\"new_password\":\"***\"}', '{\"message\":\"修改成功\"}', 156, NULL);
INSERT INTO `operation_logs` VALUES (13, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '查询操作日志', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/operation-logs', '', '{\"data\":[...]}', 134, NULL);
INSERT INTO `operation_logs` VALUES (14, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '删除操作日志', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'DELETE', '/api/v1/operation-logs/1', '', '{\"message\":\"删除成功\"}', 98, NULL);
INSERT INTO `operation_logs` VALUES (15, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '用户登录失败', 9, 'disabled', '192.168.1.109', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 0, 'POST', '/api/v1/login', '{\"username\":\"disabled\",\"password\":\"***\"}', '{\"error\":\"用户已被禁用\"}', 67, NULL);
INSERT INTO `operation_logs` VALUES (16, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '查询用户详情', 3, 'user', '192.168.1.102', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/users/3', '', '{\"data\":{...}}', 45, NULL);
INSERT INTO `operation_logs` VALUES (17, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '获取验证码', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/captcha', '', '{\"data\":\"base64image\"}', 23, NULL);
INSERT INTO `operation_logs` VALUES (18, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '用户注册', 3, 'user', '192.168.1.102', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/register', '{\"username\":\"register_user\",\"email\":\"register@example.com\"}', '{\"message\":\"注册成功\"}', 289, NULL);
INSERT INTO `operation_logs` VALUES (19, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '更新用户状态', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'PUT', '/api/v1/users/9/status', '{\"status\":0}', '{\"message\":\"状态更新成功\"}', 134, NULL);

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `updated_at` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '请求路径',
  `method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '请求方法',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '权限描述',
  `status` tinyint NULL DEFAULT 1 COMMENT '请求路径',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_permission_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of permission
-- ----------------------------
INSERT INTO `permission` VALUES (1, '', '', NULL, '/api/v1/login', 'POST', '用户登录', 1);
INSERT INTO `permission` VALUES (2, '', '', NULL, '/api/v1/captcha', 'GET', '获取验证码', 1);
INSERT INTO `permission` VALUES (3, '', '', NULL, '/api/v1/logout', 'POST', '退出登录', 1);
INSERT INTO `permission` VALUES (4, '', '', NULL, '/api/v1/users', 'POST', '创建用户', 1);
INSERT INTO `permission` VALUES (5, '', '', NULL, '/api/v1/users', 'GET', '获取用户列表', 1);
INSERT INTO `permission` VALUES (6, '', '', NULL, '/api/v1/users/:id', 'GET', '获取用户信息', 1);
INSERT INTO `permission` VALUES (7, '', '', NULL, '/api/v1/users/:id', 'PUT', '更新用户信息', 1);
INSERT INTO `permission` VALUES (8, '', '', NULL, '/api/v1/users/:id/status', 'PUT', '更新用户状态', 1);
INSERT INTO `permission` VALUES (9, '', '', NULL, '/api/v1/users/:id', 'DELETE', '删除用户', 1);
INSERT INTO `permission` VALUES (10, '', '', NULL, '/api/v1/users/change-password', 'PUT', '修改密码', 1);
INSERT INTO `permission` VALUES (11, '', '', NULL, '/api/v1/users/:id/reset-password', 'PUT', '重置密码', 1);
INSERT INTO `permission` VALUES (12, '', '', NULL, '/api/v1/users-roles/:username', 'POST', '为用户分配角色', 1);
INSERT INTO `permission` VALUES (13, '', '', NULL, '/api/v1/users-roles/:username', 'DELETE', '移除用户的角色', 1);
INSERT INTO `permission` VALUES (14, '', '', NULL, '/api/v1/users-roles/:username', 'GET', '获取用户的角色列表', 1);
INSERT INTO `permission` VALUES (15, '', '', NULL, '/api/v1/roles', 'POST', '创建角色', 1);
INSERT INTO `permission` VALUES (16, '', '', NULL, '/api/v1/roles', 'GET', '获取角色列表', 1);
INSERT INTO `permission` VALUES (17, '', '', NULL, '/api/v1/roles/:id', 'GET', '获取角色详情', 1);
INSERT INTO `permission` VALUES (18, '', '', NULL, '/api/v1/roles/:id', 'PUT', '更新角色', 1);
INSERT INTO `permission` VALUES (19, '', '', NULL, '/api/v1/roles/:id', 'DELETE', '删除角色', 1);
INSERT INTO `permission` VALUES (20, '', '', NULL, '/api/v1/menus', 'POST', '创建菜单', 1);
INSERT INTO `permission` VALUES (21, '', '', NULL, '/api/v1/menus', 'GET', '查询用户可见菜单', 1);
INSERT INTO `permission` VALUES (22, '', '', NULL, '/api/v1/menus/all', 'GET', '查询所有菜单', 1);
INSERT INTO `permission` VALUES (23, '', '', NULL, '/api/v1/menus/:id', 'PUT', '更新菜单', 1);
INSERT INTO `permission` VALUES (24, '', '', NULL, '/api/v1/menus/:id', 'DELETE', '删除菜单', 1);
INSERT INTO `permission` VALUES (25, '', '', NULL, '/api/v1/roles/:id/menus', 'POST', '为角色分配菜单权限', 1);
INSERT INTO `permission` VALUES (26, '', '', NULL, '/api/v1/roles/:id/menus', 'GET', '获取角色的菜单权限', 1);
INSERT INTO `permission` VALUES (27, '', '', NULL, '/api/v1/roles/:id/menus', 'DELETE', '移除角色的菜单权限', 1);
INSERT INTO `permission` VALUES (28, '', '', NULL, '/api/v1/roles/:id/policies', 'POST', '添加Casbin策略', 1);
INSERT INTO `permission` VALUES (29, '', '', NULL, '/api/v1/roles/:id/policies', 'DELETE', '移除Casbin策略', 1);
INSERT INTO `permission` VALUES (30, '', '', NULL, '/api/v1/roles/:id/policies', 'GET', '获取角色的Casbin策略', 1);
INSERT INTO `permission` VALUES (31, '', '', NULL, '/api/v1/permissions', 'POST', '创建权限', 1);
INSERT INTO `permission` VALUES (32, '', '', NULL, '/api/v1/permissions', 'GET', '获取权限列表', 1);
INSERT INTO `permission` VALUES (33, '', '', NULL, '/api/v1/permissions/:id', 'GET', '获取权限详情', 1);
INSERT INTO `permission` VALUES (34, '', '', NULL, '/api/v1/permissions/:id', 'PUT', '更新权限', 1);
INSERT INTO `permission` VALUES (35, '', '', NULL, '/api/v1/permissions/:id/status', 'PUT', '更新权限状态', 1);
INSERT INTO `permission` VALUES (36, '', '', NULL, '/api/v1/permissions/:id', 'DELETE', '删除权限', 1);
INSERT INTO `permission` VALUES (37, '', '', NULL, '/api/v1/permissions/all/:id', 'GET', '获取角色管理中不分页的权限列表', 1);
INSERT INTO `permission` VALUES (38, '', '', NULL, '/api/v1/permissions/all', 'GET', '获取角色管理中不分页的权限列表', 1);

-- ----------------------------
-- Table structure for role_menus
-- ----------------------------
DROP TABLE IF EXISTS `role_menus`;
CREATE TABLE `role_menus`  (
  `role_id` bigint UNSIGNED NOT NULL,
  `menu_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of role_menus
-- ----------------------------
INSERT INTO `role_menus` VALUES (1, 1);
INSERT INTO `role_menus` VALUES (1, 2);
INSERT INTO `role_menus` VALUES (1, 11);
INSERT INTO `role_menus` VALUES (1, 12);
INSERT INTO `role_menus` VALUES (1, 13);
INSERT INTO `role_menus` VALUES (1, 14);
INSERT INTO `role_menus` VALUES (1, 15);
INSERT INTO `role_menus` VALUES (2, 1);
INSERT INTO `role_menus` VALUES (2, 2);
INSERT INTO `role_menus` VALUES (2, 11);
INSERT INTO `role_menus` VALUES (2, 12);
INSERT INTO `role_menus` VALUES (2, 13);
INSERT INTO `role_menus` VALUES (2, 14);
INSERT INTO `role_menus` VALUES (2, 15);
INSERT INTO `role_menus` VALUES (3, 1);
INSERT INTO `role_menus` VALUES (4, 1);
INSERT INTO `role_menus` VALUES (5, 1);
INSERT INTO `role_menus` VALUES (5, 2);
INSERT INTO `role_menus` VALUES (5, 11);
INSERT INTO `role_menus` VALUES (5, 12);
INSERT INTO `role_menus` VALUES (5, 13);
INSERT INTO `role_menus` VALUES (5, 14);
INSERT INTO `role_menus` VALUES (5, 15);

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `updated_at` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `status` bigint NULL DEFAULT 1,
  `ident` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_roles_name`(`name` ASC) USING BTREE,
  UNIQUE INDEX `idx_roles_ident`(`ident` ASC) USING BTREE,
  INDEX `idx_roles_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO `roles` VALUES (1, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', NULL, '超级管理员', '拥有系统所有权限', 1, 'surperadmin');
INSERT INTO `roles` VALUES (2, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', NULL, '管理员', '拥有大部分管理权限', 1, 'admin');
INSERT INTO `roles` VALUES (3, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', NULL, '普通用户', '拥有基本查看权限', 1, 'user');
INSERT INTO `roles` VALUES (4, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', NULL, '访客', '仅拥有查看权限', 1, 'Visitor');
INSERT INTO `roles` VALUES (5, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', NULL, '测试角色', '用于测试的角色', 1, 'ceshi');
INSERT INTO `roles` VALUES (7, '', '', NULL, 'TestRole', '', 1, 'test_role_ident');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `updated_at` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `status` bigint NULL DEFAULT 1,
  `last_login_at` datetime(3) NULL DEFAULT NULL,
  `last_login_ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `role_id` bigint UNSIGNED NULL DEFAULT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `ident` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_users_username`(`username` ASC) USING BTREE,
  UNIQUE INDEX `idx_users_email`(`email` ASC) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_roles_users`(`role_id` ASC) USING BTREE,
  CONSTRAINT `fk_roles_users` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_users_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (10, '2026-01-18 19:07:49.000', '2026-01-20 10:04:26.130', NULL, 'admin', '$2a$10$ocZQAxtwX0K8aSywcLjICeMqyWv8KqvqJ7ZeoKSu9bMSIKjKhpkAq', 'admin2@example.com', '13800138009', '管理员', '', 1, '2026-01-28 22:52:57.493', '', 2, '', '', NULL);
INSERT INTO `users` VALUES (16, '', '', '2026-01-27 21:18:43.440', 'ceshi', '$2a$10$3J.aR7claboqAQuKPWeWUenxV.jY7Hjasi.eBBqYicedQgyzWG4KO', '2603485744@qq.com', '', '测试', '', 1, '2026-01-27 20:59:45.972', '', 5, '', '', NULL);
INSERT INTO `users` VALUES (19, '', '', '2026-01-27 21:18:41.641', 'ceshi2', '$2a$10$lGiFopGyYy0A707gBLzwv.e7EDInuygQg/8zu.Y2Mbm78dyCM56rW', '2603485766@qq.com', '', '测试2', '', 1, NULL, '', 3, '', '', NULL);
INSERT INTO `users` VALUES (20, '', '', NULL, 'ceshiadmin', '$2a$10$zGgDA7YfYEHet7tbY9d1EePymBCGNgpVNQwgz4fw63.BqAe/9W.Xy', '26034854@qq.com', '', '测试admin', '', 1, NULL, '', 1, '', '', NULL);

SET FOREIGN_KEY_CHECKS = 1;
