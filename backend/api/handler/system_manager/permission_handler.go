package system_manager

import (
	sysModel "backend/internal/model/system_manager"
	sysService "backend/internal/service/system_manager"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// PermissionHandler 权限管理处理器
type PermissionHandler struct {
	PermissionService *sysService.PermissionService
}

func NewPermissionHandler(permissionService *sysService.PermissionService) *PermissionHandler {
	return &PermissionHandler{
		PermissionService: permissionService,
	}
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

	type CasbinPolicy struct {
		Sub string `json:"sub"`
		Obj string `json:"obj"`
		Act string `json:"act"`
	}

	var result []CasbinPolicy
	for _, p := range policies {
		if len(p) >= 3 {
			result = append(result, CasbinPolicy{
				Sub: p[0],
				Obj: p[1],
				Act: p[2],
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    result,
	})
}

// GetAllPolicies 获取所有Casbin策略
func (h *PermissionHandler) GetAllPolicies(c *gin.Context) {
	policies, err := h.PermissionService.GetAllPolicies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type CasbinPolicy struct {
		Sub string `json:"sub"`
		Obj string `json:"obj"`
		Act string `json:"act"`
	}

	var result []CasbinPolicy
	for _, p := range policies {
		if len(p) >= 3 {
			result = append(result, CasbinPolicy{
				Sub: p[0],
				Obj: p[1],
				Act: p[2],
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    result,
	})
}

// CreatePermission 创建权限
func (h *PermissionHandler) CreatePermission(c *gin.Context) {
	var req sysModel.Permission
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
	path := c.Query("path")
	method := c.Query("method")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	permissions, err := h.PermissionService.GetPermissions(pageSize, offset, path, method)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total, err := h.PermissionService.GetPermissionTotal(path, method)
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

	var req sysModel.Permission
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

// GetAllPermissions 获取所有权限
func (h *PermissionHandler) GetAllPermissions(c *gin.Context) {
	path := c.Query("path")
	method := c.Query("method")

	permissions, err := h.PermissionService.GetAllPermissions(path, method)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    permissions,
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

// GetPermissionByPathAndMethod 根据路径和方法获取权限详情
// func (h *PermissionHandler) GetPermissionByPathAndMethod(c *gin.Context) {
// 	path := c.Query("path")
// 	method := c.Query("method")

// 	if path == "" || method == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "path and method are required"})
// 		return
// 	}

// 	permission, err := h.PermissionService.GetPermissionByPathAndMethod(path, method)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "获取成功",
// 		"data":    permission,
// 	})
// }
