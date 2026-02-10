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

 Date: 10/02/2026 22:50:03
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
) ENGINE = InnoDB AUTO_INCREMENT = 127 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (16, 'g', 'admin', 'admin', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (20, 'g', 'ceshi', 'ceshi', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (19, 'g', 'ceshi2', 'user', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (91, 'g', 'ceshiadmin', 'surperadmin', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (15, 'g', 'superadmin', 'surperadmin', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (56, 'p', 'admin', '/api/v1/captcha', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (121, 'p', 'admin', '/api/v1/credentials', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (120, 'p', 'admin', '/api/v1/credentials', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (124, 'p', 'admin', '/api/v1/credentials/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (122, 'p', 'admin', '/api/v1/credentials/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (123, 'p', 'admin', '/api/v1/credentials/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (125, 'p', 'admin', '/api/v1/credentials/batch', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (126, 'p', 'admin', '/api/v1/credentials/host', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (109, 'p', 'admin', '/api/v1/host-groups', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (108, 'p', 'admin', '/api/v1/host-groups', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (112, 'p', 'admin', '/api/v1/host-groups/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (110, 'p', 'admin', '/api/v1/host-groups/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (111, 'p', 'admin', '/api/v1/host-groups/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (113, 'p', 'admin', '/api/v1/host-groups/:id/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (114, 'p', 'admin', '/api/v1/host-metrics', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (115, 'p', 'admin', '/api/v1/host-metrics/history', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (116, 'p', 'admin', '/api/v1/host-metrics/latest', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (101, 'p', 'admin', '/api/v1/hosts', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (100, 'p', 'admin', '/api/v1/hosts', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (104, 'p', 'admin', '/api/v1/hosts/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (102, 'p', 'admin', '/api/v1/hosts/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (103, 'p', 'admin', '/api/v1/hosts/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (107, 'p', 'admin', '/api/v1/hosts/:id/monitoring', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (106, 'p', 'admin', '/api/v1/hosts/:id/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (105, 'p', 'admin', '/api/v1/hosts/batch', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (117, 'p', 'admin', '/api/v1/hosts/statistics', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (55, 'p', 'admin', '/api/v1/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (57, 'p', 'admin', '/api/v1/logout', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (75, 'p', 'admin', '/api/v1/menus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (74, 'p', 'admin', '/api/v1/menus', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (78, 'p', 'admin', '/api/v1/menus/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (77, 'p', 'admin', '/api/v1/menus/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (76, 'p', 'admin', '/api/v1/menus/all', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (98, 'p', 'admin', '/api/v1/operation-logs', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (118, 'p', 'admin', '/api/v1/operation-logs/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (86, 'p', 'admin', '/api/v1/permissions', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (85, 'p', 'admin', '/api/v1/permissions', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (90, 'p', 'admin', '/api/v1/permissions/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (87, 'p', 'admin', '/api/v1/permissions/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (88, 'p', 'admin', '/api/v1/permissions/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (89, 'p', 'admin', '/api/v1/permissions/:id/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (95, 'p', 'admin', '/api/v1/permissions/all', 'GET', '', '', '');
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
INSERT INTO `casbin_rule` VALUES (119, 'p', 'admin', '/api/v1/users/profile', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (22, 'p', 'surperadmin', '/api/v1/captcha', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (21, 'p', 'surperadmin', '/api/v1/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (23, 'p', 'surperadmin', '/api/v1/logout', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (39, 'p', 'surperadmin', '/api/v1/menus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (38, 'p', 'surperadmin', '/api/v1/menus', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (42, 'p', 'surperadmin', '/api/v1/menus/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (41, 'p', 'surperadmin', '/api/v1/menus/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (40, 'p', 'surperadmin', '/api/v1/menus/all', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (96, 'p', 'surperadmin', '/api/v1/operation-logs', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (97, 'p', 'surperadmin', '/api/v1/operation-logs/stats', 'GET', '', '', '');
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
-- Table structure for credentials
-- ----------------------------
DROP TABLE IF EXISTS `credentials`;
CREATE TABLE `credentials`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '凭据唯一标识符',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '凭据名称，用于标识和搜索',
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '登录用户名',
  `password` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '加密后的密码',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '凭据描述信息',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '凭据状态：1-启用，0-禁用',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_name`(`name` ASC) USING BTREE COMMENT '凭据名称唯一索引',
  INDEX `idx_status`(`status` ASC) USING BTREE COMMENT '状态索引，用于快速筛选启用/禁用的凭据',
  INDEX `idx_username`(`username` ASC) USING BTREE COMMENT '用户名索引，便于按用户名查询'
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '主机凭据信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of credentials
-- ----------------------------
INSERT INTO `credentials` VALUES (1, 'root-default', 'root', '$2a$10$example_hashed_password', '默认root用户凭据', 1, '2026-02-06 05:54:01', '2026-02-06 05:54:01', NULL);
INSERT INTO `credentials` VALUES (2, 'admin-default', 'admin', '$2a$10$example_hashed_password', '默认管理员凭据', 1, '2026-02-06 05:54:01', '2026-02-06 05:54:01', NULL);
INSERT INTO `credentials` VALUES (3, 'ubuntu-default', 'ubuntu', '$2a$10$example_hashed_password', 'Ubuntu系统默认用户凭据', 1, '2026-02-06 05:54:01', '2026-02-06 05:54:01', NULL);
INSERT INTO `credentials` VALUES (4, 'centos-default', 'centos', '$2a$10$example_hashed_password', 'CentOS系统默认用户凭据', 1, '2026-02-06 05:54:01', '2026-02-06 05:54:01', NULL);
INSERT INTO `credentials` VALUES (5, '测试', 'root', 'Zhy20250730!', '密码', 1, '2026-02-10 22:46:37', '2026-02-10 22:46:37', NULL);

-- ----------------------------
-- Table structure for host_credentials
-- ----------------------------
DROP TABLE IF EXISTS `host_credentials`;
CREATE TABLE `host_credentials`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '关联关系唯一标识符',
  `host_id` bigint UNSIGNED NOT NULL COMMENT '主机ID，关联hosts表',
  `credential_id` bigint UNSIGNED NOT NULL COMMENT '凭据ID，关联credentials表',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_host_credential`(`host_id` ASC, `credential_id` ASC) USING BTREE COMMENT '主机-凭据组合唯一约束，防止重复关联',
  INDEX `idx_host_id`(`host_id` ASC) USING BTREE COMMENT '主机ID索引，便于查询某主机的所有凭据',
  INDEX `idx_credential_id`(`credential_id` ASC) USING BTREE COMMENT '凭据ID索引，便于查询使用某凭据的所有主机',
  CONSTRAINT `fk_host_credentials_credential_id` FOREIGN KEY (`credential_id`) REFERENCES `credentials` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_host_credentials_host_id` FOREIGN KEY (`host_id`) REFERENCES `hosts` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '主机与凭据关联关系表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of host_credentials
-- ----------------------------
INSERT INTO `host_credentials` VALUES (12, 2, 1, '2026-02-10 14:32:31', NULL, NULL);
INSERT INTO `host_credentials` VALUES (13, 2, 2, '2026-02-10 14:32:31', NULL, NULL);
INSERT INTO `host_credentials` VALUES (26, 1, 1, '2026-02-10 14:45:57', NULL, NULL);
INSERT INTO `host_credentials` VALUES (27, 1, 2, '2026-02-10 14:45:57', NULL, NULL);

-- ----------------------------
-- Table structure for host_groups
-- ----------------------------
DROP TABLE IF EXISTS `host_groups`;
CREATE TABLE `host_groups`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '主机组名称',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '描述信息',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态: 1-启用, 0-禁用',
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建人用户ID',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新人用户ID',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_host_groups_name`(`name` ASC) USING BTREE,
  INDEX `idx_host_groups_status`(`status` ASC) USING BTREE,
  INDEX `idx_host_groups_created_at`(`created_at` ASC) USING BTREE,
  INDEX `idx_host_groups_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '主机组表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of host_groups
-- ----------------------------
INSERT INTO `host_groups` VALUES (1, 'Web服务器组', '用于存放Web应用服务器的主机组', 1, 1, 1, '2026-02-05 04:11:27', '2026-02-05 04:11:27', NULL);
INSERT INTO `host_groups` VALUES (2, '数据库服务器组', '用于存放数据库服务器的主机组', 1, 1, 1, '2026-02-05 04:11:27', '2026-02-05 04:11:27', NULL);
INSERT INTO `host_groups` VALUES (3, '缓存服务器组', '用于存放Redis、Memcached等缓存服务器的主机组', 1, 1, 1, '2026-02-05 04:11:27', '2026-02-05 04:11:27', NULL);
INSERT INTO `host_groups` VALUES (4, '存储服务器组', '用于存放文件存储和备份服务器的主机组', 1, 1, 1, '2026-02-05 04:11:27', '2026-02-05 04:11:27', NULL);
INSERT INTO `host_groups` VALUES (5, 'BBB应用服务器组', 'AAA应用服务器主机组', 1, 10, 10, '2026-02-05 13:51:54', '2026-02-05 14:06:49', NULL);

-- ----------------------------
-- Table structure for host_metrics
-- ----------------------------
DROP TABLE IF EXISTS `host_metrics`;
CREATE TABLE `host_metrics`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `host_id` bigint UNSIGNED NOT NULL COMMENT '主机ID',
  `metric_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '指标类型: cpu,memory,disk,network',
  `metric_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '指标名称',
  `metric_value` decimal(10, 2) NOT NULL COMMENT '指标值',
  `unit` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '单位',
  `recorded_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_host_metrics_host_id`(`host_id` ASC) USING BTREE,
  INDEX `idx_host_metrics_metric_type`(`metric_type` ASC) USING BTREE,
  INDEX `idx_host_metrics_recorded_at`(`recorded_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '主机监控指标表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of host_metrics
-- ----------------------------
INSERT INTO `host_metrics` VALUES (1, 1, 'cpu', 'cpu_usage', 45.50, '%', '2026-02-05 04:06:27');
INSERT INTO `host_metrics` VALUES (2, 1, 'cpu', 'cpu_usage', 48.20, '%', '2026-02-05 04:07:27');
INSERT INTO `host_metrics` VALUES (3, 1, 'cpu', 'cpu_usage', 42.80, '%', '2026-02-05 04:08:27');
INSERT INTO `host_metrics` VALUES (4, 1, 'cpu', 'cpu_usage', 51.30, '%', '2026-02-05 04:09:27');
INSERT INTO `host_metrics` VALUES (5, 1, 'cpu', 'cpu_usage', 47.90, '%', '2026-02-05 04:10:27');
INSERT INTO `host_metrics` VALUES (6, 1, 'memory', 'memory_usage', 68.20, '%', '2026-02-05 04:06:27');
INSERT INTO `host_metrics` VALUES (7, 1, 'memory', 'memory_usage', 71.50, '%', '2026-02-05 04:07:27');
INSERT INTO `host_metrics` VALUES (8, 1, 'memory', 'memory_usage', 65.80, '%', '2026-02-05 04:08:27');
INSERT INTO `host_metrics` VALUES (9, 1, 'memory', 'memory_usage', 73.10, '%', '2026-02-05 04:09:27');
INSERT INTO `host_metrics` VALUES (10, 1, 'memory', 'memory_usage', 69.70, '%', '2026-02-05 04:10:27');
INSERT INTO `host_metrics` VALUES (11, 4, 'disk', 'disk_usage', 78.50, '%', '2026-02-05 04:06:27');
INSERT INTO `host_metrics` VALUES (12, 4, 'disk', 'disk_usage', 79.20, '%', '2026-02-05 04:07:27');
INSERT INTO `host_metrics` VALUES (13, 4, 'disk', 'disk_usage', 80.10, '%', '2026-02-05 04:08:27');
INSERT INTO `host_metrics` VALUES (14, 4, 'disk', 'disk_usage', 81.30, '%', '2026-02-05 04:09:27');
INSERT INTO `host_metrics` VALUES (15, 4, 'disk', 'disk_usage', 82.70, '%', '2026-02-05 04:10:27');
INSERT INTO `host_metrics` VALUES (16, 7, 'network', 'network_in', 125.60, 'Mbps', '2026-02-05 04:06:27');
INSERT INTO `host_metrics` VALUES (17, 7, 'network', 'network_in', 132.40, 'Mbps', '2026-02-05 04:07:27');
INSERT INTO `host_metrics` VALUES (18, 7, 'network', 'network_in', 118.90, 'Mbps', '2026-02-05 04:08:27');
INSERT INTO `host_metrics` VALUES (19, 7, 'network', 'network_in', 145.20, 'Mbps', '2026-02-05 04:09:27');
INSERT INTO `host_metrics` VALUES (20, 7, 'network', 'network_in', 138.70, 'Mbps', '2026-02-05 04:10:27');

-- ----------------------------
-- Table structure for hosts
-- ----------------------------
DROP TABLE IF EXISTS `hosts`;
CREATE TABLE `hosts`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `hostname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '主机名',
  `ip_address` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'IP地址(IPv4或IPv6)',
  `port` smallint UNSIGNED NOT NULL DEFAULT 22 COMMENT 'SSH端口',
  `os_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'linux' COMMENT '操作系统类型: linux,windows',
  `cpu_cores` smallint UNSIGNED NULL DEFAULT NULL COMMENT 'CPU核心数',
  `memory_gb` smallint UNSIGNED NULL DEFAULT NULL COMMENT '内存大小(GB)',
  `disk_space_gb` int UNSIGNED NULL DEFAULT NULL COMMENT '磁盘空间(GB)',
  `group_id` bigint UNSIGNED NOT NULL COMMENT '所属主机组ID',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '主机状态: 1-在线, 2-离线, 3-故障',
  `monitoring_enabled` tinyint NOT NULL DEFAULT 1 COMMENT '监控是否启用: 1-启用, 2-禁用',
  `last_heartbeat` timestamp NULL DEFAULT NULL COMMENT '最后心跳时间',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '主机描述',
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建人用户ID',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新人用户ID',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_hosts_hostname`(`hostname` ASC) USING BTREE,
  UNIQUE INDEX `uk_hosts_ip_address`(`ip_address` ASC) USING BTREE,
  INDEX `idx_hosts_group_id`(`group_id` ASC) USING BTREE,
  INDEX `idx_hosts_status`(`status` ASC) USING BTREE,
  INDEX `idx_hosts_os_type`(`os_type` ASC) USING BTREE,
  INDEX `idx_hosts_monitoring_enabled`(`monitoring_enabled` ASC) USING BTREE,
  INDEX `idx_hosts_last_heartbeat`(`last_heartbeat` ASC) USING BTREE,
  INDEX `idx_hosts_created_at`(`created_at` ASC) USING BTREE,
  INDEX `idx_hosts_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '主机表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of hosts
-- ----------------------------
INSERT INTO `hosts` VALUES (1, 'web-server-01', '192.168.1.101', 22, 'linux', 8, 16, 500, 1, 2, 1, '2026-02-05 04:06:27', '生产环境Web服务器01', 1, 10, '2026-02-05 04:11:27', '2026-02-10 22:45:57', NULL);
INSERT INTO `hosts` VALUES (2, 'web-server-02', '192.168.1.102', 22, 'linux', 8, 16, 500, 1, 1, 1, '2026-02-05 04:08:27', '生产环境Web服务器02', 1, 10, '2026-02-05 04:11:27', '2026-02-10 22:32:32', NULL);

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
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
) ENGINE = InnoDB AUTO_INCREMENT = 121 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of menus
-- ----------------------------
INSERT INTO `menus` VALUES (1, NULL, NULL, NULL, 'dashboard', '控制台', '/dashboard', 'DashboardView', NULL, 0, 'House', 1, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (2, NULL, NULL, NULL, 'system', '系统管理', '/system', '', NULL, 0, 'Setting', 100, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (3, '2026-02-06 11:20:09.305', '2026-02-06 11:20:09.305', NULL, 'asset', '主机管理', '/asset', '', NULL, 0, 'Setting', 1, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (11, NULL, NULL, NULL, 'users', '用户管理', '/users', 'UserManageView', NULL, 2, 'User', 1, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (12, NULL, NULL, NULL, 'roles', '角色管理', '/roles', 'RoleManageView', NULL, 2, 'Avatar', 2, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (13, NULL, NULL, NULL, 'menus', '菜单管理', '/menus', 'MenuManageView', NULL, 2, 'Menu', 3, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (14, NULL, NULL, NULL, 'operation-logs', '操作日志', '/operation-logs', 'OperationLogView', NULL, 2, 'Document', 4, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (15, NULL, NULL, NULL, 'permissions', '权限管理', '/permissions', 'PermissionResourceView', NULL, 2, 'Lock', 5, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (118, NULL, NULL, NULL, 'group', '主机组管理', '/asset/group', 'HostGroupView', NULL, 3, 'User', 1, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (119, NULL, NULL, NULL, 'host', '主机管理', '/asset/host', 'HostManageView', NULL, 3, 'User', 2, 0, 0, NULL, 1);
INSERT INTO `menus` VALUES (120, '2026-02-10 19:26:14.630', '2026-02-10 19:26:14.630', NULL, 'credential', '凭据管理', '/asset/credential', 'CredentialManageView', NULL, 3, 'User', 3, 0, 0, NULL, 1);

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
  `status` bigint NULL DEFAULT 0,
  `request_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `request_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `request_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `response_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `response_time` bigint NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `module` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `level` bigint NULL DEFAULT 1,
  `refer_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_operation_logs_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `idx_operation_logs_refer_id`(`refer_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 111 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of operation_logs
-- ----------------------------
INSERT INTO `operation_logs` VALUES (20, '2026-02-03 21:33:22.000', '2026-02-03 21:33:22.000', '删除用户', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'DELETE', '/api/v1/users/:id', '', '{\"message\":\"删除成功\"}', 7, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (21, '2026-02-04 11:08:37.000', '2026-02-04 11:08:37.000', '创建用户', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/users', '{\"username\":\"1111111\",\"nickname\":\"1111111\",\"email\":\"1111111@123.com\",\"phone\":\"111111\",\"status\":1,\"password\":\"111111\"}', '{\"data\":{\"id\":23,\"created_at\":\"2026-02-04T11:08:37.143+08:00\",\"updated_at\":\"2026-02-04T11:08:37.143+08:00\",\"deleted_at\":null,\"username\":\"1111111\",\"password\":\"$2a$10$SstBKyW61Anu8rt7zwgSRuPSHFq.hqhYdpOR66jZ1dcUIXSDRVLI6\",\"email\":\"1111111@123.com\",\"phone\":\"\",\"nickname\":\"1111111\",\"avatar\":\"\",\"status\":1,\"last_login_at\":null,\"last_login_ip\":\"\",\"role_id\":null},\"message\":\"用户创建成功\"}', 179, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (22, '2026-02-05 11:31:14.000', '2026-02-05 11:31:14.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/hosts\",\"method\":\"POST\"}', '{\"message\":\"策略添加成功\"}', 118, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (23, '2026-02-05 11:31:14.000', '2026-02-05 11:31:14.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/hosts\",\"method\":\"GET\"}', '{\"message\":\"策略添加成功\"}', 93, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (24, '2026-02-05 11:31:14.000', '2026-02-05 11:31:14.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/hosts/:id\",\"method\":\"GET\"}', '{\"message\":\"策略添加成功\"}', 90, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (25, '2026-02-05 11:31:14.000', '2026-02-05 11:31:14.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/hosts/:id\",\"method\":\"PUT\"}', '{\"message\":\"策略添加成功\"}', 92, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (26, '2026-02-05 11:31:14.000', '2026-02-05 11:31:14.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/hosts/:id\",\"method\":\"DELETE\"}', '{\"message\":\"策略添加成功\"}', 94, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (27, '2026-02-05 11:31:14.000', '2026-02-05 11:31:14.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/hosts/batch\",\"method\":\"DELETE\"}', '{\"message\":\"策略添加成功\"}', 92, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (28, '2026-02-05 11:31:14.000', '2026-02-05 11:31:14.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/hosts/:id/status\",\"method\":\"PUT\"}', '{\"message\":\"策略添加成功\"}', 93, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (29, '2026-02-05 11:31:14.000', '2026-02-05 11:31:14.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/hosts/:id/monitoring\",\"method\":\"PUT\"}', '{\"message\":\"策略添加成功\"}', 97, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (30, '2026-02-05 11:31:14.000', '2026-02-05 11:31:14.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/host-groups\",\"method\":\"POST\"}', '{\"message\":\"策略添加成功\"}', 96, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (31, '2026-02-05 11:31:14.000', '2026-02-05 11:31:14.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/host-groups\",\"method\":\"GET\"}', '{\"message\":\"策略添加成功\"}', 94, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (32, '2026-02-05 11:31:15.000', '2026-02-05 11:31:15.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/host-groups/:id\",\"method\":\"GET\"}', '{\"message\":\"策略添加成功\"}', 91, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (33, '2026-02-05 11:31:15.000', '2026-02-05 11:31:15.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/host-groups/:id\",\"method\":\"PUT\"}', '{\"message\":\"策略添加成功\"}', 90, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (34, '2026-02-05 11:31:15.000', '2026-02-05 11:31:15.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/host-groups/:id\",\"method\":\"DELETE\"}', '{\"message\":\"策略添加成功\"}', 89, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (35, '2026-02-05 11:31:15.000', '2026-02-05 11:31:15.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/host-groups/:id/status\",\"method\":\"PUT\"}', '{\"message\":\"策略添加成功\"}', 91, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (36, '2026-02-05 11:31:15.000', '2026-02-05 11:31:15.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/host-metrics\",\"method\":\"POST\"}', '{\"message\":\"策略添加成功\"}', 93, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (37, '2026-02-05 11:31:15.000', '2026-02-05 11:31:15.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/host-metrics/history\",\"method\":\"GET\"}', '{\"message\":\"策略添加成功\"}', 90, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (38, '2026-02-05 11:31:15.000', '2026-02-05 11:31:15.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/host-metrics/latest\",\"method\":\"GET\"}', '{\"message\":\"策略添加成功\"}', 91, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (39, '2026-02-05 11:31:15.000', '2026-02-05 11:31:15.000', '添加Casbin策略', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/hosts/statistics\",\"method\":\"GET\"}', '{\"message\":\"策略添加成功\"}', 94, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (40, '2026-02-05 12:59:31.454', '2026-02-05 12:59:31.454', '创建主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'POST', '/api/v1/hosts', '{\r\n    \"hostname\": \"web-server-01\",\r\n    \"ip_address\": \"192.168.1.100\",\r\n    \"port\": 22,\r\n    \"username\": \"root\",\r\n    \"password\": \"encrypted_password\",\r\n    \"os_type\": \"linux\",\r\n    \"cpu_cores\": 8,\r\n    \"memory_gb\": 16,\r\n    \"disk_space_gb\": 500,\r\n    \"group_id\": 1,\r\n    \"description\": \"Web服务器01\"\r\n}', '{\"success\":false,\"message\":\"数据验证失败\",\"error\":\"数据验证失败\"}', 6, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (41, '2026-02-05 12:59:52.922', '2026-02-05 12:59:52.922', '创建主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'POST', '/api/v1/hosts', '{\r\n    \"hostname\": \"web-server-01\",\r\n    \"ip_address\": \"10.100.1.100\",\r\n    \"port\": 22,\r\n    \"username\": \"root\",\r\n    \"password\": \"encrypted_password\",\r\n    \"os_type\": \"linux\",\r\n    \"cpu_cores\": 8,\r\n    \"memory_gb\": 16,\r\n    \"disk_space_gb\": 500,\r\n    \"group_id\": 1,\r\n    \"description\": \"Web服务器01\"\r\n}', '{\"success\":false,\"message\":\"数据验证失败\",\"error\":\"数据验证失败\"}', 3, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (42, '2026-02-05 13:00:03.412', '2026-02-05 13:00:03.412', '创建主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 500, 'POST', '/api/v1/hosts', '{\r\n    \"hostname\": \"数据库\",\r\n    \"ip_address\": \"10.100.1.100\",\r\n    \"port\": 22,\r\n    \"username\": \"root\",\r\n    \"password\": \"encrypted_password\",\r\n    \"os_type\": \"linux\",\r\n    \"cpu_cores\": 8,\r\n    \"memory_gb\": 16,\r\n    \"disk_space_gb\": 500,\r\n    \"group_id\": 1,\r\n    \"description\": \"Web服务器01\"\r\n}', '{\"success\":false,\"message\":\"数据库操作失败\",\"error\":\"数据库操作失败\"}', 238, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (43, '2026-02-05 13:00:18.986', '2026-02-05 13:00:18.986', '创建主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 500, 'POST', '/api/v1/hosts', '{\r\n    \"hostname\": \"database\",\r\n    \"ip_address\": \"10.100.1.100\",\r\n    \"port\": 22,\r\n    \"username\": \"root\",\r\n    \"password\": \"encrypted_password\",\r\n    \"os_type\": \"linux\",\r\n    \"cpu_cores\": 8,\r\n    \"memory_gb\": 16,\r\n    \"disk_space_gb\": 500,\r\n    \"group_id\": 1,\r\n    \"description\": \"Web服务器01\"\r\n}', '{\"success\":false,\"message\":\"数据库操作失败\",\"error\":\"数据库操作失败\"}', 211, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (44, '2026-02-05 13:06:14.043', '2026-02-05 13:06:14.043', '创建主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 500, 'POST', '/api/v1/hosts', '{\r\n    \"hostname\": \"database\",\r\n    \"ip_address\": \"10.100.1.100\",\r\n    \"port\": 22,\r\n    \"username\": \"root\",\r\n    \"password\": \"encrypted_password\",\r\n    \"os_type\": \"linux\",\r\n    \"cpu_cores\": 8,\r\n    \"memory_gb\": 16,\r\n    \"disk_space_gb\": 500,\r\n    \"group_id\": 1,\r\n    \"description\": \"Web服务器01\"\r\n}', '{\"success\":false,\"message\":\"数据库操作失败\",\"error\":\"数据库操作失败\"}', 226, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (45, '2026-02-05 13:10:12.434', '2026-02-05 13:10:12.434', '创建主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 500, 'POST', '/api/v1/hosts', '{\r\n    \"hostname\": \"database\",\r\n    \"ip_address\": \"10.100.1.100\",\r\n    \"port\": 22,\r\n    \"username\": \"root\",\r\n    \"password\": \"encrypted_password\",\r\n    \"os_type\": \"linux\",\r\n    \"cpu_cores\": 8,\r\n    \"memory_gb\": 16,\r\n    \"disk_space_gb\": 500,\r\n    \"group_id\": 1,\r\n    \"description\": \"Web服务器01\"\r\n}', '{\"success\":false,\"message\":\"数据库操作失败\",\"error\":\"数据库操作失败\"}', 223, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (46, '2026-02-05 13:14:52.725', '2026-02-05 13:14:52.725', '创建主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 500, 'POST', '/api/v1/hosts', '{\r\n    \"hostname\": \"database\",\r\n    \"ip_address\": \"10.100.1.100\",\r\n    \"port\": 22,\r\n    \"username\": \"root\",\r\n    \"password\": \"encrypted_password\",\r\n    \"os_type\": \"linux\",\r\n    \"cpu_cores\": 8,\r\n    \"memory_gb\": 16,\r\n    \"disk_space_gb\": 500,\r\n    \"group_id\": 1,\r\n    \"description\": \"Web服务器01\"\r\n}', '{\"success\":false,\"message\":\"数据库操作失败\",\"error\":\"数据库操作失败\"}', 228, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (47, '2026-02-05 13:16:43.774', '2026-02-05 13:16:43.774', '创建主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 200, 'POST', '/api/v1/hosts', '{\r\n    \"hostname\": \"database\",\r\n    \"ip_address\": \"10.100.1.100\",\r\n    \"port\": 22,\r\n    \"username\": \"root\",\r\n    \"password\": \"encrypted_password\",\r\n    \"os_type\": \"linux\",\r\n    \"cpu_cores\": 8,\r\n    \"memory_gb\": 16,\r\n    \"disk_space_gb\": 500,\r\n    \"group_id\": 1,\r\n    \"description\": \"Web服务器01\"\r\n}', '{\"success\":true,\"message\":\"主机创建成功\",\"data\":{\"id\":11,\"hostname\":\"database\",\"ip_address\":\"10.100.1.100\",\"port\":22,\"username\":\"root\",\"password\":\"$2a$10$SKqOhxfRQ0eC8WBZXzljE.MKknxoPxWrlU86koSO3hASOvYbBG.DO\",\"os_type\":\"linux\",\"cpu_cores\":8,\"memory_gb\":16,\"disk_space_gb\":500,\"group_id\":1,\"status\":1,\"monitoring_enabled\":1,\"last_heartbeat\":null,\"description\":\"Web服务器01\",\"created_by\":10,\"updated_by\":null,\"created_at\":\"2026-02-05T13:16:43.724+08:00\",\"updated_at\":\"2026-02-05T13:16:43.724+08:00\",\"deleted_at\":null,\"group\":null}}', 266, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (48, '2026-02-05 13:22:00.798', '2026-02-05 13:22:00.798', '更新主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 200, 'PUT', '/api/v1/hosts/:id', '{\r\n    \"hostname\": \"database\",\r\n    \"ip_address\": \"10.100.1.100\",\r\n    \"port\": 1202,\r\n    \"username\": \"root\",\r\n    \"password\": \"new_encrypted_password\",\r\n    \"os_type\": \"linux\",\r\n    \"cpu_cores\": 16,\r\n    \"memory_gb\": 32,\r\n    \"disk_space_gb\": 1000,\r\n    \"group_id\": 1,\r\n    \"description\": \"升级后的Web服务器01\"\r\n}', '{\"success\":true,\"message\":\"主机更新成功\",\"data\":{\"id\":11,\"hostname\":\"database\",\"ip_address\":\"10.100.1.100\",\"port\":1202,\"username\":\"root\",\"password\":\"$2a$10$u6E6v5Jcp5T9jVuyiVRMy.uFXtfU83QxCLl9/zRRtYhRqmHWkM8L6\",\"os_type\":\"linux\",\"cpu_cores\":16,\"memory_gb\":32,\"disk_space_gb\":1000,\"group_id\":1,\"status\":1,\"monitoring_enabled\":1,\"last_heartbeat\":null,\"description\":\"升级后的Web服务器01\",\"created_by\":10,\"updated_by\":10,\"created_at\":\"2026-02-05T13:16:44+08:00\",\"updated_at\":\"2026-02-05T13:22:00.7+08:00\",\"deleted_at\":null,\"group\":{\"id\":1,\"name\":\"Web服务器组\",\"description\":\"用于存放Web应用服务器的主机组\",\"status\":1,\"created_by\":1,\"updated_by\":1,\"created_at\":\"2026-02-05T04:11:27+08:00\",\"updated_at\":\"2026-02-05T04:11:27+08:00\",\"deleted_at\":null,\"hosts\":null}}}', 309, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (49, '2026-02-05 13:24:35.013', '2026-02-05 13:24:35.013', '删除主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 200, 'DELETE', '/api/v1/hosts/:id', '', '{\"success\":true,\"message\":\"主机删除成功\"}', 151, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (50, '2026-02-05 13:30:38.168', '2026-02-05 13:30:38.168', '批量删除主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 200, 'DELETE', '/api/v1/hosts/batch', '{\r\n    \"ids\": [8, 9, 10]\r\n}', '{\"success\":true,\"message\":\"批量删除主机成功\",\"data\":{\"deleted_count\":3}}', 52, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (51, '2026-02-05 13:32:00.027', '2026-02-05 13:32:00.027', '更新主机状态', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'PUT', '/api/v1/hosts/:id/status', '{\r\n    \"status\": 0\r\n}', '{\"success\":false,\"message\":\"参数验证失败\",\"error\":\"请求参数: Key: \'HostStatusUpdateRequest.Status\' Error:Field validation for \'Status\' failed on the \'required\' tag\"}', 0, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (52, '2026-02-05 13:32:16.109', '2026-02-05 13:32:16.109', '更新主机状态', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'PUT', '/api/v1/hosts/:id/status', '{\r\n    \"Status\": 0\r\n}', '{\"success\":false,\"message\":\"参数验证失败\",\"error\":\"请求参数: Key: \'HostStatusUpdateRequest.Status\' Error:Field validation for \'Status\' failed on the \'required\' tag\"}', 0, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (53, '2026-02-05 13:32:39.771', '2026-02-05 13:32:39.771', '更新主机状态', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'PUT', '/api/v1/hosts/:id/status', '{\r\n    \"status\": 0\r\n}', '{\"success\":false,\"message\":\"参数验证失败\",\"error\":\"请求参数: Key: \'HostStatusUpdateRequest.Status\' Error:Field validation for \'Status\' failed on the \'required\' tag\"}', 0, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (54, '2026-02-05 13:34:10.697', '2026-02-05 13:34:10.697', '更新主机状态', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'PUT', '/api/v1/hosts/:id/status', '{\r\n    \"status\": \"0\"\r\n}', '{\"success\":false,\"message\":\"参数验证失败\",\"error\":\"请求参数: json: cannot unmarshal string into Go struct field HostStatusUpdateRequest.status of type int8\"}', 0, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (55, '2026-02-05 13:34:17.079', '2026-02-05 13:34:17.079', '更新主机状态', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'PUT', '/api/v1/hosts/:id/status', '{\r\n    \"status\": 0\r\n}', '{\"success\":false,\"message\":\"参数验证失败\",\"error\":\"请求参数: Key: \'HostStatusUpdateRequest.Status\' Error:Field validation for \'Status\' failed on the \'required\' tag\"}', 0, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (56, '2026-02-05 13:40:05.250', '2026-02-05 13:40:05.250', '更新主机状态', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 200, 'PUT', '/api/v1/hosts/:id/status', '{\r\n    \"status\": -1\r\n}', '{\"success\":true,\"message\":\"主机状态更新成功\",\"data\":{\"id\":1,\"status\":-1}}', 143, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (57, '2026-02-05 13:42:19.743', '2026-02-05 13:42:19.743', '更新主机状态', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'PUT', '/api/v1/hosts/:id/status', '{\r\n    \"status\": 2\r\n}', '{\"success\":false,\"message\":\"参数验证失败\",\"error\":\"请求参数: Key: \'HostStatusUpdateRequest.Status\' Error:Field validation for \'Status\' failed on the \'oneof\' tag\"}', 0, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (58, '2026-02-05 13:42:24.271', '2026-02-05 13:42:24.271', '更新主机状态', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'PUT', '/api/v1/hosts/:id/status', '{\r\n    \"status\": 2\r\n}', '{\"success\":false,\"message\":\"参数验证失败\",\"error\":\"请求参数: Key: \'HostStatusUpdateRequest.Status\' Error:Field validation for \'Status\' failed on the \'oneof\' tag\"}', 0, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (59, '2026-02-05 13:42:38.148', '2026-02-05 13:42:38.148', '更新主机状态', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 200, 'PUT', '/api/v1/hosts/:id/status', '{\r\n    \"status\": 2\r\n}', '{\"success\":true,\"message\":\"主机状态更新成功\",\"data\":{\"id\":1,\"status\":2}}', 146, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (60, '2026-02-05 13:45:06.638', '2026-02-05 13:45:06.638', '更新主机监控状态', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'PUT', '/api/v1/hosts/:id/monitoring', '{\r\n    \"monitoring_enabled\": 0\r\n}', '{\"success\":false,\"message\":\"参数验证失败\",\"error\":\"请求参数: Key: \'HostMonitoringUpdateRequest.MonitoringEnabled\' Error:Field validation for \'MonitoringEnabled\' failed on the \'required\' tag\"}', 0, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (61, '2026-02-05 13:50:22.923', '2026-02-05 13:50:22.923', '创建主机组', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'POST', '/api/v1/host-groups', '{\r\n    \"name\": \"数据库服务器组\",\r\n    \"description\": \"用于存放数据库服务器的主机组\"\r\n}', '{\"success\":false,\"message\":\"数据验证失败\",\"error\":\"数据验证失败\"}', 3, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (62, '2026-02-05 13:51:54.025', '2026-02-05 13:51:54.025', '创建主机组', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 200, 'POST', '/api/v1/host-groups', '{\r\n    \"name\": \"AAA服务器组\",\r\n    \"description\": \"用于存放数据库服务器的主机组\"\r\n}', '{\"success\":true,\"message\":\"主机组创建成功\",\"data\":{\"id\":5,\"name\":\"AAA服务器组\",\"description\":\"用于存放数据库服务器的主机组\",\"status\":1,\"created_by\":10,\"updated_by\":null,\"created_at\":\"2026-02-05T13:51:53.972+08:00\",\"updated_at\":\"2026-02-05T13:51:53.972+08:00\",\"deleted_at\":null,\"hosts\":null}}', 97, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (63, '2026-02-05 13:54:31.593', '2026-02-05 13:54:31.593', '更新主机组', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'PUT', '/api/v1/host-groups/:id', '{\r\n    \"name\": \"BBB应用服务器组\",\r\n    \"description\": \"Web应用服务器主机组\"\r\n}', '{\"success\":false,\"message\":\"数据验证失败\",\"error\":\"数据验证失败\"}', 91, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (64, '2026-02-05 13:55:03.508', '2026-02-05 13:55:03.508', '更新主机组', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'PUT', '/api/v1/host-groups/:id', '{\r\n    \"name\": \"BBB应用服务器组\",\r\n    \"description\": \"AAA应用服务器主机组\"\r\n}', '{\"success\":false,\"message\":\"数据验证失败\",\"error\":\"数据验证失败\"}', 96, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (65, '2026-02-05 14:00:18.116', '2026-02-05 14:00:18.116', '更新主机组', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 500, 'PUT', '/api/v1/host-groups/:id', '{\r\n    \"name\": \"BBB应用服务器组\",\r\n    \"description\": \"AAA应用服务器主机组\"\r\n}', '{\"success\":false,\"message\":\"数据库操作失败\",\"error\":\"数据库操作失败\"}', 94, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (66, '2026-02-05 14:00:39.683', '2026-02-05 14:00:39.683', '更新主机组', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 500, 'PUT', '/api/v1/host-groups/:id', '{\r\n    \"name\": \"BBB应用服务器组\",\r\n    \"description\": \"AAA应用服务器主机组\"\r\n}', '{\"success\":false,\"message\":\"服务器内部错误\",\"error\":\"record not found\"}', 97, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (67, '2026-02-05 14:00:47.146', '2026-02-05 14:00:47.146', '更新主机组', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 500, 'PUT', '/api/v1/host-groups/:id', '{\r\n    \"name\": \"BBB应用服务器组\",\r\n    \"description\": \"AAA应用服务器主机组\"\r\n}', '{\"success\":false,\"message\":\"服务器内部错误\",\"error\":\"record not found\"}', 95, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (68, '2026-02-05 14:01:55.477', '2026-02-05 14:01:55.477', '更新主机组', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 500, 'PUT', '/api/v1/host-groups/:id', '{\r\n    \"name\": \"BBB应用服务器组\",\r\n    \"description\": \"AAA应用服务器主机组\"\r\n}', '{\"success\":false,\"message\":\"数据库操作失败\",\"error\":\"数据库操作失败\"}', 90, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (69, '2026-02-05 14:02:15.300', '2026-02-05 14:02:15.300', '更新主机组', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 500, 'PUT', '/api/v1/host-groups/:id', '{\r\n    \"name\": \"BBB应用服务器组\",\r\n    \"description\": \"AAA应用服务器主机组\"\r\n}', '{\"success\":false,\"message\":\"数据库操作失败\",\"error\":\"数据库操作失败\"}', 102, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (70, '2026-02-05 14:03:23.928', '2026-02-05 14:03:23.928', '更新主机组', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 500, 'PUT', '/api/v1/host-groups/:id', '{\r\n    \"name\": \"BBB应用服务器组\",\r\n    \"description\": \"AAA应用服务器主机组\"\r\n}', '{\"success\":false,\"message\":\"数据库操作失败\",\"error\":\"数据库操作失败\"}', 99, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (71, '2026-02-05 14:04:04.951', '2026-02-05 14:04:04.951', '更新主机组', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'PUT', '/api/v1/host-groups/:id', '{\r\n    \"name\": \"BBB应用服务器组\",\r\n    \"description\": \"AAA应用服务器主机组\"\r\n}', '{\"success\":false,\"message\":\"数据验证失败\",\"error\":\"数据验证失败\"}', 100, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (72, '2026-02-05 14:04:53.347', '2026-02-05 14:04:53.347', '更新主机组', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'PUT', '/api/v1/host-groups/:id', '{\r\n    \"name\": \"BBB应用服务器组\",\r\n    \"description\": \"AAA应用服务器主机组\"\r\n}', '{\"success\":false,\"message\":\"数据验证失败\",\"error\":\"数据验证失败\"}', 99, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (73, '2026-02-05 14:06:49.408', '2026-02-05 14:06:49.408', '更新主机组', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 200, 'PUT', '/api/v1/host-groups/:id', '{\r\n    \"name\": \"BBB应用服务器组\",\r\n    \"description\": \"AAA应用服务器主机组\"\r\n}', '{\"success\":true,\"message\":\"主机组更新成功\",\"data\":{\"id\":5,\"name\":\"BBB应用服务器组\",\"description\":\"AAA应用服务器主机组\",\"status\":1,\"created_by\":10,\"updated_by\":10,\"created_at\":\"2026-02-05T13:51:54+08:00\",\"updated_at\":\"2026-02-05T14:06:49.354+08:00\",\"deleted_at\":null,\"hosts\":[]}}', 185, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (74, '2026-02-06 11:20:09.357', '2026-02-06 11:20:09.357', '更新菜单', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'PUT', '/api/v1/menus/:id', '{\"id\":3,\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"name\":\"asset\",\"title\":\"主机管理\",\"path\":\"/asset\",\"component\":\"\",\"parent_id\":null,\"parent\":null,\"children\":[],\"icon\":\"Setting\",\"sort\":1,\"is_hidden\":false,\"status\":1,\"level\":0}', '{\"message\":\"更新成功\"}', 52, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (75, '2026-02-06 11:20:21.526', '2026-02-06 11:20:21.526', '为角色分配菜单', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/menus', '{\"menu_ids\":[3]}', '{\"message\":\"菜单权限分配成功\"}', 87, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (76, '2026-02-06 11:23:47.338', '2026-02-06 11:23:47.338', '为角色分配菜单', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/menus', '{\"menu_ids\":[118]}', '{\"message\":\"菜单权限分配成功\"}', 91, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (77, '2026-02-06 11:25:31.160', '2026-02-06 11:25:31.160', '为角色分配菜单', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/menus', '{\"menu_ids\":[119]}', '{\"message\":\"菜单权限分配成功\"}', 85, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (78, '2026-02-06 11:25:46.875', '2026-02-06 11:25:46.875', '更新主机监控状态', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'PUT', '/api/v1/hosts/:id/monitoring', '{\"monitoring_enabled\":2}', '{\"success\":true,\"message\":\"监控状态更新成功\",\"data\":{\"id\":1,\"monitoring_enabled\":2}}', 140, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (79, '2026-02-06 11:25:54.211', '2026-02-06 11:25:54.211', '更新主机监控状态', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'PUT', '/api/v1/hosts/:id/monitoring', '{\"monitoring_enabled\":1}', '{\"success\":true,\"message\":\"监控状态更新成功\",\"data\":{\"id\":1,\"monitoring_enabled\":1}}', 139, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (80, '2026-02-06 11:26:37.125', '2026-02-06 11:26:37.125', '创建主机', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/hosts', '{\"id\":0,\"hostname\":\"测试\",\"ip_address\":\"19.168.0.1\",\"port\":22,\"username\":\"root\",\"password\":\"Zhy20250730!\",\"os_type\":\"linux\",\"cpu_cores\":1,\"memory_gb\":1,\"disk_space_gb\":1,\"group_id\":5,\"description\":\"\"}', '{\"success\":true,\"message\":\"主机创建成功\",\"data\":{\"id\":12,\"hostname\":\"测试\",\"ip_address\":\"19.168.0.1\",\"port\":22,\"username\":\"root\",\"password\":\"$2a$10$TQIO.3roAbEldwr6ccmEVOifGIXvdsZsP3gC/Yav8Kv/dnxAG8BLy\",\"os_type\":\"linux\",\"cpu_cores\":1,\"memory_gb\":1,\"disk_space_gb\":1,\"group_id\":5,\"status\":1,\"monitoring_enabled\":1,\"last_heartbeat\":null,\"description\":\"\",\"created_by\":10,\"updated_by\":null,\"created_at\":\"2026-02-06T11:26:37.074+08:00\",\"updated_at\":\"2026-02-06T11:26:37.074+08:00\",\"deleted_at\":null,\"group\":null}}', 257, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (81, '2026-02-06 11:28:28.370', '2026-02-06 11:28:28.370', '删除主机', 10, 'admin', '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'DELETE', '/api/v1/hosts/:id', '', '{\"success\":true,\"message\":\"主机删除成功\"}', 147, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (82, '2026-02-10 19:26:14.636', '2026-02-10 19:26:14.636', '更新菜单', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'PUT', '/api/v1/menus/:id', '{\"id\":120,\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"name\":\"credential\",\"title\":\"凭据管理\",\"path\":\"/asset/credential\",\"component\":\"CredentialManageView\",\"parent_id\":3,\"parent\":null,\"children\":[],\"icon\":\"User\",\"sort\":3,\"is_hidden\":false,\"status\":1,\"level\":1}', '{\"message\":\"更新成功\"}', 5, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (83, '2026-02-10 19:28:20.980', '2026-02-10 19:28:20.980', '为角色分配菜单', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/menus', '{\"menu_ids\":[120]}', '{\"message\":\"菜单权限分配成功\"}', 5, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (84, '2026-02-10 19:35:20.216', '2026-02-10 19:35:20.216', '添加Casbin策略', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/operation-logs/:id\",\"method\":\"DELETE\"}', '{\"message\":\"策略添加成功\"}', 6, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (85, '2026-02-10 19:35:20.485', '2026-02-10 19:35:20.485', '添加Casbin策略', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/users/profile\",\"method\":\"GET\"}', '{\"message\":\"策略添加成功\"}', 7, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (86, '2026-02-10 19:35:20.531', '2026-02-10 19:35:20.531', '添加Casbin策略', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/credentials\",\"method\":\"POST\"}', '{\"message\":\"策略添加成功\"}', 7, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (87, '2026-02-10 19:35:20.811', '2026-02-10 19:35:20.811', '添加Casbin策略', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/credentials\",\"method\":\"GET\"}', '{\"message\":\"策略添加成功\"}', 6, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (88, '2026-02-10 19:35:20.857', '2026-02-10 19:35:20.857', '添加Casbin策略', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/credentials/:id\",\"method\":\"GET\"}', '{\"message\":\"策略添加成功\"}', 6, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (89, '2026-02-10 19:35:21.124', '2026-02-10 19:35:21.124', '添加Casbin策略', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/credentials/:id\",\"method\":\"PUT\"}', '{\"message\":\"策略添加成功\"}', 5, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (90, '2026-02-10 19:35:21.170', '2026-02-10 19:35:21.170', '添加Casbin策略', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/credentials/:id\",\"method\":\"DELETE\"}', '{\"message\":\"策略添加成功\"}', 5, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (91, '2026-02-10 19:35:21.438', '2026-02-10 19:35:21.438', '添加Casbin策略', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/credentials/batch\",\"method\":\"DELETE\"}', '{\"message\":\"策略添加成功\"}', 7, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (92, '2026-02-10 19:35:21.484', '2026-02-10 19:35:21.484', '添加Casbin策略', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/credentials/host\",\"method\":\"GET\"}', '{\"message\":\"策略添加成功\"}', 5, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (93, '2026-02-10 19:35:21.764', '2026-02-10 19:35:21.764', '移除Casbin策略', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'DELETE', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/permissions/all/:id\",\"method\":\"GET\"}', '{\"message\":\"策略移除成功\"}', 7, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (94, '2026-02-10 19:35:21.809', '2026-02-10 19:35:21.809', '移除Casbin策略', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'DELETE', '/api/v1/roles/:id/policies', '{\"path\":\"/api/v1/operation-logs/stats\",\"method\":\"GET\"}', '{\"message\":\"策略移除成功\"}', 6, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (95, '2026-02-10 20:14:37.550', '2026-02-10 20:14:37.550', '创建主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'POST', '/api/v1/hosts', '{\r\n    \"hostname\": \"web-server-01\",\r\n    \"ip_address\": \"192.168.1.100\",\r\n    \"port\": 22,\r\n    \"os_type\": \"linux\",\r\n    \"cpu_cores\": 8,\r\n    \"memory_gb\": 16,\r\n    \"disk_space_gb\": 500,\r\n    \"group_id\": 1,\r\n    \"description\": \"Web服务器01\",\r\n    \"credential_ids\": [1, 2]\r\n}', '{\"success\":false,\"message\":\"数据验证失败\",\"error\":\"数据验证失败\"}', 4, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (96, '2026-02-10 20:14:47.854', '2026-02-10 20:14:47.854', '创建主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 500, 'POST', '/api/v1/hosts', '{\r\n    \"hostname\": \"web-server-011\",\r\n    \"ip_address\": \"192.168.1.120\",\r\n    \"port\": 22,\r\n    \"os_type\": \"linux\",\r\n    \"cpu_cores\": 8,\r\n    \"memory_gb\": 16,\r\n    \"disk_space_gb\": 500,\r\n    \"group_id\": 1,\r\n    \"description\": \"Web服务器01\",\r\n    \"credential_ids\": [1, 2]\r\n}', '{\"success\":false,\"message\":\"数据库操作失败\",\"error\":\"数据库操作失败\"}', 9, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (97, '2026-02-10 20:15:35.445', '2026-02-10 20:15:35.445', '创建主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 500, 'POST', '/api/v1/hosts', '{\r\n    \"hostname\": \"web-server-011\",\r\n    \"ip_address\": \"192.168.1.120\",\r\n    \"username\": \"abc\",\r\n    \"port\": 22,\r\n    \"os_type\": \"linux\",\r\n    \"cpu_cores\": 8,\r\n    \"memory_gb\": 16,\r\n    \"disk_space_gb\": 500,\r\n    \"group_id\": 1,\r\n    \"description\": \"Web服务器01\",\r\n    \"credential_ids\": [1, 2]\r\n}', '{\"success\":false,\"message\":\"数据库操作失败\",\"error\":\"数据库操作失败\"}', 37, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (98, '2026-02-10 20:16:50.608', '2026-02-10 20:16:50.608', '创建主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 500, 'POST', '/api/v1/hosts', '{\r\n    \"hostname\": \"web-server-011\",\r\n    \"ip_address\": \"192.168.1.120\",\r\n    \"port\": 22,\r\n    \"os_type\": \"linux\",\r\n    \"cpu_cores\": 8,\r\n    \"memory_gb\": 16,\r\n    \"disk_space_gb\": 500,\r\n    \"group_id\": 1,\r\n    \"description\": \"Web服务器01\",\r\n    \"credential_ids\": [1, 2]\r\n}', '{\"success\":false,\"message\":\"数据库操作失败\",\"error\":\"数据库操作失败\"}', 10, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (99, '2026-02-10 20:18:33.749', '2026-02-10 20:18:33.749', '创建主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 200, 'POST', '/api/v1/hosts', '{\r\n    \"hostname\": \"web-server-011\",\r\n    \"ip_address\": \"192.168.1.120\",\r\n    \"port\": 22,\r\n    \"os_type\": \"linux\",\r\n    \"cpu_cores\": 8,\r\n    \"memory_gb\": 16,\r\n    \"disk_space_gb\": 500,\r\n    \"group_id\": 1,\r\n    \"description\": \"Web服务器01\",\r\n    \"credential_ids\": [1, 2]\r\n}', '{\"success\":true,\"message\":\"主机创建成功\",\"data\":{\"id\":13,\"hostname\":\"web-server-011\",\"ip_address\":\"192.168.1.120\",\"port\":22,\"os_type\":\"linux\",\"cpu_cores\":8,\"memory_gb\":16,\"disk_space_gb\":500,\"group_id\":1,\"status\":1,\"monitoring_enabled\":1,\"last_heartbeat\":null,\"description\":\"Web服务器01\",\"created_by\":10,\"updated_by\":10,\"created_at\":\"2026-02-10T20:18:33.732+08:00\",\"updated_at\":\"2026-02-10T20:18:33.739+08:00\",\"deleted_at\":null,\"group\":null,\"credentials\":[{\"id\":1,\"name\":\"root-default\",\"username\":\"root\",\"password\":\"$2a$10$example_hashed_password\",\"description\":\"默认root用户凭据\",\"created_at\":\"2026-02-06T05:54:01+08:00\",\"updated_at\":\"2026-02-06T05:54:01+08:00\",\"deleted_at\":null,\"hosts\":null},{\"id\":2,\"name\":\"admin-default\",\"username\":\"admin\",\"password\":\"$2a$10$example_hashed_password\",\"description\":\"默认管理员凭据\",\"created_at\":\"2026-02-06T05:54:01+08:00\",\"updated_at\":\"2026-02-06T05:54:01+08:00\",\"deleted_at\":null,\"hosts\":null}]}}', 23, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (100, '2026-02-10 20:19:09.264', '2026-02-10 20:19:09.264', '创建主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'POST', '/api/v1/hosts', '{\r\n    \"hostname\": \"web-server-011\",\r\n    \"ip_address\": \"192.168.1.120\",\r\n    \"port\": 22,\r\n    \"os_type\": \"linux\",\r\n    \"cpu_cores\": 8,\r\n    \"memory_gb\": 16,\r\n    \"disk_space_gb\": 500,\r\n    \"group_id\": 1,\r\n    \"description\": \"Web服务器01\",\r\n    \"credential_ids\": [1, 2]\r\n}', '{\"success\":false,\"message\":\"数据验证失败\",\"error\":\"数据验证失败\"}', 2, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (101, '2026-02-10 20:19:15.219', '2026-02-10 20:19:15.219', '创建主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 400, 'POST', '/api/v1/hosts', '{\r\n    \"hostname\": \"web-server-01111\",\r\n    \"ip_address\": \"192.168.1.120\",\r\n    \"port\": 22,\r\n    \"os_type\": \"linux\",\r\n    \"cpu_cores\": 8,\r\n    \"memory_gb\": 16,\r\n    \"disk_space_gb\": 500,\r\n    \"group_id\": 1,\r\n    \"description\": \"Web服务器01\",\r\n    \"credential_ids\": [1, 2]\r\n}', '{\"success\":false,\"message\":\"数据验证失败\",\"error\":\"数据验证失败\"}', 4, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (102, '2026-02-10 20:19:20.831', '2026-02-10 20:19:20.831', '创建主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 200, 'POST', '/api/v1/hosts', '{\r\n    \"hostname\": \"web-server-01111\",\r\n    \"ip_address\": \"192.168.1.10\",\r\n    \"port\": 22,\r\n    \"os_type\": \"linux\",\r\n    \"cpu_cores\": 8,\r\n    \"memory_gb\": 16,\r\n    \"disk_space_gb\": 500,\r\n    \"group_id\": 1,\r\n    \"description\": \"Web服务器01\",\r\n    \"credential_ids\": [1, 2]\r\n}', '{\"success\":true,\"message\":\"主机创建成功\",\"data\":{\"id\":14,\"hostname\":\"web-server-01111\",\"ip_address\":\"192.168.1.10\",\"port\":22,\"os_type\":\"linux\",\"cpu_cores\":8,\"memory_gb\":16,\"disk_space_gb\":500,\"group_id\":1,\"status\":1,\"monitoring_enabled\":1,\"last_heartbeat\":null,\"description\":\"Web服务器01\",\"created_by\":10,\"updated_by\":10,\"created_at\":\"2026-02-10T20:19:20.813+08:00\",\"updated_at\":\"2026-02-10T20:19:20.821+08:00\",\"deleted_at\":null,\"group\":null,\"credentials\":[{\"id\":1,\"name\":\"root-default\",\"username\":\"root\",\"password\":\"$2a$10$example_hashed_password\",\"description\":\"默认root用户凭据\",\"created_at\":\"2026-02-06T05:54:01+08:00\",\"updated_at\":\"2026-02-06T05:54:01+08:00\",\"deleted_at\":null,\"hosts\":null},{\"id\":2,\"name\":\"admin-default\",\"username\":\"admin\",\"password\":\"$2a$10$example_hashed_password\",\"description\":\"默认管理员凭据\",\"created_at\":\"2026-02-06T05:54:01+08:00\",\"updated_at\":\"2026-02-06T05:54:01+08:00\",\"deleted_at\":null,\"hosts\":null}]}}', 27, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (103, '2026-02-10 20:58:24.346', '2026-02-10 20:58:24.346', '批量删除主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 200, 'DELETE', '/api/v1/hosts/batch', '{\r\n    \"ids\": [8, 9, 10]\r\n}', '{\"success\":true,\"message\":\"批量删除主机成功\",\"data\":{\"deleted_count\":0}}', 4, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (104, '2026-02-10 20:59:02.939', '2026-02-10 20:59:02.939', '批量删除主机', 10, 'admin', '::1', 'Apifox/1.0.0 (https://apifox.com)', 200, 'DELETE', '/api/v1/hosts/batch', '{\r\n    \"ids\": [13,14]\r\n}', '{\"success\":true,\"message\":\"批量删除主机成功\",\"data\":{\"deleted_count\":2}}', 5, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (105, '2026-02-10 22:32:14.824', '2026-02-10 22:32:14.824', '更新主机', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'PUT', '/api/v1/hosts/:id', '{\"id\":2,\"hostname\":\"\",\"ip_address\":\"\",\"port\":22,\"os_type\":\"linux\",\"description\":\"\",\"credential_ids\":[1]}', '{\"success\":true,\"message\":\"主机更新成功\",\"data\":{\"id\":2,\"hostname\":\"web-server-02\",\"ip_address\":\"192.168.1.102\",\"port\":22,\"os_type\":\"linux\",\"cpu_cores\":8,\"memory_gb\":16,\"disk_space_gb\":500,\"group_id\":1,\"status\":1,\"monitoring_enabled\":1,\"last_heartbeat\":\"2026-02-05T04:08:27+08:00\",\"description\":\"生产环境Web服务器02\",\"created_by\":1,\"updated_by\":10,\"created_at\":\"2026-02-05T04:11:27+08:00\",\"updated_at\":\"2026-02-10T22:32:14.817+08:00\",\"deleted_at\":null,\"group\":{\"id\":1,\"name\":\"Web服务器组\",\"description\":\"用于存放Web应用服务器的主机组\",\"status\":1,\"created_by\":1,\"updated_by\":1,\"created_at\":\"2026-02-05T04:11:27+08:00\",\"updated_at\":\"2026-02-05T04:11:27+08:00\",\"deleted_at\":null,\"hosts\":null},\"credentials\":[{\"id\":1,\"name\":\"root-default\",\"username\":\"root\",\"password\":\"$2a$10$example_hashed_password\",\"description\":\"默认root用户凭据\",\"created_at\":\"2026-02-06T05:54:01+08:00\",\"updated_at\":\"2026-02-06T05:54:01+08:00\",\"deleted_at\":null,\"hosts\":null}]}}', 36, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (106, '2026-02-10 22:32:31.685', '2026-02-10 22:32:31.685', '更新主机', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'PUT', '/api/v1/hosts/:id', '{\"id\":2,\"hostname\":\"\",\"ip_address\":\"\",\"port\":22,\"os_type\":\"linux\",\"description\":\"\",\"credential_ids\":[1,2]}', '{\"success\":true,\"message\":\"主机更新成功\",\"data\":{\"id\":2,\"hostname\":\"web-server-02\",\"ip_address\":\"192.168.1.102\",\"port\":22,\"os_type\":\"linux\",\"cpu_cores\":8,\"memory_gb\":16,\"disk_space_gb\":500,\"group_id\":1,\"status\":1,\"monitoring_enabled\":1,\"last_heartbeat\":\"2026-02-05T04:08:27+08:00\",\"description\":\"生产环境Web服务器02\",\"created_by\":1,\"updated_by\":10,\"created_at\":\"2026-02-05T04:11:27+08:00\",\"updated_at\":\"2026-02-10T22:32:31.676+08:00\",\"deleted_at\":null,\"group\":{\"id\":1,\"name\":\"Web服务器组\",\"description\":\"用于存放Web应用服务器的主机组\",\"status\":1,\"created_by\":1,\"updated_by\":1,\"created_at\":\"2026-02-05T04:11:27+08:00\",\"updated_at\":\"2026-02-05T04:11:27+08:00\",\"deleted_at\":null,\"hosts\":null},\"credentials\":[{\"id\":1,\"name\":\"root-default\",\"username\":\"root\",\"password\":\"$2a$10$example_hashed_password\",\"description\":\"默认root用户凭据\",\"created_at\":\"2026-02-06T05:54:01+08:00\",\"updated_at\":\"2026-02-06T05:54:01+08:00\",\"deleted_at\":null,\"hosts\":null},{\"id\":2,\"name\":\"admin-default\",\"username\":\"admin\",\"password\":\"$2a$10$example_hashed_password\",\"description\":\"默认管理员凭据\",\"created_at\":\"2026-02-06T05:54:01+08:00\",\"updated_at\":\"2026-02-06T05:54:01+08:00\",\"deleted_at\":null,\"hosts\":null}]}}', 28, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (107, '2026-02-10 22:38:20.598', '2026-02-10 22:38:20.598', '更新主机', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'PUT', '/api/v1/hosts/:id', '{\"id\":1,\"hostname\":\"\",\"ip_address\":\"\",\"port\":22,\"os_type\":\"linux\",\"description\":\"\",\"credential_ids\":[1]}', '{\"success\":true,\"message\":\"主机更新成功\",\"data\":{\"id\":1,\"hostname\":\"web-server-01\",\"ip_address\":\"192.168.1.101\",\"port\":22,\"os_type\":\"linux\",\"cpu_cores\":8,\"memory_gb\":16,\"disk_space_gb\":500,\"group_id\":1,\"status\":2,\"monitoring_enabled\":1,\"last_heartbeat\":\"2026-02-05T04:06:27+08:00\",\"description\":\"生产环境Web服务器01\",\"created_by\":1,\"updated_by\":10,\"created_at\":\"2026-02-05T04:11:27+08:00\",\"updated_at\":\"2026-02-10T22:38:20.59+08:00\",\"deleted_at\":null,\"group\":{\"id\":1,\"name\":\"Web服务器组\",\"description\":\"用于存放Web应用服务器的主机组\",\"status\":1,\"created_by\":1,\"updated_by\":1,\"created_at\":\"2026-02-05T04:11:27+08:00\",\"updated_at\":\"2026-02-05T04:11:27+08:00\",\"deleted_at\":null,\"hosts\":null},\"credentials\":[{\"id\":1,\"name\":\"root-default\",\"username\":\"root\",\"password\":\"$2a$10$example_hashed_password\",\"description\":\"默认root用户凭据\",\"created_at\":\"2026-02-06T05:54:01+08:00\",\"updated_at\":\"2026-02-06T05:54:01+08:00\",\"deleted_at\":null,\"hosts\":null}]}}', 34, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (108, '2026-02-10 22:40:03.875', '2026-02-10 22:40:03.875', '更新主机', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'PUT', '/api/v1/hosts/:id', '{\"id\":1,\"hostname\":\"\",\"ip_address\":\"\",\"port\":22,\"os_type\":\"linux\",\"description\":\"\",\"credential_ids\":[1,2,3,4]}', '{\"success\":true,\"message\":\"主机更新成功\",\"data\":{\"id\":1,\"hostname\":\"web-server-01\",\"ip_address\":\"192.168.1.101\",\"port\":22,\"os_type\":\"linux\",\"cpu_cores\":8,\"memory_gb\":16,\"disk_space_gb\":500,\"group_id\":1,\"status\":2,\"monitoring_enabled\":1,\"last_heartbeat\":\"2026-02-05T04:06:27+08:00\",\"description\":\"生产环境Web服务器01\",\"created_by\":1,\"updated_by\":10,\"created_at\":\"2026-02-05T04:11:27+08:00\",\"updated_at\":\"2026-02-10T22:40:03.865+08:00\",\"deleted_at\":null,\"group\":{\"id\":1,\"name\":\"Web服务器组\",\"description\":\"用于存放Web应用服务器的主机组\",\"status\":1,\"created_by\":1,\"updated_by\":1,\"created_at\":\"2026-02-05T04:11:27+08:00\",\"updated_at\":\"2026-02-05T04:11:27+08:00\",\"deleted_at\":null,\"hosts\":null},\"credentials\":[{\"id\":1,\"name\":\"root-default\",\"username\":\"root\",\"password\":\"$2a$10$example_hashed_password\",\"description\":\"默认root用户凭据\",\"created_at\":\"2026-02-06T05:54:01+08:00\",\"updated_at\":\"2026-02-06T05:54:01+08:00\",\"deleted_at\":null,\"hosts\":null},{\"id\":2,\"name\":\"admin-default\",\"username\":\"admin\",\"password\":\"$2a$10$example_hashed_password\",\"description\":\"默认管理员凭据\",\"created_at\":\"2026-02-06T05:54:01+08:00\",\"updated_at\":\"2026-02-06T05:54:01+08:00\",\"deleted_at\":null,\"hosts\":null},{\"id\":3,\"name\":\"ubuntu-default\",\"username\":\"ubuntu\",\"password\":\"$2a$10$example_hashed_password\",\"description\":\"Ubuntu系统默认用户凭据\",\"created_at\":\"2026-02-06T05:54:01+08:00\",\"updated_at\":\"2026-02-06T05:54:01+08:00\",\"deleted_at\":null,\"hosts\":null},{\"id\":4,\"name\":\"centos-default\",\"username\":\"centos\",\"password\":\"$2a$10$example_hashed_password\",\"description\":\"CentOS系统默认用户凭据\",\"created_at\":\"2026-02-06T05:54:01+08:00\",\"updated_at\":\"2026-02-06T05:54:01+08:00\",\"deleted_at\":null,\"hosts\":null}]}}', 35, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (109, '2026-02-10 22:45:57.388', '2026-02-10 22:45:57.388', '更新主机', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'PUT', '/api/v1/hosts/:id', '{\"id\":1,\"hostname\":\"\",\"ip_address\":\"\",\"port\":22,\"os_type\":\"linux\",\"description\":\"\",\"credential_ids\":[1,2]}', '{\"success\":true,\"message\":\"主机更新成功\",\"data\":{\"id\":1,\"hostname\":\"web-server-01\",\"ip_address\":\"192.168.1.101\",\"port\":22,\"os_type\":\"linux\",\"cpu_cores\":8,\"memory_gb\":16,\"disk_space_gb\":500,\"group_id\":1,\"status\":2,\"monitoring_enabled\":1,\"last_heartbeat\":\"2026-02-05T04:06:27+08:00\",\"description\":\"生产环境Web服务器01\",\"created_by\":1,\"updated_by\":10,\"created_at\":\"2026-02-05T04:11:27+08:00\",\"updated_at\":\"2026-02-10T22:45:57.382+08:00\",\"deleted_at\":null,\"group\":{\"id\":1,\"name\":\"Web服务器组\",\"description\":\"用于存放Web应用服务器的主机组\",\"status\":1,\"created_by\":1,\"updated_by\":1,\"created_at\":\"2026-02-05T04:11:27+08:00\",\"updated_at\":\"2026-02-05T04:11:27+08:00\",\"deleted_at\":null,\"hosts\":null},\"credentials\":[{\"id\":1,\"name\":\"root-default\",\"username\":\"root\",\"password\":\"$2a$10$example_hashed_password\",\"description\":\"默认root用户凭据\",\"created_at\":\"2026-02-06T05:54:01+08:00\",\"updated_at\":\"2026-02-06T05:54:01+08:00\",\"deleted_at\":null,\"hosts\":null},{\"id\":2,\"name\":\"admin-default\",\"username\":\"admin\",\"password\":\"$2a$10$example_hashed_password\",\"description\":\"默认管理员凭据\",\"created_at\":\"2026-02-06T05:54:01+08:00\",\"updated_at\":\"2026-02-06T05:54:01+08:00\",\"deleted_at\":null,\"hosts\":null}]}}', 27, NULL, NULL, 1, NULL);
INSERT INTO `operation_logs` VALUES (110, '2026-02-10 22:46:37.027', '2026-02-10 22:46:37.027', '创建凭据', 10, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0', 200, 'POST', '/api/v1/credentials', '{\"name\":\"测试\",\"username\":\"root\",\"password\":\"Zhy20250730!\",\"description\":\"密码\"}', '{\"success\":true,\"message\":\"凭据创建成功\",\"data\":{\"id\":5,\"name\":\"测试\",\"username\":\"root\",\"password\":\"Zhy20250730!\",\"description\":\"密码\",\"created_at\":\"2026-02-10T22:46:37.022+08:00\",\"updated_at\":\"2026-02-10T22:46:37.022+08:00\",\"deleted_at\":null,\"hosts\":null}}', 5, NULL, NULL, 1, NULL);

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '请求路径',
  `method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '请求方法',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '权限描述',
  `status` tinyint NULL DEFAULT 1 COMMENT '请求路径',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_permission_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 66 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of permission
-- ----------------------------
INSERT INTO `permission` VALUES (1, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/login', 'POST', '用户登录', 1);
INSERT INTO `permission` VALUES (2, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/captcha', 'GET', '获取验证码', 1);
INSERT INTO `permission` VALUES (3, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/logout', 'POST', '退出登录', 1);
INSERT INTO `permission` VALUES (4, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/users', 'POST', '创建用户', 1);
INSERT INTO `permission` VALUES (5, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/users', 'GET', '获取用户列表', 1);
INSERT INTO `permission` VALUES (6, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/users/:id', 'GET', '获取用户信息', 1);
INSERT INTO `permission` VALUES (7, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/users/:id', 'PUT', '更新用户信息', 1);
INSERT INTO `permission` VALUES (8, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/users/:id/status', 'PUT', '更新用户状态', 1);
INSERT INTO `permission` VALUES (9, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/users/:id', 'DELETE', '删除用户', 1);
INSERT INTO `permission` VALUES (10, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/users/change-password', 'PUT', '修改密码', 1);
INSERT INTO `permission` VALUES (11, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/users/:id/reset-password', 'PUT', '重置密码', 1);
INSERT INTO `permission` VALUES (12, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/users-roles/:username', 'POST', '为用户分配角色', 1);
INSERT INTO `permission` VALUES (13, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/users-roles/:username', 'DELETE', '移除用户的角色', 1);
INSERT INTO `permission` VALUES (14, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/users-roles/:username', 'GET', '获取用户的角色列表', 1);
INSERT INTO `permission` VALUES (15, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/roles', 'POST', '创建角色', 1);
INSERT INTO `permission` VALUES (16, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/roles', 'GET', '获取角色列表', 1);
INSERT INTO `permission` VALUES (17, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/roles/:id', 'GET', '获取角色详情', 1);
INSERT INTO `permission` VALUES (18, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/roles/:id', 'PUT', '更新角色', 1);
INSERT INTO `permission` VALUES (19, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/roles/:id', 'DELETE', '删除角色', 1);
INSERT INTO `permission` VALUES (20, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/menus', 'POST', '创建菜单', 1);
INSERT INTO `permission` VALUES (21, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/menus', 'GET', '查询用户可见菜单', 1);
INSERT INTO `permission` VALUES (22, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/menus/all', 'GET', '查询所有菜单', 1);
INSERT INTO `permission` VALUES (23, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/menus/:id', 'PUT', '更新菜单', 1);
INSERT INTO `permission` VALUES (24, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/menus/:id', 'DELETE', '删除菜单', 1);
INSERT INTO `permission` VALUES (25, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/roles/:id/menus', 'POST', '为角色分配菜单权限', 1);
INSERT INTO `permission` VALUES (26, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/roles/:id/menus', 'GET', '获取角色的菜单权限', 1);
INSERT INTO `permission` VALUES (27, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/roles/:id/menus', 'DELETE', '移除角色的菜单权限', 1);
INSERT INTO `permission` VALUES (28, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/roles/:id/policies', 'POST', '添加Casbin策略', 1);
INSERT INTO `permission` VALUES (29, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/roles/:id/policies', 'DELETE', '移除Casbin策略', 1);
INSERT INTO `permission` VALUES (30, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/roles/:id/policies', 'GET', '获取角色的Casbin策略', 1);
INSERT INTO `permission` VALUES (31, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/permissions', 'POST', '创建权限', 1);
INSERT INTO `permission` VALUES (32, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/permissions', 'GET', '获取权限列表', 1);
INSERT INTO `permission` VALUES (33, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/permissions/:id', 'GET', '获取权限详情', 1);
INSERT INTO `permission` VALUES (34, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/permissions/:id', 'PUT', '更新权限', 1);
INSERT INTO `permission` VALUES (35, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/permissions/:id/status', 'PUT', '更新权限状态', 1);
INSERT INTO `permission` VALUES (36, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/permissions/:id', 'DELETE', '删除权限', 1);
INSERT INTO `permission` VALUES (37, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/permissions/all', 'GET', '获取所有权限（不分页）', 1);
INSERT INTO `permission` VALUES (38, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/operation-logs', 'GET', '查询操作日志', 1);
INSERT INTO `permission` VALUES (39, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/operation-logs/:id', 'DELETE', '删除操作日志', 1);
INSERT INTO `permission` VALUES (40, '2026-02-05 03:35:07.000', '2026-02-05 03:35:07.000', NULL, '/api/v1/users/profile', 'GET', '获取当前用户信息', 1);
INSERT INTO `permission` VALUES (41, '2026-02-05 11:29:09.000', NULL, NULL, '/api/v1/hosts', 'POST', '创建主机', 1);
INSERT INTO `permission` VALUES (42, '2026-02-05 11:29:11.000', NULL, NULL, '/api/v1/hosts', 'GET', '获取主机列表', 1);
INSERT INTO `permission` VALUES (43, '2026-02-05 11:29:14.000', NULL, NULL, '/api/v1/hosts/:id', 'GET', '获取主机详情', 1);
INSERT INTO `permission` VALUES (44, '2026-02-05 11:29:18.000', NULL, NULL, '/api/v1/hosts/:id', 'PUT', '更新主机信息', 1);
INSERT INTO `permission` VALUES (45, '2026-02-05 11:29:21.000', NULL, NULL, '/api/v1/hosts/:id', 'DELETE', '删除主机', 1);
INSERT INTO `permission` VALUES (46, '2026-02-05 11:29:24.000', NULL, NULL, '/api/v1/hosts/batch', 'DELETE', '批量删除主机', 1);
INSERT INTO `permission` VALUES (47, '2026-02-05 11:29:27.000', NULL, NULL, '/api/v1/hosts/:id/status', 'PUT', '更新主机状态', 1);
INSERT INTO `permission` VALUES (48, '2026-02-05 11:29:46.000', NULL, NULL, '/api/v1/hosts/:id/monitoring', 'PUT', '更新主机监控状态', 1);
INSERT INTO `permission` VALUES (49, '2026-02-05 11:29:50.000', NULL, NULL, '/api/v1/host-groups', 'POST', '创建主机组', 1);
INSERT INTO `permission` VALUES (50, '2026-02-05 11:29:53.000', NULL, NULL, '/api/v1/host-groups', 'GET', '获取主机组列表', 1);
INSERT INTO `permission` VALUES (51, '2026-02-05 11:29:56.000', NULL, NULL, '/api/v1/host-groups/:id', 'GET', '获取主机组详情', 1);
INSERT INTO `permission` VALUES (52, '2026-02-05 03:35:40.000', '2026-02-05 03:35:40.000', NULL, '/api/v1/host-groups/:id', 'PUT', '更新主机组', 1);
INSERT INTO `permission` VALUES (53, '2026-02-05 11:30:03.000', NULL, NULL, '/api/v1/host-groups/:id', 'DELETE', '删除主机组', 1);
INSERT INTO `permission` VALUES (54, '2026-02-05 11:30:05.000', NULL, NULL, '/api/v1/host-groups/:id/status', 'PUT', '更新主机组状态', 1);
INSERT INTO `permission` VALUES (55, '2026-02-05 11:30:08.000', NULL, NULL, '/api/v1/host-metrics', 'POST', '上报主机指标', 1);
INSERT INTO `permission` VALUES (56, '2026-02-05 11:30:11.000', NULL, NULL, '/api/v1/host-metrics/history', 'GET', '获取主机指标历史', 1);
INSERT INTO `permission` VALUES (57, '2026-02-05 11:30:13.000', NULL, NULL, '/api/v1/host-metrics/latest', 'GET', '获取主机最新指标', 1);
INSERT INTO `permission` VALUES (58, '2026-02-05 11:30:16.000', NULL, NULL, '/api/v1/hosts/statistics', 'GET', '获取主机统计信息', 1);
INSERT INTO `permission` VALUES (59, '2026-02-05 11:30:19.000', NULL, NULL, '/api/v1/credentials', 'POST', '创建凭据', 1);
INSERT INTO `permission` VALUES (60, '2026-02-05 11:30:22.000', NULL, NULL, '/api/v1/credentials', 'GET', '获取凭据列表', 1);
INSERT INTO `permission` VALUES (61, '2026-02-05 11:30:25.000', NULL, NULL, '/api/v1/credentials/:id', 'GET', '获取凭据详情', 1);
INSERT INTO `permission` VALUES (62, '2026-02-05 11:30:28.000', NULL, NULL, '/api/v1/credentials/:id', 'PUT', '更新凭据', 1);
INSERT INTO `permission` VALUES (63, '2026-02-05 11:30:31.000', NULL, NULL, '/api/v1/credentials/:id', 'DELETE', '删除凭据', 1);
INSERT INTO `permission` VALUES (64, '2026-02-05 11:30:34.000', NULL, NULL, '/api/v1/credentials/batch', 'DELETE', '批量删除凭据', 1);
INSERT INTO `permission` VALUES (65, '2026-02-05 11:30:37.000', NULL, NULL, '/api/v1/credentials/host', 'GET', '获取主机关联凭据', 1);

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
INSERT INTO `role_menus` VALUES (2, 3);
INSERT INTO `role_menus` VALUES (2, 11);
INSERT INTO `role_menus` VALUES (2, 12);
INSERT INTO `role_menus` VALUES (2, 13);
INSERT INTO `role_menus` VALUES (2, 14);
INSERT INTO `role_menus` VALUES (2, 15);
INSERT INTO `role_menus` VALUES (2, 118);
INSERT INTO `role_menus` VALUES (2, 119);
INSERT INTO `role_menus` VALUES (2, 120);
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
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
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
INSERT INTO `roles` VALUES (1, NULL, NULL, NULL, '超级管理员', '拥有系统所有权限', 1, 'surperadmin');
INSERT INTO `roles` VALUES (2, NULL, NULL, NULL, '管理员', '拥有大部分管理权限', 1, 'admin');
INSERT INTO `roles` VALUES (3, NULL, NULL, NULL, '普通用户', '拥有基本查看权限', 1, 'user');
INSERT INTO `roles` VALUES (4, NULL, NULL, NULL, '访客', '仅拥有查看权限', 1, 'Visitor');
INSERT INTO `roles` VALUES (5, NULL, NULL, NULL, '测试角色', '用于测试的角色', 1, 'ceshi');
INSERT INTO `roles` VALUES (7, NULL, NULL, NULL, 'TestRole', '', 1, 'test_role_ident');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3),
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
  `ident` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_users_username`(`username` ASC) USING BTREE,
  UNIQUE INDEX `idx_users_email`(`email` ASC) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_roles_users`(`role_id` ASC) USING BTREE,
  CONSTRAINT `fk_roles_users` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_users_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (10, '2026-02-10 22:31:42.074', '2026-02-10 22:31:42.074', NULL, 'admin', '$2a$10$ocZQAxtwX0K8aSywcLjICeMqyWv8KqvqJ7ZeoKSu9bMSIKjKhpkAq', 'admin2@example.com', '13800138009', '管理员', '', 1, '2026-02-10 22:31:42.073', '', 2, '', NULL);
INSERT INTO `users` VALUES (16, NULL, NULL, NULL, 'ceshi', '$2a$10$3J.aR7claboqAQuKPWeWUenxV.jY7Hjasi.eBBqYicedQgyzWG4KO', '2603485744@qq.com', '', '测试', '', 1, '2026-01-27 20:59:45.972', '', 5, '', NULL);
INSERT INTO `users` VALUES (22, NULL, NULL, NULL, '1111', '$2a$10$gwhoeJZElt8CrJQp5Kb64uig2X6uEc0BHrn36IIM/e7vE04tkD12e', '1111@12313.com', '', '1111', '', 1, NULL, '', NULL, NULL, NULL);
INSERT INTO `users` VALUES (23, '2026-02-05 03:40:17.941', '2026-02-05 03:43:07.705', NULL, '11111', '$2a$10$SstBKyW61Anu8rt7zwgSRuPSHFq.hqhYdpOR66jZ1dcUIXSDRVLI6', '1111111@123.com', '', '1111111', '', 1, NULL, '', NULL, NULL, NULL);

SET FOREIGN_KEY_CHECKS = 1;
