package asset_management

import (
	assModel "backend/internal/model/asset_management"
	assService "backend/internal/service/asset_management"
	"backend/pkg/response"
	"backend/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CredentialHandler struct {
	CredentialService *assService.CredentialService
}

func NewCredentialHandler(CredentialService *assService.CredentialService) *CredentialHandler {
	return &CredentialHandler{
		CredentialService: CredentialService,
	}
}

// CreateCredential 创建凭据
// @Summary 创建凭据
// @Description 创建新的凭据记录
// @Tags 凭据管理
// @Accept json
// @Produce json
// @Param credential body assModel.CredentialCreateRequest true "凭据信息"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/credentials [post]
// @Security Bearer
func (h *CredentialHandler) CreateCredential(c *gin.Context) {
	var req assModel.CredentialCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "请求参数", err.Error())
		return
	}

	// 获取当前用户ID
	userID := utils.GetUserIDFromContext(c)

	credential, err := h.CredentialService.CreateCredential(&req, userID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "凭据创建成功", credential)
}

// GetCredentialList 获取凭据列表
// @Summary 获取凭据列表
// @Description 分页获取凭据列表，支持多种筛选条件
// @Tags 凭据管理
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param name query string false "凭据名称模糊搜索"
// @Param username query string false "用户名模糊搜索"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/credentials [get]
// @Security Bearer
func (h *CredentialHandler) GetCredentialList(c *gin.Context) {
	var req assModel.CredentialListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ValidationError(c, "查询参数", err.Error())
		return
	}

	// 设置默认值
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	name := c.Query("name")
	username := c.Query("username")

	credentials, total, err := h.CredentialService.ListCredentials(req.Page, req.PageSize, name, username)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "获取凭据列表成功", gin.H{
		"list": credentials,
		"pagination": gin.H{
			"page":      req.Page,
			"page_size": req.PageSize,
			"total":     total,
		},
	})
}

// GetCredentialByID 获取凭据详情
// @Summary 获取凭据详情
// @Description 根据ID获取凭据详细信息
// @Tags 凭据管理
// @Produce json
// @Param id path int true "凭据ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/credentials/{id} [get]
// @Security Bearer
func (h *CredentialHandler) GetCredentialByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "id", "无效的凭据ID")
		return
	}

	credential, err := h.CredentialService.GetCredentialByID(uint(id))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "获取凭据详情成功", credential)
}

// UpdateCredential 更新凭据信息
// @Summary 更新凭据信息
// @Description 更新凭据的基本信息
// @Tags 凭据管理
// @Accept json
// @Produce json
// @Param id path int true "凭据ID"
// @Param credential body assModel.CredentialUpdateRequest true "更新的凭据信息"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/credentials/{id} [put]
// @Security Bearer
func (h *CredentialHandler) UpdateCredential(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "id", "无效的凭据ID")
		return
	}

	var req assModel.CredentialUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "请求参数", err.Error())
		return
	}

	userID := utils.GetUserIDFromContext(c)

	credential, err := h.CredentialService.UpdateCredential(uint(id), &req, userID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "凭据更新成功", credential)
}

// DeleteCredential 删除凭据
// @Summary 删除凭据
// @Description 删除指定的凭据记录
// @Tags 凭据管理
// @Produce json
// @Param id path int true "凭据ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/credentials/{id} [delete]
// @Security Bearer
func (h *CredentialHandler) DeleteCredential(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "id", "无效的凭据ID")
		return
	}

	if err := h.CredentialService.DeleteCredential(uint(id)); err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "凭据删除成功", nil)
}

// BatchDeleteCredentials 批量删除凭据
// @Summary 批量删除凭据
// @Description 批量删除多个凭据记录
// @Tags 凭据管理
// @Accept json
// @Produce json
// @Param ids body []uint true "凭据ID数组"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/credentials/batch [delete]
// @Security Bearer
func (h *CredentialHandler) BatchDeleteCredentials(c *gin.Context) {
	var req struct {
		IDs []uint `json:"ids" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "ids", "ids参数不能为空")
		return
	}

	affected, err := h.CredentialService.BatchDeleteCredentials(req.IDs)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "批量删除凭据成功", gin.H{"deleted_count": affected})
}

// GetCredentialsByHost 获取主机关联的凭据
// @Summary 获取主机关联的凭据
// @Description 获取指定主机关联的凭据信息
// @Tags 凭据管理
// @Produce json
// @Param host_id query int true "主机ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/credentials/host [get]
// @Security Bearer
func (h *CredentialHandler) GetCredentialsByHost(c *gin.Context) {
	hostID, err := strconv.ParseUint(c.Query("host_id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "host_id", "host_id参数不能为空")
		return
	}

	credentials, err := h.CredentialService.GetCredentialsByHost(uint(hostID))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "获取主机凭据成功", credentials)
}
