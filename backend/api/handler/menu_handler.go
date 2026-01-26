package handler

import (
	"backend/internal/model"
	"backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
