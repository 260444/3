package system_manager

import (
	sysService "backend/internal/service/system_manager"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OperationLogHandler struct {
	OperationLogService *sysService.OperationLogService
}

func NewOperationLogHandler(operationLogService *sysService.OperationLogService) *OperationLogHandler {
	return &OperationLogHandler{
		OperationLogService: operationLogService,
	}
}

func (h *OperationLogHandler) GetOperationLogs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	logs, total, err := h.OperationLogService.GetOperationLogs(pageSize, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data": gin.H{
			"list":      logs,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func (h *OperationLogHandler) DeleteOperationLog(c *gin.Context) {
	logID, _ := strconv.Atoi(c.Param("id"))

	err := h.OperationLogService.DeleteOperationLog(uint(logID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}
