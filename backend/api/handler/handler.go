package handler

import (
	"backend/internal/service"
)

// HandlerGroup 包含所有处理器
type HandlerGroup struct {
	UserHandler         *UserHandler
	RoleHandler         *RoleHandler
	MenuHandler         *MenuHandler
	OperationLogHandler *OperationLogHandler
	PermissionHandler   *PermissionHandler
	RoleMenuHandler     *RoleMenuHandler
}

// NewHandlerGroup 创建所有处理器组
func NewHandlerGroup(
	userService *service.UserService,
	roleService *service.RoleService,
	menuService *service.MenuService,
	operationLogService *service.OperationLogService,
	permissionService *service.PermissionService,
	roleMenuService *service.RoleMenuService,
) *HandlerGroup {
	return &HandlerGroup{
		UserHandler:         NewUserHandler(userService),
		RoleHandler:         NewRoleHandler(roleService),
		MenuHandler:         NewMenuHandler(menuService),
		OperationLogHandler: NewOperationLogHandler(operationLogService),
		PermissionHandler:   NewPermissionHandler(permissionService),
		RoleMenuHandler:     NewRoleMenuHandler(roleMenuService),
	}
}
