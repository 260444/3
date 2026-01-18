package database

import (
	"backend/config"
	"backend/internal/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 根据配置决定日志级别
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
	)
	if err != nil {
		return nil, fmt.Errorf("数据库迁移失败: %w", err)
	}

	return db, nil
}
