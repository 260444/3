-- 凭据表
-- 存储主机连接所需的认证信息，支持多个主机共享同一凭据
CREATE TABLE IF NOT EXISTS `credentials` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '凭据唯一标识符',
    `name` VARCHAR(100) NOT NULL COMMENT '凭据名称，用于标识和搜索',
    `username` VARCHAR(100) NOT NULL COMMENT '登录用户名',
    `password` TEXT NOT NULL COMMENT '加密后的密码',
    `description` TEXT COMMENT '凭据描述信息',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '凭据状态：1-启用，0-禁用',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_name` (`name`) COMMENT '凭据名称唯一索引',
    KEY `idx_status` (`status`) COMMENT '状态索引，用于快速筛选启用/禁用的凭据',
    KEY `idx_username` (`username`) COMMENT '用户名索引，便于按用户名查询'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='主机凭据信息表';

-- 主机凭据关联表
-- 实现主机与凭据的多对多关系，一个主机可以使用多个凭据，一个凭据可以被多个主机使用
CREATE TABLE IF NOT EXISTS `host_credentials` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '关联关系唯一标识符',
    `host_id` BIGINT UNSIGNED NOT NULL COMMENT '主机ID，关联hosts表',
    `credential_id` BIGINT UNSIGNED NOT NULL COMMENT '凭据ID，关联credentials表',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '关联创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_host_credential` (`host_id`, `credential_id`) COMMENT '主机-凭据组合唯一约束，防止重复关联',
    KEY `idx_host_id` (`host_id`) COMMENT '主机ID索引，便于查询某主机的所有凭据',
    KEY `idx_credential_id` (`credential_id`) COMMENT '凭据ID索引，便于查询使用某凭据的所有主机',
    CONSTRAINT `fk_host_credentials_host_id` FOREIGN KEY (`host_id`) REFERENCES `hosts` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_host_credentials_credential_id` FOREIGN KEY (`credential_id`) REFERENCES `credentials` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='主机与凭据关联关系表';

-- 插入默认凭据数据
-- 提供一些常用的默认凭据模板，方便快速配置
INSERT INTO `credentials` (`name`, `username`, `password`, `description`, `status`) VALUES
('root-default', 'root', '$2a$10$example_hashed_password', '默认root用户凭据', 1),
('admin-default', 'admin', '$2a$10$example_hashed_password', '默认管理员凭据', 1),
('ubuntu-default', 'ubuntu', '$2a$10$example_hashed_password', 'Ubuntu系统默认用户凭据', 1),
('centos-default', 'centos', '$2a$10$example_hashed_password', 'CentOS系统默认用户凭据', 1)
ON DUPLICATE KEY UPDATE 
    `username` = VALUES(`username`),
    `password` = VALUES(`password`),
    `description` = VALUES(`description`),
    `status` = VALUES(`status`),
    `updated_at` = CURRENT_TIMESTAMP;

-- 插入测试用的主机凭据关联数据
-- 为测试环境预设一些主机凭据关联关系
INSERT INTO `host_credentials` (`host_id`, `credential_id`) VALUES
(1, 1),  -- 主机1使用root凭据
(1, 2),  -- 主机1同时使用admin凭据
(2, 3),  -- 主机2使用ubuntu凭据
(3, 4)   -- 主机3使用centos凭据
ON DUPLICATE KEY UPDATE 
    `created_at` = CURRENT_TIMESTAMP;