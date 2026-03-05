package operationtool

import (
	assService "backend/internal/service/asset_management"
	"backend/pkg/response"
	"backend/pkg/ssh"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DeploymentAgentHandler struct {
	SSHManager        *ssh.SSHManager
	HostService       *assService.HostService
	CredentialService *assService.CredentialService
}

func NewDeploymentAgentHandler(HostService *assService.HostService, CredentialService *assService.CredentialService) *DeploymentAgentHandler {
	return &DeploymentAgentHandler{
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
func (h *DeploymentAgentHandler) DeploymentAgent(c *gin.Context) {
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
	result, err := ssh.UploadFile(h.SSHManager, "node_exporter-1.10.2.linux-amd64.tar.gz", host.IPAddress, credential.Username, credential.Password, host.Port)
	if err != nil {
		response.Error(c, err)
		return
	}
	fmt.Println("上传文件结果:", result.Output)
	// 定义要执行的命令序列
	commands := []string{
		"[ -e /usr/local/node_exporter-1.10.2.linux-amd64 ] || tar xvf /tmp/node_exporter-1.10.2.linux-amd64.tar.gz -C /usr/local",
		"[ -L /usr/local/node_exporter ] || (cd /usr/local && ln -s node_exporter-1.10.2.linux-amd64 node_exporter)",
	}

	// 执行 SSH 命令
	err = ssh.ExecuteSSHCommands(
		h.SSHManager,
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
		fmt.Sprintf("命令执行成功"),
		gin.H{
			"host_id":       hostID,
			"credential_id": credentialID,
			"status":        "completed",
		})
}
