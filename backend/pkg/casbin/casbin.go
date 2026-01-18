package casbin

import (
	"fmt"

	"github.com/casbin/casbin/v3"
	fileadapter "github.com/casbin/casbin/v3/persist/file-adapter"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var Enforcer *casbin.Enforcer

// InitCasbinWithGormAdapter 初始化Casbin（使用GORM适配器）
func InitCasbinWithGormAdapter(db *gorm.DB) error {
	// 使用GORM适配器（存储策略到数据库）
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("初始化Casbin GORM适配器失败: %w", err)
	}

	// 创建Enforcer
	Enforcer, err = casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("创建Casbin Enforcer失败: %w", err)
	}

	// 加载策略
	err = Enforcer.LoadPolicy()
	if err != nil {
		return fmt.Errorf("加载Casbin策略失败: %w", err)
	}

	return nil
}

// InitCasbinWithFileAdapter 初始化Casbin（使用文件适配器）
func InitCasbinWithFileAdapter() error {
	// 使用文件适配器
	adapter := fileadapter.NewAdapter("config/policy.csv")

	var err error
	Enforcer, err = casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("创建Casbin Enforcer失败: %w", err)
	}

	// 加载策略
	err = Enforcer.LoadPolicy()
	if err != nil {
		return fmt.Errorf("加载Casbin策略失败: %w", err)
	}

	return nil
}
