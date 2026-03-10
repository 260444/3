package operationtool

import (
	hostModel "backend/internal/model/asset_management"
	operModel "backend/internal/model/operationTool"
	"backend/pkg/logger"
	"backend/pkg/response"
	"backend/pkg/ssh"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ExecuteCommandHandler 批量在多台主机上执行命令
// @Summary 批量执行命令
// @Description 使用指定的凭据在多台主机上并行执行命令
// @Tags 运维工具
// @Accept json
// @Produce json
// @Param credential_id path int true "凭据 ID"
// @Param body body ExecuteCommandRequest true "执行请求"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/execute-command/{credential_id} [post]
// @Security Bearer

func (h *OperationToolsHandler) ExecuteCommandOnHosts(c *gin.Context) {
	// 解析路径参数
	credentialID, err := strconv.ParseUint(c.Param("credential_id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "credential_id", "无效的凭据 ID")
		return
	}

	// 解析请求体
	var req operModel.ExecuteCommandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorWithCode(c, 400, "请求参数错误："+err.Error())
		return
	}

	// 验证主机列表
	if len(req.HostIDs) == 0 {
		response.ValidationError(c, "host_ids", "至少需要指定一台主机")
		return
	}

	// 验证命令
	if strings.TrimSpace(req.Commands) == "" {
		response.ValidationError(c, "commands", "命令不能为空")
		return
	}

	// 获取凭据信息
	credential, err := h.CredentialService.GetCredentialByID(uint(credentialID))
	if err != nil {
		response.Error(c, fmt.Errorf("获取凭据失败：%w", err))
		return
	}

	// 获取所有目标主机信息
	hosts := make([]*hostModel.Host, 0, len(req.HostIDs))
	for _, hostID := range req.HostIDs {
		host, err := h.HostService.GetHostByID(hostID)
		if err != nil {
			logger.Logger.Warn("获取主机信息失败",
				zap.Uint("host_id", hostID),
				zap.Error(err))

			continue
		}
		hosts = append(hosts, host)
	}

	if len(hosts) == 0 {
		response.Error(c, fmt.Errorf("未找到任何有效的主机"))
		return
	}

	// 并发执行命令
	var (
		wg           sync.WaitGroup
		mu           sync.Mutex
		results      = make([]operModel.HostCommandResult, 0, len(hosts))
		successCount = 0
		failedCount  = 0
	)

	for _, host := range hosts {
		wg.Add(1)
		go func(h *hostModel.Host) {
			defer wg.Done()

			startTime := time.Now()
			result := operModel.HostCommandResult{
				HostID:    h.ID,
				Hostname:  h.Hostname,
				IPAddress: h.IPAddress,
			}

			sshManager := ssh.NewSSHManager()
			// 执行命令
			execResult, err := sshManager.ExecuteSingleCommand(
				h.IPAddress,
				credential.Username,
				credential.Password,
				req.Commands,
				h.Port,
				2*time.Second,
			)

			result.ExecuteTime = time.Since(startTime).Milliseconds()

			if err != nil {
				result.Success = false
				result.Error = err.Error()
				failedCount++
				logger.Logger.Error("主机命令执行失败",
					zap.String("hostname", h.Hostname),
					zap.String("ip", h.IPAddress),
					zap.Error(err))
			} else {
				result.Success = true
				result.Output = execResult.Output
				successCount++
				logger.Logger.Info("主机命令执行成功",
					zap.String("hostname", h.Hostname),
					zap.String("ip", h.IPAddress),
					zap.Int64("耗时_ms", result.ExecuteTime))
			}

			// 安全地添加到结果切片
			mu.Lock()
			results = append(results, result)
			mu.Unlock()
		}(host)
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	// 返回结果
	response.Success(c, gin.H{
		"total":   len(hosts),
		"success": successCount,
		"failed":  failedCount,
		"results": results,
	})
}
