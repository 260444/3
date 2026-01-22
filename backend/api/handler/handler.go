package handler

import (
	"backend/internal/model"
	"backend/internal/service"
	"backend/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email"`
		Nickname string `json:"nickname"`
		RoleID   uint   `json:"role_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserService.CreateUser(req.Username, req.Password, req.Email, req.Nickname, req.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "用户创建成功",
		"data":    user,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 生成token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成token失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"data": gin.H{
			"user":  user,
			"token": token,
		},
	})
}

func (h *UserHandler) Logout(c *gin.Context) {
	// 从中间件获取用户信息
	userID, _ := c.Get("userID")
	username, _ := c.Get("username")

	// TODO: 如果使用 Redis 存储 token，可以在这里将 token 加入黑名单
	// 目前 JWT 是无状态的，客户端删除 token 即可实现退出登录

	c.JSON(http.StatusOK, gin.H{
		"message": "退出登录成功",
		"data": gin.H{
			"user_id":  userID,
			"username": username,
		},
	})
}

func (h *UserHandler) GetUserInfo(c *gin.Context) {
	// 从中间件获取用户信息
	userID, _ := c.Get("userID")
	user, err := h.UserService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户信息失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    user,
	})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	users, total, err := h.UserService.GetUsers(pageSize, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data": gin.H{
			"list":      users,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))

	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.UserService.UpdateUser(uint(userID), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}

func (h *UserHandler) UpdateUserStatus(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	status, _ := strconv.Atoi(c.PostForm("status"))

	err := h.UserService.UpdateUserStatus(uint(userID), status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新状态成功",
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))

	err := h.UserService.DeleteUser(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.UserService.ChangePassword(userID.(uint), req.OldPassword, req.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "密码修改成功",
	})
}

func (h *UserHandler) ResetPassword(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		NewPassword string `json:"new_password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.UserService.ResetPassword(uint(userID), req.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "密码重置成功",
	})
}

type RoleHandler struct {
	RoleService *service.RoleService
}

func NewRoleHandler(roleService *service.RoleService) *RoleHandler {
	return &RoleHandler{
		RoleService: roleService,
	}
}

func (h *RoleHandler) CreateRole(c *gin.Context) {
	var req model.Role
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.RoleService.CreateRole(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "角色创建成功",
		"data":    req,
	})
}

func (h *RoleHandler) GetRole(c *gin.Context) {
	roleID, _ := strconv.Atoi(c.Param("id"))

	role, err := h.RoleService.GetRoleByID(uint(roleID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    role,
	})
}

func (h *RoleHandler) GetRoles(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	roles, total, err := h.RoleService.GetRoles(pageSize, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data": gin.H{
			"list":      roles,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func (h *RoleHandler) UpdateRole(c *gin.Context) {
	roleID, _ := strconv.Atoi(c.Param("id"))

	var req model.Role
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = uint(roleID)
	err := h.RoleService.UpdateRole(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}

func (h *RoleHandler) DeleteRole(c *gin.Context) {
	roleID, _ := strconv.Atoi(c.Param("id"))

	err := h.RoleService.DeleteRole(uint(roleID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

type MenuHandler struct {
	MenuService *service.MenuService
}

func NewMenuHandler(menuService *service.MenuService) *MenuHandler {
	return &MenuHandler{
		MenuService: menuService,
	}
}

func (h *MenuHandler) CreateMenu(c *gin.Context) {
	var req model.Menu
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.MenuService.CreateMenu(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "菜单创建成功",
		"data":    req,
	})
}

func (h *MenuHandler) GetMenu(c *gin.Context) {
	menuID, _ := strconv.Atoi(c.Param("id"))

	menu, err := h.MenuService.GetMenuByID(uint(menuID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    menu,
	})
}

//func (h *MenuHandler) GetMenuTree(c *gin.Context) {
//	var parentID *uint
//	parentIDStr := c.Query("parent_id")
//	if parentIDStr != "" {
//		id, err := strconv.Atoi(parentIDStr)
//		if err == nil {
//			parentID = &[]uint{uint(id)}[0]
//		}
//	}
//
//	menus, err := h.MenuService.GetMenuTree(parentID)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"message": "获取成功",
//		"data":    menus,
//	})
//}

func (h *MenuHandler) GetUserMenus(c *gin.Context) {
	//获取用户id
	userID, _ := c.Get("userID")

	//获取当前用户的菜单
	menus, err := h.MenuService.GetUserMenus(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    menus,
	})
}

func (h *MenuHandler) GetAllMenus(c *gin.Context) {
	menus, err := h.MenuService.GetAllMenus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    menus,
	})
}

func uintPtr(value uint) *uint {
	return &value
}
func (h *MenuHandler) UpdateMenu(c *gin.Context) {
	menuID, _ := strconv.Atoi(c.Param("id"))

	var req model.Menu
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = uint(menuID)
	if req.ParentID == nil {
		req.ParentID = uintPtr(0)
	}

	err := h.MenuService.UpdateMenu(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}

func (h *MenuHandler) DeleteMenu(c *gin.Context) {
	menuID, _ := strconv.Atoi(c.Param("id"))

	err := h.MenuService.DeleteMenu(uint(menuID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

type OperationLogHandler struct {
	OperationLogService *service.OperationLogService
}

func NewOperationLogHandler(operationLogService *service.OperationLogService) *OperationLogHandler {
	return &OperationLogHandler{
		OperationLogService: operationLogService,
	}
}

func (h *OperationLogHandler) GetOperationLogs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	logs, total, err := h.OperationLogService.GetOperationLogs(pageSize, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data": gin.H{
			"list":      logs,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func (h *OperationLogHandler) DeleteOperationLog(c *gin.Context) {
	logID, _ := strconv.Atoi(c.Param("id"))

	err := h.OperationLogService.DeleteOperationLog(uint(logID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

// PermissionHandler 权限管理处理器
type PermissionHandler struct {
	PermissionService *service.PermissionService
}

func NewPermissionHandler(permissionService *service.PermissionService) *PermissionHandler {
	return &PermissionHandler{
		PermissionService: permissionService,
	}
}

type RoleMeanHandler struct {
	RoleMenuService *service.RoleMenuService
}

func NewRoleMeanHandler(roleMeanHandler *service.RoleMenuService) *RoleMeanHandler {
	return &RoleMeanHandler{
		RoleMenuService: roleMeanHandler,
	}
}

// AssignMenuToRole 为角色分配菜单权限
func (h *RoleMeanHandler) AssignMenuToRole(c *gin.Context) {
	roleID, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		MenuIDs []uint `json:"menu_ids" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.RoleMenuService.AssignMenuToRole(uint(roleID), req.MenuIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "菜单权限分配成功",
	})
}

// GetRoleMenus 获取角色的菜单权限
func (h *RoleMeanHandler) GetRoleMenus(c *gin.Context) {
	roleID, _ := strconv.Atoi(c.Param("id"))

	menus, err := h.RoleMenuService.GetUserMenusByID(uint(roleID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    menus,
	})
}

// RemoveMenuFromRole 移除角色的菜单权限
func (h *RoleMeanHandler) RemoveMenuFromRole(c *gin.Context) {
	roleID, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		MenuIDs []uint `json:"menu_ids" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.RoleMenuService.RemoveMenuFromRole(uint(roleID), req.MenuIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "菜单权限移除成功",
	})
}

// AddPolicy 添加Casbin策略
func (h *PermissionHandler) AddPolicy(c *gin.Context) {
	roleID, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		Path   string `json:"path" binding:"required"`
		Method string `json:"method" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.PermissionService.AddPolicy(uint(roleID), req.Path, req.Method)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "策略添加成功",
	})
}

// RemovePolicy 移除Casbin策略
func (h *PermissionHandler) RemovePolicy(c *gin.Context) {
	roleID, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		Path   string `json:"path" binding:"required"`
		Method string `json:"method" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.PermissionService.RemovePolicy(uint(roleID), req.Path, req.Method)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "策略移除成功",
	})
}

// GetPolicies 获取角色的所有Casbin策略
func (h *PermissionHandler) GetPolicies(c *gin.Context) {
	roleID, _ := strconv.Atoi(c.Param("id"))

	policies, err := h.PermissionService.GetPolicies(uint(roleID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    policies,
	})
}

// GetAllPolicies 获取所有Casbin策略
func (h *PermissionHandler) GetAllPolicies(c *gin.Context) {
	policies, err := h.PermissionService.GetAllPolicies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    policies,
	})
}

// CreatePermission 创建权限
func (h *PermissionHandler) CreatePermission(c *gin.Context) {
	var req model.Permission
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.PermissionService.CreatePermission(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "权限创建成功",
		"data":    req,
	})
}

// GetPermission 获取权限详情
func (h *PermissionHandler) GetPermission(c *gin.Context) {
	permissionID, _ := strconv.Atoi(c.Param("id"))

	permission, err := h.PermissionService.GetPermissionByID(uint(permissionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    permission,
	})
}

// GetPermissions 获取权限列表
func (h *PermissionHandler) GetPermissions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	permissions, err := h.PermissionService.GetPermissions(pageSize, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total, err := h.PermissionService.GetPermissionTotal()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data": gin.H{
			"list":      permissions,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

// UpdatePermission 更新权限
func (h *PermissionHandler) UpdatePermission(c *gin.Context) {
	permissionID, _ := strconv.Atoi(c.Param("id"))

	var req model.Permission
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = uint(permissionID)
	err := h.PermissionService.UpdatePermission(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}

// UpdatePermissionStatus 更新权限状态
func (h *PermissionHandler) UpdatePermissionStatus(c *gin.Context) {
	permissionID, _ := strconv.Atoi(c.Param("id"))
	status, _ := strconv.Atoi(c.PostForm("status"))

	err := h.PermissionService.UpdatePermissionStatus(uint(permissionID), int8(status))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新状态成功",
	})
}

// DeletePermission 删除权限
func (h *PermissionHandler) DeletePermission(c *gin.Context) {
	permissionID, _ := strconv.Atoi(c.Param("id"))

	err := h.PermissionService.DeletePermission(uint(permissionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}
