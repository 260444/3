-- 主机管理模块测试数据
-- 执行前请确保已创建表结构

-- 清空现有数据（可选）
-- DELETE FROM host_metrics;
-- DELETE FROM hosts;
-- DELETE FROM host_groups;

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
-- Web服务器01的CPU指标（最近5分钟数据）
(1, 'cpu', 'cpu_usage', 45.5, '%', DATE_SUB(NOW(), INTERVAL 5 MINUTE)),
(1, 'cpu', 'cpu_usage', 48.2, '%', DATE_SUB(NOW(), INTERVAL 4 MINUTE)),
(1, 'cpu', 'cpu_usage', 42.8, '%', DATE_SUB(NOW(), INTERVAL 3 MINUTE)),
(1, 'cpu', 'cpu_usage', 51.3, '%', DATE_SUB(NOW(), INTERVAL 2 MINUTE)),
(1, 'cpu', 'cpu_usage', 47.9, '%', DATE_SUB(NOW(), INTERVAL 1 MINUTE)),

-- Web服务器01的内存指标（最近5分钟数据）
(1, 'memory', 'memory_usage', 68.2, '%', DATE_SUB(NOW(), INTERVAL 5 MINUTE)),
(1, 'memory', 'memory_usage', 71.5, '%', DATE_SUB(NOW(), INTERVAL 4 MINUTE)),
(1, 'memory', 'memory_usage', 65.8, '%', DATE_SUB(NOW(), INTERVAL 3 MINUTE)),
(1, 'memory', 'memory_usage', 73.1, '%', DATE_SUB(NOW(), INTERVAL 2 MINUTE)),
(1, 'memory', 'memory_usage', 69.7, '%', DATE_SUB(NOW(), INTERVAL 1 MINUTE)),

-- 数据库服务器的磁盘指标（最近5分钟数据）
(4, 'disk', 'disk_usage', 78.5, '%', DATE_SUB(NOW(), INTERVAL 5 MINUTE)),
(4, 'disk', 'disk_usage', 79.2, '%', DATE_SUB(NOW(), INTERVAL 4 MINUTE)),
(4, 'disk', 'disk_usage', 80.1, '%', DATE_SUB(NOW(), INTERVAL 3 MINUTE)),
(4, 'disk', 'disk_usage', 81.3, '%', DATE_SUB(NOW(), INTERVAL 2 MINUTE)),
(4, 'disk', 'disk_usage', 82.7, '%', DATE_SUB(NOW(), INTERVAL 1 MINUTE)),

-- 缓存服务器的网络指标（最近5分钟数据）
(7, 'network', 'network_in', 125.6, 'Mbps', DATE_SUB(NOW(), INTERVAL 5 MINUTE)),
(7, 'network', 'network_in', 132.4, 'Mbps', DATE_SUB(NOW(), INTERVAL 4 MINUTE)),
(7, 'network', 'network_in', 118.9, 'Mbps', DATE_SUB(NOW(), INTERVAL 3 MINUTE)),
(7, 'network', 'network_in', 145.2, 'Mbps', DATE_SUB(NOW(), INTERVAL 2 MINUTE)),
(7, 'network', 'network_in', 138.7, 'Mbps', DATE_SUB(NOW(), INTERVAL 1 MINUTE)),

-- Web服务器02的指标数据（用于对比测试）
(2, 'cpu', 'cpu_usage', 32.1, '%', DATE_SUB(NOW(), INTERVAL 1 MINUTE)),
(2, 'memory', 'memory_usage', 55.3, '%', DATE_SUB(NOW(), INTERVAL 1 MINUTE)),
(2, 'disk', 'disk_usage', 65.8, '%', DATE_SUB(NOW(), INTERVAL 1 MINUTE)),

-- 故障服务器的指标数据（历史数据）
(8, 'cpu', 'cpu_usage', 95.2, '%', DATE_SUB(NOW(), INTERVAL 2 HOUR)),
(8, 'memory', 'memory_usage', 98.7, '%', DATE_SUB(NOW(), INTERVAL 2 HOUR));

-- 验证数据插入
SELECT '=== 主机组数据 ===' as '';
SELECT id, name, description, status FROM host_groups;

SELECT '=== 主机数据统计 ===' as '';
SELECT 
    COUNT(*) as total_hosts,
    SUM(CASE WHEN status = 1 THEN 1 ELSE 0 END) as online_hosts,
    SUM(CASE WHEN status = 0 THEN 1 ELSE 0 END) as offline_hosts,
    SUM(CASE WHEN status = -1 THEN 1 ELSE 0 END) as fault_hosts,
    SUM(CASE WHEN monitoring_enabled = 1 THEN 1 ELSE 0 END) as monitoring_enabled
FROM hosts;

SELECT '=== 各主机组主机数量 ===' as '';
SELECT 
    hg.name as group_name,
    COUNT(h.id) as host_count
FROM host_groups hg
LEFT JOIN hosts h ON hg.id = h.group_id
GROUP BY hg.id, hg.name;

SELECT '=== 监控指标数据量 ===' as '';
SELECT 
    COUNT(*) as total_metrics,
    COUNT(DISTINCT host_id) as hosts_with_metrics
FROM host_metrics;