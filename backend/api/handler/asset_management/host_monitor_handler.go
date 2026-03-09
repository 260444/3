package asset_management

import (
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HostMonitorHandler struct {
	service *service.HostMonitorService
}

func NewHostMonitorHandler(service *service.HostMonitorService) *HostMonitorHandler {
	return &HostMonitorHandler{service: service}
}

// 获取所有主机监控指标
func (h *HostMonitorHandler) GetHostsMetrics(c *gin.Context) {
	ctx := c.Request.Context()

	metrics, err := h.service.GetAllHostsMetrics(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    metrics,
	})
}

// 获取单个主机监控指标
func (h *HostMonitorHandler) GetHostMetrics(c *gin.Context) {
	// id := c.Param("id")
	// 实现获取单个主机指标的 lógica
}
