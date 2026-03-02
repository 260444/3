package asset_management

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	assModel "backend/internal/model/asset_management"
	assService "backend/internal/service/asset_management"
	"backend/pkg/response"
	"backend/pkg/ssh"
	"backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type SSHHandler struct {
	HostService       *assService.HostService
	CredentialService *assService.CredentialService
	SSHManager        *ssh.SSHManager
}

func NewSSHHandler(HostService *assService.HostService, CredentialService *assService.CredentialService) *SSHHandler {
	return &SSHHandler{
		HostService:       HostService,
		CredentialService: CredentialService,
		SSHManager:        ssh.NewSSHManager(),
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源的 WebSocket 连接
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// HandleSSHWebSocket 处理 SSH WebSocket 连接
// @Summary SSH WebSocket 连接
// @Description 建立 SSH WebSocket 连接用于 Web 终端
// @Tags SSH
// @Param host_id query string true "主机名或ID"
// @Param credential_id query int true "凭证ID"
// @Param token query string true "JWT token"
// @Success 101 {string} string "WebSocket 连接建立"
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/ssh/ws [get]
// @Security Bearer
func (h *SSHHandler) HandleSSHWebSocket(c *gin.Context) {
	// 验证 JWT token（从 URL 参数获取）
	token := c.Query("token")
	if token == "" {
		// 尝试从 Authorization header 获取
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "缺少认证信息"})
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "认证格式错误"})
			return
		}
		token = parts[1]
	}

	// 解析 token
	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 Token"})
		return
	}

	// 设置用户信息到上下文
	c.Set("userID", claims.UserID)
	c.Set("username", claims.Username)
	c.Set("ident", claims.Ident)

	// 获取主机信息
	hostIdentifier := c.Query("host_id")
	if hostIdentifier == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "host_id 参数不能为空"})
		return
	}

	// 获取凭证ID
	credentialIDStr := c.Query("credential_id")
	if credentialIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "credential_id 参数不能为空"})
		return
	}

	credentialID, err := strconv.ParseUint(credentialIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的 credential_id"})
		return
	}

	// 获取主机信息（支持通过主机名或ID查询）
	var host *assModel.Host
	if id, err := strconv.ParseUint(hostIdentifier, 10, 32); err == nil {
		// 通过ID查询
		host, err = h.HostService.GetHostByID(uint(id))
	} else {
		// 通过主机名查询（需要实现这个方法）
		c.JSON(http.StatusBadRequest, gin.H{"error": "请提供有效的主机ID"})
		return
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "主机不存在"})
		return
	}

	// 获取凭证信息
	credential, err := h.CredentialService.GetCredentialByID(uint(credentialID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "凭证不存在"})
		return
	}

	// 检查权限
	userID := utils.GetUserIDFromContext(c)
	if !h.HostService.CheckHostAccess(host.ID, userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权访问该主机"})
		return
	}

	// 升级到 WebSocket 连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("WebSocket 升级失败: %v\n", err)
		return
	}
	defer conn.Close()

	// 转换端口为字符串
	port := strconv.Itoa(int(host.Port))

	// 处理 SSH 连接
	h.handleSSHConnection(conn, host.IPAddress, port, credential.Username, credential.Password)
}

// handleSSHConnection 处理 SSH 连接
func (h *SSHHandler) handleSSHConnection(wsConn *websocket.Conn, host, port, username, password string) {
	// 创建 SSH 客户端
	client, err := h.SSHManager.CreateSSHClient(host, port, username, password)
	if err != nil {
		sendWSMessage(wsConn, "error", fmt.Sprintf("SSH连接失败: %v", err))
		return
	}
	defer client.Close()

	// 创建会话
	session, err := h.SSHManager.CreateSession(client)
	if err != nil {
		sendWSMessage(wsConn, "error", fmt.Sprintf("创建SSH会话失败: %v", err))
		return
	}
	defer h.SSHManager.CloseSession(session)

	// 启动 Shell
	if err := h.SSHManager.StartShell(session); err != nil {
		sendWSMessage(wsConn, "error", fmt.Sprintf("启动Shell失败: %v", err))
		return
	}

	// 发送欢迎消息
	sendWSMessage(wsConn, "output", "\r\n\x1b[32mSSH 连接已建立\x1b[0m\r\n")

	// 读取输出的 channel
	outputChan := make(chan []byte, 1024)
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := session.StdoutPipe.Read(buf)
			if err != nil {
				if err.Error() != "EOF" {
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

	// 处理输入
	inputChan := make(chan ssh.SSHMessage, 10)
	go func() {
		for {
			var msg ssh.SSHMessage
			if err := wsConn.ReadJSON(&msg); err != nil {
				close(inputChan)
				return
			}
			inputChan <- msg
		}
	}()

	// 主循环
	for {
		select {
		case data, ok := <-outputChan:
			if !ok {
				sendWSMessage(wsConn, "close", "")
				return
			}
			if err := sendWSMessage(wsConn, "output", string(data)); err != nil {
				return
			}
		case msg, ok := <-inputChan:
			if !ok {
				return
			}
			switch msg.Type {
			case "input":
				if _, err := session.StdinPipe.Write([]byte(msg.Data)); err != nil {
					sendWSMessage(wsConn, "error", fmt.Sprintf("写入输入错误: %v", err))
					return
				}
			case "resize":
				if err := h.SSHManager.ResizePty(session, msg.Rows, msg.Cols); err != nil {
					sendWSMessage(wsConn, "error", fmt.Sprintf("调整终端大小错误: %v", err))
				}
			case "close":
				return
			}
		}
	}
}

// sendWSMessage 发送 WebSocket 消息
func sendWSMessage(wsConn *websocket.Conn, msgType, data string) error {
	msg := ssh.SSHMessage{
		Type: msgType,
		Data: data,
	}
	return wsConn.WriteJSON(msg)
}

// TestSSHConnection 测试 SSH 连接
// @Summary 测试 SSH 连接
// @Description 测试指定主机的 SSH 连接是否可用
// @Tags SSH
// @Accept json
// @Produce json
// @Param request body TestSSHRequest true "测试请求"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/ssh/test [post]
// @Security Bearer
func (h *SSHHandler) TestSSHConnection(c *gin.Context) {
	var req TestSSHRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err)
		return
	}

	// 获取凭证信息
	credential, err := h.CredentialService.GetCredentialByID(req.CredentialID)
	if err != nil {
		response.NotFound(c, "凭证不存在")
		return
	}

	// 获取主机信息
	host, err := h.HostService.GetHostByID(req.HostID)
	if err != nil {
		response.NotFound(c, "主机不存在")
		return
	}

	// 检查权限
	userID := utils.GetUserIDFromContext(c)
	if !h.HostService.CheckHostAccess(host.ID, userID) {
		response.Forbidden(c, "无权访问该主机")
		return
	}

	// 测试连接
	port := strconv.Itoa(int(host.Port))
	err = ssh.TestConnection(host.IPAddress, port, credential.Username, credential.Password)
	if err != nil {
		response.SuccessWithMessage(c, "连接失败", gin.H{
			"success": false,
			"message": fmt.Sprintf("连接失败: %v", err),
			"details": gin.H{
				"host":     host.IPAddress,
				"port":     port,
				"username": credential.Username,
			},
		})
		return
	}

	response.SuccessWithMessage(c, "连接成功", gin.H{
		"success": true,
		"message": "连接成功",
	})
}

// TestSSHRequest 测试 SSH 连接请求
type TestSSHRequest struct {
	HostID       uint `json:"host_id" binding:"required"`
	CredentialID uint `json:"credential_id" binding:"required"`
}
