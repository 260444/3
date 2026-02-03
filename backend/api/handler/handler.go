package handler

import (
	sysHandler "backend/api/handler/system_manager"
	sysService "backend/internal/service/system_manager"
)

// HandlerGroup 包含所有处理器
type HandlerGroup struct {
	UserHandler         *sysHandler.UserHandler
	RoleHandler         *sysHandler.RoleHandler
	MenuHandler         *sysHandler.MenuHandler
	OperationLogHandler *sysHandler.OperationLogHandler
	PermissionHandler   *sysHandler.PermissionHandler
	RoleMenuHandler     *sysHandler.RoleMenuHandler
}

// NewHandlerGroup 创建所有处理器组
func NewHandlerGroup(
	userService *sysService.UserService,
	roleService *sysService.RoleService,
	menuService *sysService.MenuService,
	operationLogService *sysService.OperationLogService,
	permissionService *sysService.PermissionService,
	roleMenuService *sysService.RoleMenuService,
) *HandlerGroup {
	return &HandlerGroup{
		UserHandler:         sysHandler.NewUserHandler(userService),
		RoleHandler:         sysHandler.NewRoleHandler(roleService),
		MenuHandler:         sysHandler.NewMenuHandler(menuService),
		OperationLogHandler: sysHandler.NewOperationLogHandler(operationLogService),
		PermissionHandler:   sysHandler.NewPermissionHandler(permissionService),
		RoleMenuHandler:     sysHandler.NewRoleMenuHandler(roleMenuService),
	}
}
