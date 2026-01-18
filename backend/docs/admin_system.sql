/*
 Navicat Premium Dump SQL

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80027 (8.0.27)
 Source Host           : localhost:3306
 Source Schema         : admin_system

 Target Server Type    : MySQL
 Target Server Version : 80027 (8.0.27)
 File Encoding         : 65001

 Date: 18/01/2026 22:17:43
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
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (2, 'p', 'role_4', '/api/users', 'GET', '', '', '');

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `updated_at` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `redirect` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `parent_id` bigint UNSIGNED NULL DEFAULT NULL,
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `sort` bigint NULL DEFAULT 0,
  `is_hidden` tinyint(1) NULL DEFAULT 0,
  `is_link` tinyint(1) NULL DEFAULT 0,
  `link_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `status` bigint NULL DEFAULT 1,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_menus_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_menus_children`(`parent_id` ASC) USING BTREE,
  CONSTRAINT `fk_menus_children` FOREIGN KEY (`parent_id`) REFERENCES `menus` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menus
-- ----------------------------
INSERT INTO `menus` VALUES (1, '2026-01-18 21:54:50', '2026-01-18 21:54:50', NULL, 'dashboard', '控制台', '/dashboard', 'DashboardView', NULL, NULL, 'House', 1, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (2, '2026-01-18 21:54:50', '2026-01-18 21:54:50', NULL, 'system', '系统管理', '/system', '', NULL, NULL, 'Setting', 100, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (11, '2026-01-18 21:54:50', '2026-01-18 21:54:50', NULL, 'users', '用户管理', '/users', 'UserManageView', NULL, 2, 'User', 1, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (12, '2026-01-18 21:54:50', '2026-01-18 21:54:50', NULL, 'roles', '角色管理', '/roles', 'RoleManageView', NULL, 2, 'Avatar', 2, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (13, '2026-01-18 21:54:50', '2026-01-18 21:54:50', NULL, 'menus', '菜单管理', '/menus', 'MenuManageView', NULL, 2, 'Menu', 3, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (14, '2026-01-18 21:54:50', '2026-01-18 21:54:50', NULL, 'operation-logs', '操作日志', '/operation-logs', 'OperationLogView', NULL, 2, 'Document', 4, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (15, '2026-01-18 21:54:50', '2026-01-18 21:54:50', NULL, 'permissions', '权限管理', '/permissions', 'PermissionManageView', NULL, 2, 'Lock', 5, 0, 0, NULL, 1);

-- ----------------------------
-- Table structure for operation_logs
-- ----------------------------
DROP TABLE IF EXISTS `operation_logs`;
CREATE TABLE `operation_logs`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of operation_logs
-- ----------------------------
INSERT INTO `operation_logs` VALUES (1, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '用户登录', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/login', '{\"username\":\"admin\",\"password\":\"***\"}', '{\"message\":\"登录成功\"}', 156);
INSERT INTO `operation_logs` VALUES (2, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '用户登录', 2, 'manager', '192.168.1.101', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/login', '{\"username\":\"manager\",\"password\":\"***\"}', '{\"message\":\"登录成功\"}', 142);
INSERT INTO `operation_logs` VALUES (3, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '查询用户列表', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/users', '', '{\"data\":[...]}', 89);
INSERT INTO `operation_logs` VALUES (4, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '创建用户', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/users', '{\"username\":\"newuser\",\"email\":\"new@example.com\"}', '{\"message\":\"创建成功\"}', 234);
INSERT INTO `operation_logs` VALUES (5, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '更新用户信息', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'PUT', '/api/v1/users/6', '{\"nickname\":\"张三更新\"}', '{\"message\":\"更新成功\"}', 178);
INSERT INTO `operation_logs` VALUES (6, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '删除用户', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'DELETE', '/api/v1/users/99', '', '{\"message\":\"删除成功\"}', 145);
INSERT INTO `operation_logs` VALUES (7, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '查询角色列表', 2, 'manager', '192.168.1.101', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/roles', '', '{\"data\":[...]}', 76);
INSERT INTO `operation_logs` VALUES (8, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '创建角色', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/roles', '{\"name\":\"新角色\",\"description\":\"测试角色\"}', '{\"message\":\"创建成功\"}', 198);
INSERT INTO `operation_logs` VALUES (9, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '分配角色权限', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'PUT', '/api/v1/roles/5/menus', '{\"menu_ids\":[1,2,11]}', '{\"message\":\"分配成功\"}', 267);
INSERT INTO `operation_logs` VALUES (10, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '查询菜单树', 3, 'user', '192.168.1.102', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/menus', '', '{\"data\":[...]}', 112);
INSERT INTO `operation_logs` VALUES (11, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '创建菜单', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/menus', '{\"name\":\"test\",\"title\":\"测试菜单\"}', '{\"message\":\"创建成功\"}', 189);
INSERT INTO `operation_logs` VALUES (12, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '修改密码', 2, 'manager', '192.168.1.101', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'PUT', '/api/v1/users/change-password', '{\"old_password\":\"***\",\"new_password\":\"***\"}', '{\"message\":\"修改成功\"}', 156);
INSERT INTO `operation_logs` VALUES (13, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '查询操作日志', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/operation-logs', '', '{\"data\":[...]}', 134);
INSERT INTO `operation_logs` VALUES (14, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '删除操作日志', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'DELETE', '/api/v1/operation-logs/1', '', '{\"message\":\"删除成功\"}', 98);
INSERT INTO `operation_logs` VALUES (15, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '用户登录失败', 9, 'disabled', '192.168.1.109', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 0, 'POST', '/api/v1/login', '{\"username\":\"disabled\",\"password\":\"***\"}', '{\"error\":\"用户已被禁用\"}', 67);
INSERT INTO `operation_logs` VALUES (16, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '查询用户详情', 3, 'user', '192.168.1.102', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/users/3', '', '{\"data\":{...}}', 45);
INSERT INTO `operation_logs` VALUES (17, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '获取验证码', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'GET', '/api/v1/captcha', '', '{\"data\":\"base64image\"}', 23);
INSERT INTO `operation_logs` VALUES (18, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '用户注册', 3, 'user', '192.168.1.102', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'POST', '/api/v1/register', '{\"username\":\"register_user\",\"email\":\"register@example.com\"}', '{\"message\":\"注册成功\"}', 289);
INSERT INTO `operation_logs` VALUES (19, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', '更新用户状态', 1, 'admin', '192.168.1.100', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36', 1, 'PUT', '/api/v1/users/9/status', '{\"status\":0}', '{\"message\":\"状态更新成功\"}', 134);

-- ----------------------------
-- Table structure for role_menus
-- ----------------------------
DROP TABLE IF EXISTS `role_menus`;
CREATE TABLE `role_menus`  (
  `role_id` bigint UNSIGNED NOT NULL,
  `menu_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

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
INSERT INTO `role_menus` VALUES (3, 1);
INSERT INTO `role_menus` VALUES (3, 2);
INSERT INTO `role_menus` VALUES (3, 11);
INSERT INTO `role_menus` VALUES (4, 1);

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
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_roles_name`(`name` ASC) USING BTREE,
  INDEX `idx_roles_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO `roles` VALUES (1, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', NULL, '超级管理员', '拥有系统所有权限', 1);
INSERT INTO `roles` VALUES (2, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', NULL, '管理员', '拥有大部分管理权限', 1);
INSERT INTO `roles` VALUES (3, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', NULL, '普通用户', '拥有基本查看权限', 1);
INSERT INTO `roles` VALUES (4, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', NULL, '访客', '仅拥有查看权限', 1);
INSERT INTO `roles` VALUES (5, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', NULL, '测试角色', '用于测试的角色', 1);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
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
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_users_username`(`username` ASC) USING BTREE,
  UNIQUE INDEX `idx_users_email`(`email` ASC) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_roles_users`(`role_id` ASC) USING BTREE,
  CONSTRAINT `fk_roles_users` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '2026-01-18 19:07:49.000', '2026-01-18 20:41:08.778', '2026-01-18 21:02:33.077', 'testuser_1768740068', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'testuser_1768740068_updated@example.com', '', '测试用户（已更新）', '', 0, '2026-01-18 20:41:08.689', '', 1);
INSERT INTO `users` VALUES (2, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', NULL, 'manager', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'manager@example.com', '13800138001', '管理员', '', 1, NULL, NULL, 2);
INSERT INTO `users` VALUES (3, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', NULL, 'user', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'user@example.com', '13800138002', '普通用户', '', 1, NULL, NULL, 3);
INSERT INTO `users` VALUES (4, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', NULL, 'guest', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'guest@example.com', '13800138003', '访客', '', 1, NULL, NULL, 4);
INSERT INTO `users` VALUES (5, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', NULL, 'test', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'test@example.com', '13800138004', '测试用户', '', 1, NULL, NULL, 5);
INSERT INTO `users` VALUES (6, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', NULL, 'zhangsan', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'zhangsan@example.com', '13800138005', '张三', '', 1, NULL, NULL, 3);
INSERT INTO `users` VALUES (7, '2026-01-18 19:07:49.000', '2026-01-18 19:07:49.000', NULL, 'lisi', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'lisi@example.com', '13800138006', '李四', '', 1, NULL, NULL, 3);
INSERT INTO `users` VALUES (8, '2026-01-18 19:07:49.000', '2026-01-18 19:22:14.965', NULL, 'wangwu', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', 'wangwu@example.com', '13800138007', '王五', '', 1, NULL, NULL, 3);
INSERT INTO `users` VALUES (9, '2026-01-18 19:07:49.000', '2026-01-18 19:24:28.961', NULL, 'disabled', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', '132@example.com', '123', '禁用用户', '', 1, NULL, '', 3);
INSERT INTO `users` VALUES (10, '2026-01-18 19:07:49.000', '2026-01-18 21:55:15.651', NULL, 'admin', '$2a$10$ocZQAxtwX0K8aSywcLjICeMqyWv8KqvqJ7ZeoKSu9bMSIKjKhpkAq', 'admin2@example.com', '13800138009', '管理员2', '', 1, '2026-01-18 21:55:15.650', '', 2);
INSERT INTO `users` VALUES (14, '2026-01-18 19:20:14.750', '2026-01-18 19:20:14.750', '2026-01-18 19:20:34.435', 'ceshi', '$2a$10$UVGAcMTvIU6xWteku4lBdOrw8OMTThhAh0PBltB50aruK2K6ogPCy', '2603485744@qq.com', '', '测试', '', 1, NULL, '', 2);

SET FOREIGN_KEY_CHECKS = 1;
