package router

import (
	assHandler "backend/api/handler/asset_management"
	sysHandler "backend/api/handler/system_manager"
	"backend/api/middleware"
	sysService "backend/internal/service/system_manager"
	"backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userHandler *sysHandler.UserHandler,
	roleHandler *sysHandler.RoleHandler,
	menuHandler *sysHandler.MenuHandler,
	operationLogHandler *sysHandler.OperationLogHandler,
	permissionHandler *sysHandler.PermissionHandler,
	roleMenuHandler *sysHandler.RoleMenuHandler,
	hostHandler *assHandler.HostHandler,
	hostGroupHandler *assHandler.HostGroupHandler,
	credentialHandler *assHandler.CredentialHandler,
	operationLogService *sysService.OperationLogService,

) *gin.Engine {
	r := gin.New()

	// 使用全局中间件
	r.Use(gin.Logger())
	r.Use(middleware.RecoveryMiddleware())
	r.Use(middleware.CORSMiddleware())

	// 使用日志中间件（全局应用）
	// r.Use(middleware.LoggerToFile())

	// 公开路由（不需要认证）
	public := r.Group("/api/v1")
	{
		public.POST("/login", middleware.OperationLogMiddleware(operationLogService, "用户登录"), userHandler.Login)
		public.GET("/captcha", GenerateCaptcha) // 验证码接口
	}

	// 需要认证的路由
	protected := r.Group("/api/v1")

	// 使用Casbin中间件进行权限控制
	protected.Use(middleware.JWTAuthMiddleware())
	protected.Use(middleware.CasbinMiddleware())
	{
		// 用户相关路由
		protected.POST("/logout", userHandler.Logout)
		protected.POST("/users", middleware.OperationLogMiddleware(operationLogService, "创建用户"), userHandler.CreateUser)
		protected.GET("/users", userHandler.GetUsers)
		protected.PUT("/users/:id", middleware.OperationLogMiddleware(operationLogService, "更新用户"), userHandler.UpdateUser)
		protected.DELETE("/users/:id", middleware.OperationLogMiddleware(operationLogService, "删除用户"), userHandler.DeleteUser)
		// TODO 有bug，每次获取的都是当前用户信息
		protected.GET("/users/:id", userHandler.GetUserInfo)
		protected.PUT("/users/:id/status", middleware.OperationLogMiddleware(operationLogService, "更新用户状态"), userHandler.UpdateUserStatus)
		protected.PUT("/users/change-password", middleware.OperationLogMiddleware(operationLogService, "修改密码"), userHandler.ChangePassword)
		protected.PUT("/users/:id/reset-password", middleware.OperationLogMiddleware(operationLogService, "重置密码"), userHandler.ResetPassword)

		// 角色分配相关路由
		protected.POST("/users-roles/:username", middleware.OperationLogMiddleware(operationLogService, "分配用户角色"), userHandler.AssignRole)   // 为用户分配角色
		protected.DELETE("/users-roles/:username", middleware.OperationLogMiddleware(operationLogService, "移除用户角色"), userHandler.RemoveRole) // 移除用户的角色
		protected.GET("/users-roles/:username", userHandler.GetUserRoles)                                                                    // 获取用户的角色列表

		// 角色相关路由
		protected.POST("/roles", middleware.OperationLogMiddleware(operationLogService, "创建角色"), roleHandler.CreateRole)
		protected.GET("/roles", roleHandler.GetRoles)
		protected.GET("/roles/:id", roleHandler.GetRole)
		protected.PUT("/roles/:id", middleware.OperationLogMiddleware(operationLogService, "更新角色"), roleHandler.UpdateRole)
		protected.DELETE("/roles/:id", middleware.OperationLogMiddleware(operationLogService, "删除角色"), roleHandler.DeleteRole)

		// 菜单相关路由
		protected.POST("/menus", middleware.OperationLogMiddleware(operationLogService, "创建菜单"), menuHandler.CreateMenu)
		// 查询用户可见菜单（包含子菜单）
		protected.GET("/menus", menuHandler.GetUserMenus)
		// 查询所有菜单（包含子菜单）
		protected.GET("/menus/all", menuHandler.GetAllMenus)
		protected.PUT("/menus/:id", middleware.OperationLogMiddleware(operationLogService, "更新菜单"), menuHandler.UpdateMenu)
		protected.DELETE("/menus/:id", middleware.OperationLogMiddleware(operationLogService, "删除菜单"), menuHandler.DeleteMenu)

		// 操作日志相关路由
		protected.GET("/operation-logs", operationLogHandler.GetOperationLogs)
		protected.DELETE("/operation-logs/:id", middleware.OperationLogMiddleware(operationLogService, "删除操作日志"), operationLogHandler.DeleteOperationLog)

		// 分配菜单管理相关路由
		protected.POST("/roles/:id/menus", middleware.OperationLogMiddleware(operationLogService, "为角色分配菜单"), roleMenuHandler.AssignMenuToRole) // 为角色分配菜单权限
		protected.GET("/roles/:id/menus", roleMenuHandler.GetRoleMenus)                                                                         // 获取角色的菜单权限
		// protected.PUT("/roles/menus", roleMenuHandler.UpdateMenuInRole)
		protected.DELETE("/roles/:id/menus", middleware.OperationLogMiddleware(operationLogService, "移除角色菜单"), roleMenuHandler.RemoveMenuFromRole) // 移除角色的菜单权限

		// 分配权限资源管理相关路由
		protected.POST("/roles/:id/policies", middleware.OperationLogMiddleware(operationLogService, "添加Casbin策略"), permissionHandler.AddPolicy)      // 添加Casbin策略
		protected.DELETE("/roles/:id/policies", middleware.OperationLogMiddleware(operationLogService, "移除Casbin策略"), permissionHandler.RemovePolicy) // 移除Casbin策略
		protected.GET("/roles/:id/policies", permissionHandler.GetPolicies)                                                                           // 获取角色的Casbin策略

		// 权限资源管理相关路由
		protected.POST("/permissions", middleware.OperationLogMiddleware(operationLogService, "创建权限"), permissionHandler.CreatePermission)                   // 创建权限
		protected.GET("/permissions", permissionHandler.GetPermissions)                                                                                      // 分页查询，获取权限列表
		protected.GET("/permissions/all", permissionHandler.GetAllPermissions)                                                                               // 不进行分页查询，获取所有权限
		protected.GET("/permissions/:id", permissionHandler.GetPermission)                                                                                   // 获取权限详情
		protected.PUT("/permissions/:id", middleware.OperationLogMiddleware(operationLogService, "更新权限"), permissionHandler.UpdatePermission)                // 更新权限
		protected.PUT("/permissions/:id/status", middleware.OperationLogMiddleware(operationLogService, "更新权限状态"), permissionHandler.UpdatePermissionStatus) // 更新权限状态
		protected.DELETE("/permissions/:id", middleware.OperationLogMiddleware(operationLogService, "删除权限"), permissionHandler.DeletePermission)             // 删除权限

		// 获取当前用户信息
		protected.GET("/users/profile", userHandler.GetCurrentUser)

		// 主机管理相关路由
		protected.POST("/hosts", middleware.OperationLogMiddleware(operationLogService, "创建主机"), hostHandler.CreateHost)
		protected.GET("/hosts", hostHandler.GetHostList)
		protected.GET("/hosts/:id", hostHandler.GetHostByID)
		protected.PUT("/hosts/:id", middleware.OperationLogMiddleware(operationLogService, "更新主机"), hostHandler.UpdateHost)
		protected.DELETE("/hosts/:id", middleware.OperationLogMiddleware(operationLogService, "删除主机"), hostHandler.DeleteHost)
		protected.DELETE("/hosts/batch", middleware.OperationLogMiddleware(operationLogService, "批量删除主机"), hostHandler.BatchDeleteHosts)
		protected.PUT("/hosts/:id/status", middleware.OperationLogMiddleware(operationLogService, "更新主机状态"), hostHandler.UpdateHostStatus)
		protected.PUT("/hosts/:id/monitoring", middleware.OperationLogMiddleware(operationLogService, "更新主机监控状态"), hostHandler.UpdateHostMonitoring)
		protected.GET("/hosts/statistics", hostHandler.GetHostStatistics)

		// 主机组管理相关路由
		protected.POST("/host-groups", middleware.OperationLogMiddleware(operationLogService, "创建主机组"), hostGroupHandler.CreateHostGroup)
		protected.GET("/host-groups", hostGroupHandler.GetHostGroupList)
		protected.GET("/host-groups/:id", hostGroupHandler.GetHostGroupByID)
		protected.PUT("/host-groups/:id", middleware.OperationLogMiddleware(operationLogService, "更新主机组"), hostGroupHandler.UpdateHostGroup)
		protected.DELETE("/host-groups/:id", middleware.OperationLogMiddleware(operationLogService, "删除主机组"), hostGroupHandler.DeleteHostGroup)
		protected.PUT("/host-groups/:id/status", middleware.OperationLogMiddleware(operationLogService, "更新主机组状态"), hostGroupHandler.UpdateHostGroupStatus)

		// 主机监控指标相关路由
		protected.POST("/host-metrics", middleware.OperationLogMiddleware(operationLogService, "上报主机指标"), hostHandler.ReportHostMetrics)
		protected.GET("/host-metrics/history", hostHandler.GetHostMetricsHistory)
		protected.GET("/host-metrics/latest", hostHandler.GetHostLatestMetrics)

		// 凭据管理相关路由
		protected.POST("/credentials", middleware.OperationLogMiddleware(operationLogService, "创建凭据"), credentialHandler.CreateCredential)
		protected.GET("/credentials", credentialHandler.GetCredentialList)
		protected.GET("/credentials/:id", credentialHandler.GetCredentialByID)
		protected.PUT("/credentials/:id", middleware.OperationLogMiddleware(operationLogService, "更新凭据"), credentialHandler.UpdateCredential)
		protected.DELETE("/credentials/:id", middleware.OperationLogMiddleware(operationLogService, "删除凭据"), credentialHandler.DeleteCredential)
		protected.DELETE("/credentials/batch", middleware.OperationLogMiddleware(operationLogService, "批量删除凭据"), credentialHandler.BatchDeleteCredentials)
		protected.GET("/credentials/host", credentialHandler.GetCredentialsByHost)

		return r
	}
}

// GenerateCaptcha 生成验证码
func GenerateCaptcha(c *gin.Context) {
	// 生成验证码
	id, b64s, err := utils.GenerateCaptcha()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "验证码生成失败",
		})
		return
	}

	// 返回验证码数据
	c.JSON(200, gin.H{
		"message": "验证码生成成功",
		"data": gin.H{
			"id":    id,
			"image": b64s,
		},
	})
}
