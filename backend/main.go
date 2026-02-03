package main

import (
	sysHandler "backend/api/handler/system_manager"
	"backend/api/router"
	"backend/config"
	sysRepository "backend/internal/repository/system_manager"
	sysService "backend/internal/service/system_manager"
	"backend/pkg/casbin"
	"backend/pkg/database"
	"backend/pkg/logger"
	"backend/pkg/redis"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	userRepo := sysRepository.NewUserRepository(db)
	roleRepo := sysRepository.NewRoleRepository(db)
	menuRepo := sysRepository.NewMenuRepository(db)
	operationLogRepo := sysRepository.NewOperationLogRepository(db)
	permissionRepo := sysRepository.NewPermissionRepository(db)
	roleMenuRepository := sysRepository.NewRoleMenuRepository(db)

	// 初始化服务层
	userService := sysService.NewUserService(userRepo, roleRepo)
	roleService := sysService.NewRoleService(roleRepo)
	menuService := sysService.NewMenuService(menuRepo)
	operationLogService := sysService.NewOperationLogService(operationLogRepo)
	permissionService := sysService.NewPermissionService(roleRepo, menuRepo, permissionRepo)
	roleMenuService := sysService.NewRoleMenuService(roleMenuRepository)

	// 初始化处理器层
	userHandler := sysHandler.NewUserHandler(userService)
	roleHandler := sysHandler.NewRoleHandler(roleService)
	menuHandler := sysHandler.NewMenuHandler(menuService)
	operationLogHandler := sysHandler.NewOperationLogHandler(operationLogService)
	permissionHandler := sysHandler.NewPermissionHandler(permissionService)
	roleMenuHandler := sysHandler.NewRoleMenuHandler(roleMenuService)

	// 设置路由
	r := router.SetupRouter(userHandler, roleHandler, menuHandler, operationLogHandler, permissionHandler, roleMenuHandler)

	// 启动服务器
	port := config.GlobalConfig.Server.Port

	// if err := r.Run(port); err != nil {
	// 	logger.Logger.Error("启动服务器失败", zap.Error(err))
	// 	os.Exit(1)
	// }

	// 创建 HTTP 服务器并应用超时配置
	srv := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  config.GlobalConfig.Server.ReadTimeout * time.Second,
		WriteTimeout: config.GlobalConfig.Server.WriteTimeout * time.Second,
	}

	// 启动服务器
	err = srv.ListenAndServe()
	if err != nil {
		logger.Logger.Error("启动服务器失败", zap.Error(err))
		os.Exit(1)
	}
	logger.Logger.Info(fmt.Sprintf("服务器启动在端口 %s", port))
}
