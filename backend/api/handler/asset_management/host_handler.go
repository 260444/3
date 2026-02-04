package asset_management

import (
	"d:/ai/3/backend/internal/service/system_manager"
	"d:/ai/3/backend/pkg/response"
	"d:/ai/3/backend/pkg/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type HostHandler struct {
	hostService       *system_manager.HostService
	hostGroupService  *system_manager.HostGroupService
	hostMetricService *system_manager.HostMetricService
}

func NewHostHandler(
	hostService *system_manager.HostService,
	hostGroupService *system_manager.HostGroupService,
	hostMetricService *system_manager.HostMetricService,
) *HostHandler {
	return &HostHandler{
		hostService:       hostService,
		hostGroupService:  hostGroupService,
		hostMetricService: hostMetricService,
	}
}

// CreateHost 创建主机
// @Summary 创建主机
// @Description 创建新的主机记录
// @Tags 主机管理
// @Accept json
// @Produce json
// @Param host body system_manager.HostCreateRequest true "主机信息"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/hosts [post]
// @Security Bearer
func (h *HostHandler) CreateHost(c *gin.Context) {
	var req system_manager.HostCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "请求参数", err.Error())
		return
	}

	// 获取当前用户ID
	userID := utils.GetUserIDFromContext(c)

	host, err := h.hostService.CreateHost(&req, userID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "主机创建成功", host)
}

// GetHostList 获取主机列表
// @Summary 获取主机列表
// @Description 分页获取主机列表，支持多种筛选条件
// @Tags 主机管理
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param hostname query string false "主机名模糊搜索"
// @Param ip_address query string false "IP地址模糊搜索"
// @Param group_id query int false "主机组ID"
// @Param status query int false "主机状态"
// @Param os_type query string false "操作系统类型"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/hosts [get]
// @Security Bearer
func (h *HostHandler) GetHostList(c *gin.Context) {
	// 解析查询参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	hostname := c.Query("hostname")
	ipAddress := c.Query("ip_address")
	osType := c.Query("os_type")

	var groupID *uint
	if groupIDStr := c.Query("group_id"); groupIDStr != "" {
		if id, err := strconv.ParseUint(groupIDStr, 10, 32); err == nil {
			gid := uint(id)
			groupID = &gid
		}
	}

	var status *int8
	if statusStr := c.Query("status"); statusStr != "" {
		if s, err := strconv.ParseInt(statusStr, 10, 8); err == nil {
			st := int8(s)
			status = &st
		}
	}

	hosts, total, err := h.hostService.ListHosts(page, pageSize, hostname, ipAddress, groupID, status, osType)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "获取主机列表成功", gin.H{
		"list": hosts,
		"pagination": gin.H{
			"page":      page,
			"page_size": pageSize,
			"total":     total,
		},
	})
}

// GetHostByID 获取主机详情
// @Summary 获取主机详情
// @Description 根据ID获取主机详细信息
// @Tags 主机管理
// @Produce json
// @Param id path int true "主机ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/hosts/{id} [get]
// @Security Bearer
func (h *HostHandler) GetHostByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "id", "无效的主机ID")
		return
	}

	host, err := h.hostService.GetHostByID(uint(id))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "获取主机详情成功", host)
}

// UpdateHost 更新主机信息
// @Summary 更新主机信息
// @Description 更新主机的基本信息
// @Tags 主机管理
// @Accept json
// @Produce json
// @Param id path int true "主机ID"
// @Param host body system_manager.HostUpdateRequest true "更新的主机信息"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/hosts/{id} [put]
// @Security Bearer
func (h *HostHandler) UpdateHost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "id", "无效的主机ID")
		return
	}

	var req system_manager.HostUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "请求参数", err.Error())
		return
	}

	userID := utils.GetUserIDFromContext(c)

	host, err := h.hostService.UpdateHost(uint(id), &req, userID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "主机更新成功", host)
}

// DeleteHost 删除主机
// @Summary 删除主机
// @Description 删除指定的主机记录
// @Tags 主机管理
// @Produce json
// @Param id path int true "主机ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/hosts/{id} [delete]
// @Security Bearer
func (h *HostHandler) DeleteHost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "id", "无效的主机ID")
		return
	}

	if err := h.hostService.DeleteHost(uint(id)); err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "主机删除成功", nil)
}

// BatchDeleteHosts 批量删除主机
// @Summary 批量删除主机
// @Description 批量删除多个主机记录
// @Tags 主机管理
// @Accept json
// @Produce json
// @Param ids body []uint true "主机ID数组"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/hosts/batch [delete]
// @Security Bearer
func (h *HostHandler) BatchDeleteHosts(c *gin.Context) {
	var req struct {
		IDs []uint `json:"ids" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "ids", "ids参数不能为空")
		return
	}

	affected, err := h.hostService.BatchDeleteHosts(req.IDs)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "批量删除主机成功", gin.H{"deleted_count": affected})
}

// UpdateHostStatus 更新主机状态
// @Summary 更新主机状态
// @Description 更新主机的在线/离线/故障状态
// @Tags 主机管理
// @Accept json
// @Produce json
// @Param id path int true "主机ID"
// @Param status body system_manager.HostStatusUpdateRequest true "状态信息"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/hosts/{id}/status [put]
// @Security Bearer
func (h *HostHandler) UpdateHostStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "id", "无效的主机ID")
		return
	}

	var req system_manager.HostStatusUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "请求参数", err.Error())
		return
	}

	if err := h.hostService.UpdateHostStatus(uint(id), req.Status); err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "主机状态更新成功", gin.H{"id": id, "status": req.Status})
}

// UpdateHostMonitoring 更新主机监控状态
// @Summary 更新主机监控状态
// @Description 启用或禁用主机监控
// @Tags 主机管理
// @Accept json
// @Produce json
// @Param id path int true "主机ID"
// @Param monitoring body system_manager.HostMonitoringUpdateRequest true "监控状态"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/hosts/{id}/monitoring [put]
// @Security Bearer
func (h *HostHandler) UpdateHostMonitoring(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "id", "无效的主机ID")
		return
	}

	var req system_manager.HostMonitoringUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "请求参数", err.Error())
		return
	}

	if err := h.hostService.UpdateHostMonitoring(uint(id), req.MonitoringEnabled); err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "监控状态更新成功", gin.H{"id": id, "monitoring_enabled": req.MonitoringEnabled})
}

// GetHostStatistics 获取主机统计信息
// @Summary 获取主机统计信息
// @Description 获取主机的各种统计信息
// @Tags 主机管理
// @Produce json
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/hosts/statistics [get]
// @Security Bearer
func (h *HostHandler) GetHostStatistics(c *gin.Context) {
	stats, err := h.hostService.GetHostStatistics()
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "获取主机统计信息成功", stats)
}

// ReportHostMetrics 上报主机指标
// @Summary 上报主机指标
// @Description 上报主机的监控指标数据
// @Tags 主机监控
// @Accept json
// @Produce json
// @Param metrics body system_manager.HostMetricsRequest true "指标数据"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/host-metrics [post]
// @Security Bearer
func (h *HostHandler) ReportHostMetrics(c *gin.Context) {
	var req system_manager.HostMetricsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "请求参数", err.Error())
		return
	}

	insertedCount, err := h.hostMetricService.ReportHostMetrics(&req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "指标上报成功", gin.H{"inserted_count": insertedCount})
}

// GetHostMetricsHistory 获取主机指标历史
// @Summary 获取主机指标历史
// @Description 获取主机的历史监控指标数据
// @Tags 主机监控
// @Produce json
// @Param host_id query int true "主机ID"
// @Param metric_type query string false "指标类型"
// @Param metric_name query string false "指标名称"
// @Param start_time query string false "开始时间"
// @Param end_time query string false "结束时间"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(50)
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/host-metrics/history [get]
// @Security Bearer
func (h *HostHandler) GetHostMetricsHistory(c *gin.Context) {
	hostID, err := strconv.ParseUint(c.Query("host_id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "host_id", "host_id参数不能为空")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "50"))
	metricType := c.Query("metric_type")
	metricName := c.Query("metric_name")

	var startTime, endTime *time.Time
	if startStr := c.Query("start_time"); startStr != "" {
		if t, err := time.Parse(time.RFC3339, startStr); err == nil {
			startTime = &t
		}
	}
	if endStr := c.Query("end_time"); endStr != "" {
		if t, err := time.Parse(time.RFC3339, endStr); err == nil {
			endTime = &t
		}
	}

	metrics, total, err := h.hostMetricService.GetHostMetricsHistory(
		uint(hostID), metricType, metricName, startTime, endTime, page, pageSize)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "获取指标历史成功", gin.H{
		"list": metrics,
		"pagination": gin.H{
			"page":      page,
			"page_size": pageSize,
			"total":     total,
		},
	})
}

// GetHostLatestMetrics 获取主机最新指标
// @Summary 获取主机最新指标
// @Description 获取主机的最新监控指标数据
// @Tags 主机监控
// @Produce json
// @Param host_id query int true "主机ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/host-metrics/latest [get]
// @Security Bearer
func (h *HostHandler) GetHostLatestMetrics(c *gin.Context) {
	hostID, err := strconv.ParseUint(c.Query("host_id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "host_id", "host_id参数不能为空")
		return
	}

	metrics, err := h.hostMetricService.GetHostLatestMetrics(uint(hostID))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.SuccessWithMessage(c, "获取最新指标成功", metrics)
}
