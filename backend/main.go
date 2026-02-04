package main

import (
	"backend/api/router"
	"backend/config"
	"backend/pkg/di"
	"backend/pkg/logger"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 初始化配置
	if err := config.InitConfig("config/config.yaml"); err != nil {
		logger.Logger.Error("初始化配置失败", zap.Error(err))
		os.Exit(1)
	}

	// 初始化日志
	if err := logger.InitLogger(); err != nil {
		logger.Logger.Error("初始化日志失败", zap.Error(err))
		os.Exit(1)
	}

	// 设置Gin模式
	gin.SetMode(config.GlobalConfig.Server.Mode)

	// 使用依赖注入容器初始化所有依赖
	container, err := di.InitializeContainer()
	if err != nil {
		logger.Logger.Error("初始化依赖注入容器失败", zap.Error(err))
		os.Exit(1)
	}
	defer func() {
		if closeErr := container.Close(); closeErr != nil {
			logger.Logger.Error("关闭资源失败", zap.Error(closeErr))
		}
	}()

	// 从容器获取处理器层组件（用于路由设置）
	userHandler := container.GetUserHandler()
	roleHandler := container.GetRoleHandler()
	menuHandler := container.GetMenuHandler()
	operationLogHandler := container.GetOperationLogHandler()
	permissionHandler := container.GetPermissionHandler()
	roleMenuHandler := container.GetRoleMenuHandler()
	hostHandler := container.GetHostHandler()
	hostGroupHandler := container.GetHostGroupHandler()
	operationLogService := container.GetOperationLogService()

	// 设置路由
	r := router.SetupRouter(userHandler, roleHandler, menuHandler, operationLogHandler, permissionHandler, roleMenuHandler, hostHandler, hostGroupHandler, operationLogService)

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
	logger.Logger.Info("服务器启动成功", zap.String("port", port))
}
