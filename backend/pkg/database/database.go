package database

import (
	"backend/config"
	model "backend/internal/model/system_manager"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// LogLevelMap 定义日志级别字符串到 logger.LogLevel 的映射
var LogLevelMap = map[string]logger.LogLevel{
	"silent": logger.Silent,
	"error":  logger.Error,
	"warn":   logger.Warn,
	"info":   logger.Info,
}

// ParseLogLevel 将配置文件中的日志级别字符串转换为 logger.LogLevel
func ParseLogLevel(level string) (logger.LogLevel, error) {
	if lvl, ok := LogLevelMap[level]; ok {
		return lvl, nil
	}
	return logger.Info, fmt.Errorf("未知的日志级别: %s", level)
}

// InitDB 初始化数据库连接
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
	err = db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Menu{},
		&model.RoleMenu{},
		&model.OperationLog{},
		&model.Permission{},
	)
	if err != nil {
		return nil, fmt.Errorf("数据库迁移失败: %w", err)
	}

	return db, nil
}
