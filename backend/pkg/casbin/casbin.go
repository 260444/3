// Package casbin 提供基于 Casbin 的 RBAC 权限控制功能。
//
// 该包封装了 Casbin 的核心功能，包括策略管理、角色分配、权限检查等。
// 使用 GORM 适配器将策略存储在数据库中。
//
// 使用示例：
//
//	err := casbin.InitCasbinWithGormAdapter(db)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	allowed, _ := casbin.Enforcer.Enforce("alice", "data1", "read")
package casbin

import (
	"backend/pkg/logger"
	"errors"

	"github.com/casbin/casbin/v3"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Enforcer 是 Casbin 的全局执行器实例。
//
// 该实例用于执行所有权限检查和策略管理操作。
var Enforcer *casbin.Enforcer

// InitCasbinWithGormAdapter 初始化 Casbin（使用 GORM 适配器）。
//
// 该函数会创建 Casbin Enforcer 实例，使用 GORM 适配器将策略存储在数据库中。
// 初始化后会自动加载策略。
//
// 参数：
//   - db: GORM 数据库实例
//
// 返回：
//   - error: 如果初始化失败，返回错误
func InitCasbinWithGormAdapter(db *gorm.DB) error {
	// 使用GORM适配器（存储策略到数据库）
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return err
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
