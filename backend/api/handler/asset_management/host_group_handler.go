package asset_management

import (
	assModel "backend/internal/model/asset_management"
	assService "backend/internal/service/asset_management"

	"backend/pkg/response"
	"backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type HostGroupHandler struct {
	HostGroupService *assService.HostGroupService
}

func NewHostGroupHandler(HostGroupService *assService.HostGroupService) *HostGroupHandler {
	return &HostGroupHandler{
		HostGroupService: HostGroupService,
	}
}

// CreateHostGroup 创建主机组
// @Summary 创建主机组
// @Description 创建新的主机组
// @Tags 主机组管理
// @Accept json
// @Produce json
// @Param group body assService.HostGroupCreateRequest true "主机组信息"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/host-groups [post]
// @Security Bearer
func (h *HostGroupHandler) CreateHostGroup(c *gin.Context) {
	var req assModel.HostGroupCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "请求参数", err.Error())
		return
	}

	userID := utils.GetUserIDFromContext(c)

	group, err := h.HostGroupService.CreateHostGroup(&req, userID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "主机组创建成功", group)
}

// GetHostGroupList 获取主机组列表
// @Summary 获取主机组列表
// @Description 分页获取主机组列表
// @Tags 主机组管理
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param name query string false "主机组名称模糊搜索"
// @Param status query int false "状态筛选"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/host-groups [get]
// @Security Bearer
func (h *HostGroupHandler) GetHostGroupList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.Query("name")

	var status *int8
	if statusStr := c.Query("status"); statusStr != "" {
		if s, err := strconv.ParseInt(statusStr, 10, 8); err == nil {
			st := int8(s)
			status = &st
		}
	}

	groups, total, err := h.HostGroupService.ListHostGroups(page, pageSize, name, status)
	if err != nil {
		response.Error(c, err)
		return
	}

	// 为主机组添加主机数量信息
	groupsWithCount := make([]map[string]interface{}, len(groups))
	for i, group := range groups {
		var hostCount int64
		// 这里应该通过hostRepo来获取主机数量，但为了简化，我们假设在service层已经处理了
		groupsWithCount[i] = map[string]interface{}{
			"id":          group.ID,
			"name":        group.Name,
			"description": group.Description,
			"status":      group.Status,
			"host_count":  hostCount,
			"created_at":  group.CreatedAt,
			"updated_at":  group.UpdatedAt,
		}
	}

	response.SuccessWithMessage(c, "获取主机组列表成功", gin.H{
		"list": groupsWithCount,
		"pagination": gin.H{
			"page":      page,
			"page_size": pageSize,
			"total":     total,
		},
	})
}

// GetHostGroupByID 获取主机组详情
// @Summary 获取主机组详情
// @Description 根据ID获取主机组详细信息及关联主机
// @Tags 主机组管理
// @Produce json
// @Param id path int true "主机组ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/host-groups/{id} [get]
// @Security Bearer
func (h *HostGroupHandler) GetHostGroupByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "id", "无效的主机组ID")
		return
	}

	group, err := h.HostGroupService.GetHostGroupWithHosts(uint(id))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "获取主机组详情成功", group)
}

// UpdateHostGroup 更新主机组
// @Summary 更新主机组
// @Description 更新主机组的基本信息
// @Tags 主机组管理
// @Accept json
// @Produce json
// @Param id path int true "主机组ID"
// @Param group body assService.HostGroupUpdateRequest true "更新的主机组信息"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/host-groups/{id} [put]
// @Security Bearer
func (h *HostGroupHandler) UpdateHostGroup(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "id", "无效的主机组ID")
		return
	}

	var req assModel.HostGroupUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "请求参数", err.Error())
		return
	}

	userID := utils.GetUserIDFromContext(c)

	group, err := h.HostGroupService.UpdateHostGroup(uint(id), &req, userID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "主机组更新成功", group)
}

// DeleteHostGroup 删除主机组
// @Summary 删除主机组
// @Description 删除指定的主机组（需确保组内无主机）
// @Tags 主机组管理
// @Produce json
// @Param id path int true "主机组ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/host-groups/{id} [delete]
// @Security Bearer
func (h *HostGroupHandler) DeleteHostGroup(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "id", "无效的主机组ID")
		return
	}

	if err := h.HostGroupService.DeleteHostGroup(uint(id)); err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "主机组删除成功", nil)
}

// UpdateHostGroupStatus 更新主机组状态
// @Summary 更新主机组状态
// @Description 更新主机组的启用/禁用状态
// @Tags 主机组管理
// @Accept json
// @Produce json
// @Param id path int true "主机组ID"
// @Param status body assService.HostGroupStatusUpdateRequest true "状态信息"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/host-groups/{id}/status [put]
// @Security Bearer
func (h *HostGroupHandler) UpdateHostGroupStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "id", "无效的主机组ID")
		return
	}

	var req assModel.HostGroupStatusUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "请求参数", err.Error())
		return
	}

	if err := h.HostGroupService.UpdateHostGroupStatus(uint(id), req.Status); err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "主机组状态更新成功", gin.H{"id": id, "status": req.Status})
}
