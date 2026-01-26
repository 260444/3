package handler

import (
	"backend/internal/model"
	"backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
