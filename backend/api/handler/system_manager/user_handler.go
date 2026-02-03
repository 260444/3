package system_manager

import (
	"backend/internal/model/system_manager"
	"backend/pkg/logger"
	"backend/pkg/utils"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	sysService "backend/internal/service/system_manager"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *sysService.UserService
}

func NewUserHandler(userService *sysService.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// CreateUser 创建用户 *
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email"`
		Nickname string `json:"nickname"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Logger.Error("绑定角色失败:", zap.Error(err))
		return
	}

	user, err := h.UserService.CreateUser(req.Username, req.Password, req.Email, req.Nickname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("创建角色失败:", zap.Error(err))
		return
	}

	logger.Logger.Info("创建角色成功:", zap.String("username", user.Username))
	c.JSON(http.StatusOK, gin.H{
		"message": "用户创建成功",
		"data":    user,
	})
}

// AssignRole 为用户分配角色 *
func (h *UserHandler) AssignRole(c *gin.Context) {
	username := c.Param("username")
	var req struct {
		RoleIdent string `json:"role_ident" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.UserService.AddRoleForUser(username, req.RoleIdent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("创建角色失败:", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "角色分配成功",
	})
}

// GetUserRoles 获取用户的角色列表 *
func (h *UserHandler) GetUserRoles(c *gin.Context) {
	username := c.Param("username")

	roles, err := h.UserService.GetUserRoles(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("获取角色失败:", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    roles,
	})
}

// RemoveRole 移除用户的角色 *
func (h *UserHandler) RemoveRole(c *gin.Context) {
	username := c.Param("username")
	var req struct {
		RoleIdent string `json:"role_ident" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.UserService.RemoveRoleForUser(username, req.RoleIdent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("创建角色失败:", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "角色移除成功",
	})
}

// Login 登录 *
func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//这里返回的是带ident的user
	user, err := h.UserService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 生成token
	token, err := utils.GenerateToken(user.ID, user.Username, user.RoleIdent)
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

// Logout 退出登录 *
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

// GetUserInfo 获取用户信息 *
func (h *UserHandler) GetUserInfo(c *gin.Context) {
	// 从中间件获取用户信息
	userID, _ := c.Get("userID")
	user, err := h.UserService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户信息失败"})
		logger.Logger.Error("获取用户信息失败:", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    user,
	})
}

// GetCurrentUser 获取当前用户信息 *
func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	// 从中间件获取用户信息
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未认证"})
		return
	}

	user, err := h.UserService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户信息失败"})
		logger.Logger.Error("获取当前用户信息失败:", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    user,
	})
}

// GetUsers 获取用户列表 *
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
		logger.Logger.Error("获取用户列表失败:", zap.Error(err))
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

// UpdateUser 更新用户信息 *
func (h *UserHandler) UpdateUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))

	var req system_manager.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Logger.Error("绑定用户信息失败:", zap.Error(err))
		return
	}

	err := h.UserService.UpdateUser(uint(userID), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("更新用户信息失败:", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}

// UpdateUserStatus 更新用户状态 *
func (h *UserHandler) UpdateUserStatus(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	status, _ := strconv.Atoi(c.PostForm("status"))

	err := h.UserService.UpdateUserStatus(uint(userID), status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("更新用户状态失败:", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新状态成功",
	})
}

// DeleteUser 删除用户 *
func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))

	err := h.UserService.DeleteUser(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("删除用户失败:", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

// ChangePassword 修改密码 *
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
	logger.Logger.Info("密码修改请求",
		zap.String("旧密码", req.OldPassword),
		zap.String("新密码", req.NewPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("密码修改失败:", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "密码修改成功",
	})
}

// ResetPassword 重置密码 *
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
		logger.Logger.Error("密码重置失败:", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "密码重置成功",
	})
}
