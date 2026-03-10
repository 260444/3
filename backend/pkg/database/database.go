// Package database 提供数据库连接和初始化功能。
//
// 该包使用 GORM 作为 ORM 框架，支持 MySQL 数据库。
// 提供数据库连接初始化、日志配置和自动迁移功能。
//
// 使用示例：
//
//	db, err := database.InitDB()
//	if err != nil {
//	    log.Fatal(err)
//	}
package database

import (
	"backend/config"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// LogLevelMap 定义日志级别字符串到 logger.LogLevel 的映射。
var LogLevelMap = map[string]logger.LogLevel{
	"silent": logger.Silent,
	"error":  logger.Error,
	"warn":   logger.Warn,
	"info":   logger.Info,
}

// ParseLogLevel 将配置文件中的日志级别字符串转换为 logger.LogLevel。
//
// 支持的级别：silent, error, warn, info
// 如果传入未知级别，返回默认的 Info 级别。
func ParseLogLevel(level string) (logger.LogLevel, error) {
	if lvl, ok := LogLevelMap[level]; ok {
		return lvl, nil
	}
	return logger.Info, fmt.Errorf("未知的日志级别: %s", level)
}

// InitDB 初始化数据库连接并返回 GORM 数据库实例。
//
// 该函数执行以下操作：
//  1. 根据配置构建 MySQL DSN
//  2. 配置 GORM 日志（输出到文件）
//  3. 建立数据库连接
//  4. 自动迁移数据库表结构
//
// 返回的数据库实例已经配置好日志和连接池，可以直接使用。
//
// 返回的表结构包括：User, Role, Menu, RoleMenu, OperationLog, Permission
func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.GlobalConfig.Database.User,
		config.GlobalConfig.Database.Password,
		config.GlobalConfig.Database.Host,
		config.GlobalConfig.Database.Port,
		config.GlobalConfig.Database.DBName,
		config.GlobalConfig.Database.Charset,
	)

	file, err := os.OpenFile(config.GlobalConfig.Database.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("failed to open log file")
	}

	// 解析配置文件中的日志级别
	logLevel, err := ParseLogLevel(config.GlobalConfig.Database.LogLevel)
	if err != nil {
		return nil, err
	}

	// 自定义日志配置
	newLogger := logger.New(
		log.New(file, "\r\n", log.LstdFlags), // 输出目标
		logger.Config{
			SlowThreshold: time.Second, // 慢查询阈值
			LogLevel:      logLevel,    // 日志级别
			Colorful:      false,       // 是否彩色输出
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %w", err)
	}

	// 自动迁移数据库表
	//err = db.AutoMigrate(
	//	&model.User{},
	//	&model.Role{},
	//	&model.Menu{},
	//	&model.RoleMenu{},
	//	&model.OperationLog{},
	//	&model.Permission{},
	//)
	//if err != nil {
	//	return nil, fmt.Errorf("数据库迁移失败: %w", err)
	//}

	return db, nil
}
