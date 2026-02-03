package middleware

import (
	sysModel "backend/internal/model/system_manager"
	sysService "backend/internal/service/system_manager"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

// OperationLogMiddleware 操作日志中间件 - 记录用户操作日志并写入数据库
func OperationLogMiddleware(operationLogService *sysService.OperationLogService, operation string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户信息
		username, exists1 := c.Get("username")
		userID, exists2 := c.Get("userID")

		// 如果没有用户信息，跳过日志记录
		if !exists1 || !exists2 {
			c.Next()
			return
		}

		// 记录请求开始前的信息
		startTime := time.Now()
		method := c.Request.Method
		path := c.FullPath()
		clientIP := c.ClientIP()

		// 捕获响应体
		buffer := &bytes.Buffer{}
		writer := &ResponseCaptureWriter{c.Writer, buffer}
		c.Writer = writer

		// 读取请求体内容
		var requestBody string
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			requestBody = string(bodyBytes)
			// 恢复请求体供后续使用
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// 处理请求
		c.Next()

		// 请求结束后记录操作日志
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		statusCode := c.Writer.Status()

		// 构造操作日志对象
		logEntry := &sysModel.OperationLog{
			UserID:        userID.(uint),
			Username:      username.(string),
			Operation:     operation,
			IP:            clientIP,
			UserAgent:     c.Request.UserAgent(),
			RequestMethod: method,
			RequestPath:   path,
			RequestBody:   requestBody,
			ResponseBody:  buffer.String(),
			Status:        statusCode,
			ResponseTime:  int64(latency.Milliseconds()),
		}

		// 异步保存到数据库，避免阻塞主流程
		go func(log *sysModel.OperationLog) {
			if err := operationLogService.CreateOperationLog(log); err != nil {
				// 记录错误日志（这里可以使用zap或其他日志库）
				// logger.Error("保存操作日志失败", zap.Error(err))
			}
		}(logEntry)
	}
}

// ResponseCaptureWriter 包装 gin.ResponseWriter 以捕获响应体
type ResponseCaptureWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w ResponseCaptureWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // 将响应体写入 buffer
	return w.ResponseWriter.Write(b)
}
