package ssh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Session SSH 会话
type Session struct {
	SSHClient  *ssh.Client
	SSHSess    *ssh.Session
	StdinPipe  io.WriteCloser
	StdoutPipe io.Reader
	StderrPipe io.Reader
}

// SSHManager SSH 管理器
type SSHManager struct {
	sessions map[string]*Session
	mu       sync.RWMutex
}

// NewSSHManager 创建 SSH 管理器
func NewSSHManager() *SSHManager {
	return &SSHManager{
		sessions: make(map[string]*Session),
	}
}

// CreateSSHClient 创建 SSH 客户端连接
func (m *SSHManager) CreateSSHClient(host, port, username, password string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	address := fmt.Sprintf("%s:%s", host, port)
	client, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	return client, nil
}

// CreateSession 创建 SSH 会话
func (m *SSHManager) CreateSession(client *ssh.Client) (*Session, error) {
	session, err := client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	// 请求伪终端
	if err := session.RequestPty("xterm", 24, 80, ssh.TerminalModes{}); err != nil {
		session.Close()
		return nil, fmt.Errorf("failed to request pty: %w", err)
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		session.Close()
		return nil, fmt.Errorf("failed to get stdin pipe: %w", err)
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		session.Close()
		return nil, fmt.Errorf("failed to get stdout pipe: %w", err)
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		session.Close()
		return nil, fmt.Errorf("failed to get stderr pipe: %w", err)
	}

	return &Session{
		SSHClient:  client,
		SSHSess:    session,
		StdinPipe:  stdin,
		StdoutPipe: stdout,
		StderrPipe: stderr,
	}, nil
}

// StartShell 启动 Shell
func (m *SSHManager) StartShell(session *Session) error {
	return session.SSHSess.Shell()
}

// CloseSession 关闭会话
func (m *SSHManager) CloseSession(session *Session) {
	if session.SSHSess != nil {
		session.SSHSess.Close()
	}
	if session.SSHClient != nil {
		session.SSHClient.Close()
	}
}

// ResizePty 调整终端大小
func (m *SSHManager) ResizePty(session *Session, rows, cols uint16) error {
	if session.SSHSess == nil {
		return fmt.Errorf("session is nil")
	}
	return session.SSHSess.WindowChange(int(rows), int(cols))
}

// CombineOutput 合并标准输出和标准错误
func CombineOutput(session *Session) io.Reader {
	return io.MultiReader(session.StdoutPipe, session.StderrPipe)
}

// SSHMessage WebSocket 消息结构
type SSHMessage struct {
	Type string `json:"type"` // input, output, resize, error, close
	Data string `json:"data"` // 输入/输出数据
	Cols uint16 `json:"cols"` // 终端列数
	Rows uint16 `json:"rows"` // 终端行数
}

// HandleSSHWebSocket 处理 SSH WebSocket 连接
func (m *SSHManager) HandleSSHWebSocket(conn net.Conn, host, port, username, password string) {
	defer conn.Close()

	// 创建 SSH 客户端
	client, err := m.CreateSSHClient(host, port, username, password)
	if err != nil {
		sendError(conn, fmt.Sprintf("SSH连接失败: %v", err))
		return
	}
	defer client.Close()

	// 创建会话
	session, err := m.CreateSession(client)
	if err != nil {
		sendError(conn, fmt.Sprintf("创建SSH会话失败: %v", err))
		return
	}
	defer m.CloseSession(session)

	// 启动 Shell
	if err := m.StartShell(session); err != nil {
		sendError(conn, fmt.Sprintf("启动Shell失败: %v", err))
		return
	}

	// 读取输出的 goroutine
	outputChan := make(chan []byte, 1024)
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := session.StdoutPipe.Read(buf)
			if err != nil {
				if err != io.EOF {
					outputChan <- []byte(fmt.Sprintf("\r\n连接断开: %v\r\n", err))
				}
				close(outputChan)
				return
			}
			if n > 0 {
				data := make([]byte, n)
				copy(data, buf[:n])
				outputChan <- data
			}
		}
	}()

	// 处理输入输出的循环
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				if err != io.EOF {
					fmt.Printf("读取WebSocket连接错误: %v\n", err)
				}
				return
			}

			var msg SSHMessage
			if err := json.Unmarshal(buf[:n], &msg); err != nil {
				fmt.Printf("解析消息错误: %v\n", err)
				continue
			}

			switch msg.Type {
			case "input":
				if _, err := session.StdinPipe.Write([]byte(msg.Data)); err != nil {
					fmt.Printf("写入输入错误: %v\n", err)
					return
				}
			case "resize":
				if err := m.ResizePty(session, msg.Rows, msg.Cols); err != nil {
					fmt.Printf("调整终端大小错误: %v\n", err)
				}
			case "close":
				return
			}
		}
	}()

	// 发送输出到客户端
	for data := range outputChan {
		if err := sendOutput(conn, data); err != nil {
			fmt.Printf("发送输出错误: %v\n", err)
			return
		}
	}
}

// sendOutput 发送输出消息
func sendOutput(conn net.Conn, data []byte) error {
	msg := SSHMessage{
		Type: "output",
		Data: string(data),
	}
	return sendJSON(conn, msg)
}

// sendError 发送错误消息
func sendError(conn net.Conn, errMsg string) error {
	msg := SSHMessage{
		Type: "error",
		Data: errMsg,
	}
	return sendJSON(conn, msg)
}

// sendJSON 发送 JSON 消息
func sendJSON(conn net.Conn, msg SSHMessage) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = conn.Write(data)
	return err
}

// TestConnection 测试 SSH 连接
func TestConnection(host, port, username, password string) error {
	manager := NewSSHManager()
	client, err := manager.CreateSSHClient(host, port, username, password)
	if err != nil {
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	var buf bytes.Buffer
	session.Stdout = &buf
	if err := session.Run("echo test"); err != nil {
		return err
	}

	return nil
}

// SSHCommandResult SSH 命令执行结果
type SSHCommandResult struct {
	SuccessCount  int    // 成功执行的命令数
	TotalCommands int    // 总命令数
	Output        string // 标准输出
	ErrorOutput   string // 错误输出
	Success       bool   // 是否全部成功
}

// ExecuteSSHCommands 执行多条命令
func ExecuteSSHCommands(
	sshManager *SSHManager,
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

func (m *SSHManager) CreateSFTPClient(client *ssh.Client) (*sftp.Client, error) {
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		return nil, fmt.Errorf("创建 SFTP 客户端失败：%w", err)
	}
	return sftpClient, nil
}

// UploadFile 上传文件
func UploadFile(sshManager *SSHManager, filename, host, username, password string, port uint16) (*SSHCommandResult, error) {
	filepath := fmt.Sprintf("tools/%s", filename)
	// 检查文件是否存在
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil, fmt.Errorf("文件 %s 不存在", filepath)
	}

	// 建立 SSH 连接
	portStr := strconv.FormatUint(uint64(port), 10)
	client, err := sshManager.CreateSSHClient(host, portStr, username, password)
	if err != nil {
		return nil, fmt.Errorf("SSH 连接失败：%w", err)
	}
	defer client.Close()

	// 创建会话
	sftpClient, err := sshManager.CreateSFTPClient(client)
	if err != nil {
		return nil, fmt.Errorf("创建 SSH 会话失败：%w", err)
	}
	defer sftpClient.Close()

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("打开文件 %s 失败：%w", file, err)
	}

	remotePath := fmt.Sprintf("/tmp/%s", filename)
	remotefile, err := sftpClient.Create(remotePath)
	if err != nil {
		return nil, fmt.Errorf("创建远程文件 %s 失败：%w", remotePath, err)
	}
	_, err = io.Copy(remotefile, file)
	if err != nil {
		return nil, fmt.Errorf("写入远程文件 %s 失败：%w", remotePath, err)
	}
	return &SSHCommandResult{
		Success:      true,
		Output:       fmt.Sprintf("文件 %s 上传成功", filepath),
		ErrorOutput:  "",
		SuccessCount: 1,
	}, nil
}
