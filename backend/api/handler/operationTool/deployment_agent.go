package operationtool

import (
	assService "backend/internal/service/asset_management"

	"backend/pkg/logger"
	"backend/pkg/response"
	"backend/pkg/ssh"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OperationToolsHandler struct {
	SSHManager        *ssh.SSHManager
	HostService       *assService.HostService
	CredentialService *assService.CredentialService
}

func NewOperationToolsHandler(HostService *assService.HostService, CredentialService *assService.CredentialService) *OperationToolsHandler {
	return &OperationToolsHandler{
		SSHManager:        ssh.NewSSHManager(),
		HostService:       HostService,
		CredentialService: CredentialService,
	}
}

// DeploymentAgent 部署 node_exporter
// @Summary 部署node_exporter
// @Description 通过 SSH 在远程主机上部署node_exporter 监控代理
// @Tags 运维工具
// @Accept json
// @Produce json
// @Param host_id path int true "主机 ID"
// @Param credential_id path int true "凭据 ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/deployment-agent/{host_id}/{credential_id} [post]
// @Security Bearer
func (h *OperationToolsHandler) DeploymentAgentHandler(c *gin.Context) {
	// 解析路径参数
	hostID, err := strconv.ParseUint(c.Param("host_id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "host_id", "无效的主机 ID")
		return
	}

	credentialID, err := strconv.ParseUint(c.Param("credential_id"), 10, 32)
	if err != nil {
		response.ValidationError(c, "credential_id", "无效的凭据 ID")
		return
	}

	// TODO: 根据 hostID 和 credentialID 获取主机和凭据信息
	host, err := h.HostService.GetHostByID(uint(hostID))
	if err != nil {
		response.Error(c, err)
		return
	}

	credential, err := h.CredentialService.GetCredentialByID(uint(credentialID))
	if err != nil {
		response.Error(c, err)
		return
	}

	// 上传文件
	result, err := h.SSHManager.UploadFile("node_export.zip", host.IPAddress, credential.Username, credential.Password, host.Port)
	if err != nil {
		response.Error(c, err)
		return
	}

	logger.Logger.Info("上传文件结果:", zap.String("output", result.Output))
	// 定义要执行的命令序列
	commands := []string{
		// 解压 zip 包
		"[ -f /tmp/node_export.zip ] && unzip -o /tmp/node_export.zip -d /tmp/",

		// 执行部署脚本
		"bash /tmp/deploy_node_exporter.sh &> /tmp/deploy_node_exporter.log",
	}
	// 执行 SSH 命令
	err = h.SSHManager.ExecuteMultipleCommands(
		host.IPAddress,
		host.Port,
		credential.Username,
		credential.Password,
		commands,
		2*time.Second, // 每个命令等待 2 秒
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	// 响应成功
	response.SuccessWithMessage(c,
		"命令执行成功",
		gin.H{
			"host_id":       hostID,
			"credential_id": credentialID,
			"status":        "completed",
		})
}
