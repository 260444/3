// Package middleware 全局中间件
package middleware

import (
	"backend/pkg/logger"
	"backend/pkg/response"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RecoveryMiddleware 全局异常恢复中间件
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录详细的错误日志
				logger.Logger.Error("panic recovered",
					zap.Any("error", err),
					zap.String("stack", string(debug.Stack())),
					zap.String("method", c.Request.Method),
					zap.String("path", c.Request.URL.Path),
					zap.String("client_ip", c.ClientIP()),
				)

				// 返回统一错误响应
				response.ErrorWithCode(c, http.StatusInternalServerError, "服务器内部错误，请稍后重试")
				c.Abort()
			}
		}()
		c.Next()
	}
}

// ErrorHandlingMiddleware 统一错误处理中间件
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 处理累积的错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// 根据错误类型返回不同的响应
			switch err.Type {
			case gin.ErrorTypeBind:
				response.ValidationError(c, "请求参数", err.Error())
			case gin.ErrorTypeRender:
				response.ErrorWithCode(c, http.StatusInternalServerError, "响应渲染失败")
			default:
				response.ErrorWithCode(c, http.StatusInternalServerError, "未知错误")
			}
		}
	}
}

// RequestIDMiddleware 请求ID中间件
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成唯一的请求ID
		requestID := fmt.Sprintf("%d", c.GetInt("request_id"))
		if requestID == "0" {
			requestID = fmt.Sprintf("%d", c.Writer.Size())
		}

		// 添加到响应头
		c.Header("X-Request-ID", requestID)

		// 添加到上下文
		c.Set("request_id", requestID)

		c.Next()
	}
}

// SecurityHeadersMiddleware 安全头中间件
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 添加安全相关的HTTP头
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		c.Next()
	}
}

// ValidationMiddleware 参数验证中间件
func ValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 可以在这里添加通用的参数验证逻辑
		// 例如：验证必填字段、格式检查等

		c.Next()
	}
}
