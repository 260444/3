package system_manager

import (
	"net/http"
	"strconv"

	sysService "backend/internal/service/system_manager"

	"github.com/gin-gonic/gin"
)

type RoleMenuHandler struct {
	RoleMenuService *sysService.RoleMenuService
}

func NewRoleMenuHandler(roleMenuService *sysService.RoleMenuService) *RoleMenuHandler {
	return &RoleMenuHandler{
		RoleMenuService: roleMenuService,
	}
}

// AssignMenuToRole 为角色分配菜单权限
func (h *RoleMenuHandler) AssignMenuToRole(c *gin.Context) {
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
func (h *RoleMenuHandler) GetRoleMenus(c *gin.Context) {
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
func (h *RoleMenuHandler) RemoveMenuFromRole(c *gin.Context) {
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
