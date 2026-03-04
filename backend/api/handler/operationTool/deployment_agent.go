package operationtool

import (
	assService "backend/internal/service/asset_management"
	"backend/pkg/response"
	"backend/pkg/ssh"
	"bytes"
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

// DeploymentAgent 部署node_exporter
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

	// 建立 SSH 连接
	portStr := strconv.FormatUint(uint64(host.Port), 10)
	client, err := h.SSHManager.CreateSSHClient(host.IPAddress, portStr, credential.Username, credential.Password)
	if err != nil {
		response.Error(c, fmt.Errorf("SSH 连接失败：%w", err))
		return
	}
	defer client.Close()

	// // 创建会话
	session, err := h.SSHManager.CreateSession(client)
	if err != nil {
		response.Error(c, fmt.Errorf("创建 SSH 会话失败：%w", err))
		return
	}
	defer h.SSHManager.CloseSession(session)

	// 启动 Shell
	if err := h.SSHManager.StartShell(session); err != nil {
		response.Error(c, fmt.Errorf("启动 Shell 失败：%w", err))
		return
	}

	// 部署 node_exporter 的命令序列
	commands := []string{
		"ls -la",
		"pwd",
		"whoami",
		// "wget https://github.com/prometheus/node_exporter/releases/download/v1.7.0/node_exporter-1.7.0.linux-amd64.tar.gz",
		// "tar xvfz node_exporter-1.7.0.linux-amd64.tar.gz",
		// "sudo mv node_exporter-1.7.0.linux-amd64/node_exporter /usr/local/bin/",
		// "rm -rf node_exporter-1.7.0.linux-amd64*",
	}

	// 执行命令并检查状态
	var output bytes.Buffer
	var errorOutput bytes.Buffer
	successCount := 0
	totalCommands := len(commands)

	// 读取标准输出和错误输出的 goroutine
	go func() {
		buf := make([]byte, 4086)
		for {
			n, err := session.StdoutPipe.Read(buf)
			if n > 0 {
				output.Write(buf[:n])
			}
			if err != nil {
				break
			}
		}
	}()

	go func() {
		buf := make([]byte, 4086)
		for {
			n, err := session.StderrPipe.Read(buf)
			if n > 0 {
				errorOutput.Write(buf[:n])
			}
			if err != nil {
				break
			}
		}
	}()

	// 执行每个命令并等待完成
	for i, cmd := range commands {
		fmt.Printf("正在执行命令 %d/%d: %s\n", i+1, totalCommands, cmd)

		if _, err := session.StdinPipe.Write([]byte(cmd + "\n")); err != nil {
			response.Error(c, fmt.Errorf("写入命令失败：%w", err))
			return
		}

		// 等待命令执行（可以通过读取特定标记来判断）
		time.Sleep(2 * time.Second)
		successCount++
	}

	// 执行 echo $? 获取最后一个命令的退出码
	fmt.Println("检查命令执行状态...")
	if _, err := session.StdinPipe.Write([]byte("echo $?\n")); err != nil {
		response.Error(c, fmt.Errorf("写入状态检查命令失败：%w", err))
		return
	}
	time.Sleep(1 * time.Second)

	// 读取所有输出
	finalOutput := output.String()
	errorStr := errorOutput.String()

	if errorStr != "" {
		fmt.Println("========== 错误输出 ==========")
		fmt.Println(errorStr)
	}

	// 判断执行是否成功
	if errorStr != "" {
		response.Error(c, fmt.Errorf("命令执行过程中出现错误：\n%s", errorStr))
		return
	}

	if successCount != totalCommands {
		response.Error(c, fmt.Errorf("部分命令执行失败：成功 %d/%d", successCount, totalCommands))
		return
	}

	// 响应成功
	response.SuccessWithMessage(c, fmt.Sprintf("命令执行成功 (共执行 %d 条命令)", successCount), gin.H{
		"host_id":        hostID,
		"credential_id":  credentialID,
		"status":         "completed",
		"success_count":  successCount,
		"total_commands": totalCommands,
		"output":         finalOutput,
	})
}
