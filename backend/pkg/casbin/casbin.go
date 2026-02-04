package casbin

import (
	"backend/pkg/logger"
	"errors"

	"github.com/casbin/casbin/v3"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Enforcer *casbin.Enforcer

// InitCasbinWithGormAdapter 初始化Casbin（使用GORM适配器）
func InitCasbinWithGormAdapter(db *gorm.DB) error {
	// 使用GORM适配器（存储策略到数据库）
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		logger.Logger.Error("初始化Casbin GORM适配器失败", zap.Error(err))
		return errors.New("初始化Casbin GORM适配器失败")
	}

	// 创建Enforcer
	Enforcer, err = casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		logger.Logger.Error("创建Casbin Enforcer失败", zap.Error(err))
		return errors.New("创建Casbin Enforcer失败")
	}

	// 加载策略
	err = Enforcer.LoadPolicy()
	if err != nil {
		return errors.New("加载Casbin策略失败")
	}

	return nil
}

//检查权限
//添加策略
//删除策略
//为用户添加角色
//删除用户的角色
//获取用户的所有角色
//获取角色的所有用户
