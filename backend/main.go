package main

import (
	"backend/api/handler"
	"backend/api/router"
	"backend/config"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/pkg/casbin"
	"backend/pkg/database"
	"backend/pkg/logger"
	"backend/pkg/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

func main() {
	// 初始化配置
	if err := config.InitConfig("config/config.yaml"); err != nil {
		fmt.Printf("初始化配置失败: %v\n", err)
		os.Exit(1)
	}

	// 初始化日志
	if err := logger.InitLogger(); err != nil {
		fmt.Printf("初始化日志失败: %v\n", err)
		os.Exit(1)
	}

	// 设置Gin模式
	gin.SetMode(config.GlobalConfig.Server.Mode)

	// 初始化数据库
	db, err := database.InitDB()
	if err != nil {
		logger.Logger.Error("初始化数据库失败", zap.Error(err))
		os.Exit(1)
	}

	// 初始化Redis (可选)
	if err := redis.InitRedis(); err != nil {
		logger.Logger.Warn("初始化Redis失败", zap.Error(err))
		// Redis初始化失败不影响系统运行，仅记录警告
	}

	// 初始化Casbin权限管理
	if err := casbin.InitCasbinWithGormAdapter(db); err != nil {
		logger.Logger.Error("初始化Casbin失败", zap.Error(err))
		os.Exit(1)
	}

	// 初始化仓库层
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	menuRepo := repository.NewMenuRepository(db)
	operationLogRepo := repository.NewOperationLogRepository(db)
	permissionRepo := repository.NewPermissionRepository(db)

	// 初始化服务层
	userService := service.NewUserService(userRepo, roleRepo)
	roleService := service.NewRoleService(roleRepo)
	menuService := service.NewMenuService(menuRepo)
	operationLogService := service.NewOperationLogService(operationLogRepo)
	permissionService := service.NewPermissionService(roleRepo, menuRepo, permissionRepo)

	// 初始化处理器层
	userHandler := handler.NewUserHandler(userService)
	roleHandler := handler.NewRoleHandler(roleService)
	menuHandler := handler.NewMenuHandler(menuService)
	operationLogHandler := handler.NewOperationLogHandler(operationLogService)
	permissionHandler := handler.NewPermissionHandler(permissionService)

	// 设置路由
	r := router.SetupRouter(userHandler, roleHandler, menuHandler, operationLogHandler, permissionHandler)

	// 启动服务器
	port := config.GlobalConfig.Server.Port
	logger.Logger.Info(fmt.Sprintf("服务器启动在端口 %s", port))
	if err := r.Run(port); err != nil {
		logger.Logger.Error("启动服务器失败", zap.Error(err))
		os.Exit(1)
	}
}
