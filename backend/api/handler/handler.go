package handler

import (
	assHandler "backend/api/handler/asset_management"
	sysHandler "backend/api/handler/system_manager"
	assService "backend/internal/service/asset_management"
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

type HostHandler struct {
	HostGroupHandler *assHandler.HostGroupHandler
	HostHandler      *assHandler.HostHandler
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

func NewHostHandler(
	hostService *assService.HostService,
	hostGroupService *assService.HostGroupService,
	hostMetricService *assService.HostMetricService,
) *HostHandler {
	return &HostHandler{
		HostHandler:      assHandler.NewHostHandler(hostService, hostMetricService),
		HostGroupHandler: assHandler.NewHostGroupHandler(hostGroupService),
	}
}
