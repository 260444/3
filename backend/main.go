// Package main 是企业级后台管理系统的入口包。
//
// 本系统提供用户管理、角色管理、菜单管理、权限分配、操作日志以及资产管理等功能。
// 使用 Gin 框架作为 Web 服务器，采用分层架构设计（Handler-Service-Repository），
// 集成了 Casbin 进行权限控制，使用 JWT 进行身份认证。
//
// 技术栈：
//   - Web 框架：Gin v1.11.0
//   - ORM：GORM v1.31.1
//   - 数据库：MySQL 8.0
//   - 缓存：Redis 7.2
//   - 权限：Casbin v3.8.1
//   - 日志：Zap v1.27.1
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

// main 是应用程序的入口函数。
//
// 该函数负责：
//  1. 初始化配置文件
//  2. 初始化日志系统
//  3. 设置 Gin 运行模式
//  4. 初始化依赖注入容器
//  5. 设置路由
//  6. 启动 HTTP 服务器
//
// 如果初始化过程中出现错误，程序会记录错误日志并退出。
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
	credentialHandler := container.GetCredentialHandler()
	sshHandler := container.GetSSHHandler()
	operationLogService := container.GetOperationLogService()
	deploymentHandler := container.GetDeploymentHandler()

	// 设置路由
	r := router.SetupRouter(
		userHandler,
		roleHandler,
		menuHandler,
		operationLogHandler,
		permissionHandler,
		roleMenuHandler,
		hostHandler,
		hostGroupHandler,
		credentialHandler,
		sshHandler,
		operationLogService,
		deploymentHandler)

	// 创建 HTTP 服务器并应用超时配置
	port := config.GlobalConfig.Server.Port
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
