package operationtool

import (
	assService "backend/internal/service/asset_management"
	"backend/pkg/response"
	"backend/pkg/ssh"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

// SSHCommandResult SSH 命令执行结果
type SSHCommandResult struct {
	SuccessCount  int    // 成功执行的命令数
	TotalCommands int    // 总命令数
	Output        string // 标准输出
	ErrorOutput   string // 错误输出
	Success       bool   // 是否全部成功
}

// ExecuteSSHCommands 通过 SSH 执行多个命令
// 参数:
//   - sshManager: SSH 管理器
//   - host: 主机 IP 地址
//   - port: 端口号 (uint16)
//   - username: 用户名
//   - password: 密码
//   - commands: 要执行的命令切片
//   - cmdWaitTime: 每个命令等待时间 (秒)
//
// 返回:
//   - *SSHCommandResult: 执行结果
//   - error: 错误信息
func ExecuteSSHCommands(
	sshManager *ssh.SSHManager,
	host string,
	port uint16,
	username string,
	password string,
	commands []string,
	cmdWaitTime time.Duration,
) (*SSHCommandResult, error) {
	// 建立 SSH 连接
	portStr := strconv.FormatUint(uint64(port), 10)
	client, err := sshManager.CreateSSHClient(host, portStr, username, password)
	if err != nil {
		return nil, fmt.Errorf("SSH 连接失败：%w", err)
	}
	defer client.Close()

	// 创建会话
	session, err := sshManager.CreateSession(client)
	if err != nil {
		return nil, fmt.Errorf("创建 SSH 会话失败：%w", err)
	}
	defer sshManager.CloseSession(session)

	// 启动 Shell
	if err := sshManager.StartShell(session); err != nil {
		return nil, fmt.Errorf("启动 Shell 失败：%w", err)
	}

	// 初始化输出缓冲区
	var output bytes.Buffer
	var errorOutput bytes.Buffer
	successCount := 0
	totalCommands := len(commands)

	// 使用 channel 同步读取输出
	stdoutChan := make(chan []byte, 100)
	stderrChan := make(chan []byte, 100)
	doneChan := make(chan struct{})

	// 启动读取 goroutine
	go func() {
		buf := make([]byte, 4086)
		for {
			n, err := session.StdoutPipe.Read(buf)
			if n > 0 {
				data := make([]byte, n)
				copy(data, buf[:n])
				stdoutChan <- data
			}
			if err != nil {
				if err != io.EOF {
					zap.L().Warn("读取 stdout 错误", zap.Error(err))
				}
				break
			}
		}
	}()

	go func() {
		buf := make([]byte, 4086)
		for {
			n, err := session.StderrPipe.Read(buf)
			if n > 0 {
				data := make([]byte, n)
				copy(data, buf[:n])
				stderrChan <- data
			}
			if err != nil {
				if err != io.EOF {
					zap.L().Warn("读取 stderr 错误", zap.Error(err))
				}
				break
			}
		}
	}()

	// 收集输出
	go func() {
		for {
			select {
			case data := <-stdoutChan:
				output.Write(data)
			case data := <-stderrChan:
				errorOutput.Write(data)
			case <-doneChan:
				return
			}
		}
	}()

	// 执行每个命令
	for i, cmd := range commands {
		zap.L().Info("执行命令",
			zap.Int("index", i+1),
			zap.Int("total", totalCommands),
			zap.String("command", cmd))

		if _, err := session.StdinPipe.Write([]byte(cmd + "\n")); err != nil {
			close(doneChan)
			return nil, fmt.Errorf("写入命令失败：%w", err)
		}

		// 等待命令执行
		time.Sleep(cmdWaitTime)
		successCount++
	}

	// 执行 echo $? 获取最后一个命令的退出码
	if _, err := session.StdinPipe.Write([]byte("echo $?\n")); err != nil {
		close(doneChan)
		return nil, fmt.Errorf("写入状态检查命令失败：%w", err)
	}
	time.Sleep(1 * time.Second)

	// 等待输出收集完成
	close(doneChan)
	time.Sleep(500 * time.Millisecond)

	// 判断执行结果
	errorStr := errorOutput.String()
	if errorStr != "" {
		zap.L().Error("命令执行出错",
			zap.String("error_output", errorStr),
			zap.Int("success_count", successCount),
			zap.Int("total_commands", totalCommands))
	}

	if successCount != totalCommands {
		return nil, fmt.Errorf("部分命令执行失败：成功 %d/%d", successCount, totalCommands)
	}

	return &SSHCommandResult{
		SuccessCount:  successCount,
		TotalCommands: totalCommands,
		Output:        output.String(),
		ErrorOutput:   errorStr,
		Success:       errorStr == "",
	}, nil
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

	// 定义要执行的命令序列
	commands := []string{
		"ls -la",
		"pwd",
		"whoami",
		// "wget https://github.com/prometheus/node_exporter/releases/download/v1.7.0/node_exporter-1.7.0.linux-amd64.tar.gz",
		// "tar xvfz node_exporter-1.7.0.linux-amd64.tar.gz",
		// "sudo mv node_exporter-1.7.0.linux-amd64/node_exporter /usr/local/bin/",
		// "rm -rf node_exporter-1.7.0.linux-amd64*",
	}

	// 执行 SSH 命令
	result, err := ExecuteSSHCommands(
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
		fmt.Sprintf("命令执行成功 (共执行 %d 条命令)", result.SuccessCount),
		gin.H{
			"host_id":        hostID,
			"credential_id":  credentialID,
			"status":         "completed",
			"success_count":  result.SuccessCount,
			"total_commands": result.TotalCommands,
			"output":         result.Output,
		})
}
