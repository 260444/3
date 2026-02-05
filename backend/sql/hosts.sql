CREATE TABLE `hosts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `hostname` varchar(100) NOT NULL COMMENT '主机名',
  `ip_address` varchar(45) NOT NULL COMMENT 'IP地址(IPv4或IPv6)',
  `port` smallint unsigned NOT NULL DEFAULT '22' COMMENT 'SSH端口',
  `username` varchar(50) NOT NULL COMMENT '登录用户名',
  `password` varchar(255) NOT NULL COMMENT '加密后的密码',
  `os_type` varchar(20) NOT NULL DEFAULT 'linux' COMMENT '操作系统类型: linux,windows',
  `cpu_cores` smallint unsigned DEFAULT NULL COMMENT 'CPU核心数',
  `memory_gb` smallint unsigned DEFAULT NULL COMMENT '内存大小(GB)',
  `disk_space_gb` int unsigned DEFAULT NULL COMMENT '磁盘空间(GB)',
  `group_id` bigint unsigned NOT NULL COMMENT '所属主机组ID',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '主机状态: 1-在线, 0-离线, -1-故障',
  `monitoring_enabled` tinyint NOT NULL DEFAULT '1' COMMENT '监控是否启用: 1-启用, 0-禁用',
  `last_heartbeat` timestamp NULL DEFAULT NULL COMMENT '最后心跳时间',
  `description` varchar(500) DEFAULT NULL COMMENT '主机描述',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人用户ID',
  `updated_by` bigint unsigned DEFAULT NULL COMMENT '更新人用户ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_hosts_hostname` (`hostname`),
  UNIQUE KEY `uk_hosts_ip_address` (`ip_address`),
  KEY `idx_hosts_group_id` (`group_id`),
  KEY `idx_hosts_status` (`status`),
  KEY `idx_hosts_os_type` (`os_type`),
  KEY `idx_hosts_monitoring_enabled` (`monitoring_enabled`),
  KEY `idx_hosts_last_heartbeat` (`last_heartbeat`),
  KEY `idx_hosts_created_at` (`created_at`),
  KEY `idx_hosts_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='主机表';

CREATE TABLE `host_groups` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(100) NOT NULL COMMENT '主机组名称',
  `description` varchar(500) DEFAULT NULL COMMENT '描述信息',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态: 1-启用, 0-禁用',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人用户ID',
  `updated_by` bigint unsigned DEFAULT NULL COMMENT '更新人用户ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_host_groups_name` (`name`),
  KEY `idx_host_groups_status` (`status`),
  KEY `idx_host_groups_created_at` (`created_at`),
  KEY `idx_host_groups_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='主机组表';

CREATE TABLE `host_metrics` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `host_id` bigint unsigned NOT NULL COMMENT '主机ID',
  `metric_type` varchar(30) NOT NULL COMMENT '指标类型: cpu,memory,disk,network',
  `metric_name` varchar(50) NOT NULL COMMENT '指标名称',
  `metric_value` decimal(10,2) NOT NULL COMMENT '指标值',
  `unit` varchar(20) DEFAULT NULL COMMENT '单位',
  `recorded_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录时间',
  PRIMARY KEY (`id`),
  KEY `idx_host_metrics_host_id` (`host_id`),
  KEY `idx_host_metrics_metric_type` (`metric_type`),
  KEY `idx_host_metrics_recorded_at` (`recorded_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='主机监控指标表';

-- 插入测试数据

-- 插入主机组测试数据
INSERT INTO `host_groups` (`id`, `name`, `description`, `status`, `created_by`, `updated_by`, `created_at`, `updated_at`) VALUES
(1, 'Web服务器组', '用于存放Web应用服务器的主机组', 1, 1, 1, NOW(), NOW()),
(2, '数据库服务器组', '用于存放数据库服务器的主机组', 1, 1, 1, NOW(), NOW()),
(3, '缓存服务器组', '用于存放Redis、Memcached等缓存服务器的主机组', 1, 1, 1, NOW(), NOW()),
(4, '存储服务器组', '用于存放文件存储和备份服务器的主机组', 1, 1, 1, NOW(), NOW());

-- 插入主机测试数据
INSERT INTO `hosts` (`id`, `hostname`, `ip_address`, `port`, `username`, `password`, `os_type`, `cpu_cores`, `memory_gb`, `disk_space_gb`, `group_id`, `status`, `monitoring_enabled`, `last_heartbeat`, `description`, `created_by`, `updated_by`, `created_at`, `updated_at`) VALUES
(1, 'web-server-01', '192.168.1.101', 22, 'root', '$2a$10$example_hash_01', 'linux', 8, 16, 500, 1, 1, 1, DATE_SUB(NOW(), INTERVAL 5 MINUTE), '生产环境Web服务器01', 1, 1, NOW(), NOW()),
(2, 'web-server-02', '192.168.1.102', 22, 'root', '$2a$10$example_hash_02', 'linux', 8, 16, 500, 1, 1, 1, DATE_SUB(NOW(), INTERVAL 3 MINUTE), '生产环境Web服务器02', 1, 1, NOW(), NOW()),
(3, 'web-server-03', '192.168.1.103', 22, 'root', '$2a$10$example_hash_03', 'linux', 16, 32, 1000, 1, 1, 1, DATE_SUB(NOW(), INTERVAL 1 MINUTE), '生产环境Web服务器03', 1, 1, NOW(), NOW()),
(4, 'db-master-01', '192.168.2.101', 22, 'root', '$2a$10$example_hash_04', 'linux', 16, 64, 2000, 2, 1, 1, DATE_SUB(NOW(), INTERVAL 2 MINUTE), '主数据库服务器', 1, 1, NOW(), NOW()),
(5, 'db-slave-01', '192.168.2.102', 22, 'root', '$2a$10$example_hash_05', 'linux', 16, 64, 2000, 2, 1, 1, DATE_SUB(NOW(), INTERVAL 4 MINUTE), '从数据库服务器01', 1, 1, NOW(), NOW()),
(6, 'db-slave-02', '192.168.2.103', 22, 'root', '$2a$10$example_hash_06', 'linux', 16, 64, 2000, 2, 0, 1, DATE_SUB(NOW(), INTERVAL 30 MINUTE), '从数据库服务器02（离线）', 1, 1, NOW(), NOW()),
(7, 'redis-cache-01', '192.168.3.101', 22, 'root', '$2a$10$example_hash_07', 'linux', 8, 32, 500, 3, 1, 1, DATE_SUB(NOW(), INTERVAL 1 MINUTE), 'Redis缓存服务器01', 1, 1, NOW(), NOW()),
(8, 'redis-cache-02', '192.168.3.102', 22, 'root', '$2a$10$example_hash_08', 'linux', 8, 32, 500, 3, -1, 1, DATE_SUB(NOW(), INTERVAL 2 HOUR), 'Redis缓存服务器02（故障）', 1, 1, NOW(), NOW()),
(9, 'storage-nfs-01', '192.168.4.101', 22, 'root', '$2a$10$example_hash_09', 'linux', 12, 64, 5000, 4, 1, 0, DATE_SUB(NOW(), INTERVAL 10 MINUTE), 'NFS文件存储服务器（监控已禁用）', 1, 1, NOW(), NOW()),
(10, 'backup-server-01', '192.168.4.102', 22, 'backup', '$2a$10$example_hash_10', 'windows', 8, 16, 3000, 4, 1, 1, DATE_SUB(NOW(), INTERVAL 15 MINUTE), 'Windows备份服务器', 1, 1, NOW(), NOW());

-- 插入主机监控指标测试数据
INSERT INTO `host_metrics` (`host_id`, `metric_type`, `metric_name`, `metric_value`, `unit`, `recorded_at`) VALUES
-- Web服务器01的CPU指标
(1, 'cpu', 'cpu_usage', 45.5, '%', DATE_SUB(NOW(), INTERVAL 5 MINUTE)),
(1, 'cpu', 'cpu_usage', 48.2, '%', DATE_SUB(NOW(), INTERVAL 4 MINUTE)),
(1, 'cpu', 'cpu_usage', 42.8, '%', DATE_SUB(NOW(), INTERVAL 3 MINUTE)),
(1, 'cpu', 'cpu_usage', 51.3, '%', DATE_SUB(NOW(), INTERVAL 2 MINUTE)),
(1, 'cpu', 'cpu_usage', 47.9, '%', DATE_SUB(NOW(), INTERVAL 1 MINUTE)),

-- Web服务器01的内存指标
(1, 'memory', 'memory_usage', 68.2, '%', DATE_SUB(NOW(), INTERVAL 5 MINUTE)),
(1, 'memory', 'memory_usage', 71.5, '%', DATE_SUB(NOW(), INTERVAL 4 MINUTE)),
(1, 'memory', 'memory_usage', 65.8, '%', DATE_SUB(NOW(), INTERVAL 3 MINUTE)),
(1, 'memory', 'memory_usage', 73.1, '%', DATE_SUB(NOW(), INTERVAL 2 MINUTE)),
(1, 'memory', 'memory_usage', 69.7, '%', DATE_SUB(NOW(), INTERVAL 1 MINUTE)),

-- 数据库服务器的磁盘指标
(4, 'disk', 'disk_usage', 78.5, '%', DATE_SUB(NOW(), INTERVAL 5 MINUTE)),
(4, 'disk', 'disk_usage', 79.2, '%', DATE_SUB(NOW(), INTERVAL 4 MINUTE)),
(4, 'disk', 'disk_usage', 80.1, '%', DATE_SUB(NOW(), INTERVAL 3 MINUTE)),
(4, 'disk', 'disk_usage', 81.3, '%', DATE_SUB(NOW(), INTERVAL 2 MINUTE)),
(4, 'disk', 'disk_usage', 82.7, '%', DATE_SUB(NOW(), INTERVAL 1 MINUTE)),

-- 缓存服务器的网络指标
(7, 'network', 'network_in', 125.6, 'Mbps', DATE_SUB(NOW(), INTERVAL 5 MINUTE)),
(7, 'network', 'network_in', 132.4, 'Mbps', DATE_SUB(NOW(), INTERVAL 4 MINUTE)),
(7, 'network', 'network_in', 118.9, 'Mbps', DATE_SUB(NOW(), INTERVAL 3 MINUTE)),
(7, 'network', 'network_in', 145.2, 'Mbps', DATE_SUB(NOW(), INTERVAL 2 MINUTE)),
(7, 'network', 'network_in', 138.7, 'Mbps', DATE_SUB(NOW(), INTERVAL 1 MINUTE));