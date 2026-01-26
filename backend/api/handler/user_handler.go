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
