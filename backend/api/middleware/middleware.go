package middleware

import (
	"backend/pkg/casbin"
	"backend/pkg/logger"
	"backend/pkg/utils"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// JWTAuthMiddleware JWT认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "请求头缺少Authorization字段",
			})
			c.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization格式错误",
			})
			c.Abort()
			return
		}

		// 解析JWT
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "无效的Token",
			})
			c.Abort()
			return
		}

		// 将用户ID和用户名保存到上下文中
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("ident", claims.Ident)

		c.Next()
	}
}

// CORSMiddleware 跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, X-CSRF-Token")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// LoggerToFile 日志中间件 - 记录HTTP请求日志到文件
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 记录日志到文件
		logger.Logger.Info("HTTP请求",
			zap.String("method", reqMethod),
			zap.String("uri", reqUri),
			zap.Int("status", statusCode),
			zap.String("client_ip", clientIP),
			zap.Duration("latency", latencyTime),
			zap.String("user_agent", c.Request.UserAgent()),
		)
	}
}

// RateLimitMiddleware 限流中间件（简单实现）
func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里可以集成Redis进行限流
		// 简单示例，实际应用中需要更复杂的逻辑
		c.Next()
	}
}

// OperationLogMiddleware 操作日志中间件 - 记录用户操作日志
func OperationLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户信息
		username, _ := c.Get("username")
		userID, _ := c.Get("userID")

		// 记录请求开始前的信息
		startTime := time.Now()
		method := c.Request.Method
		path := c.FullPath()
		clientIP := c.ClientIP()

		// 处理请求
		c.Next()

		// 请求结束后记录操作日志
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		statusCode := c.Writer.Status()

		// 记录操作日志
		logger.Logger.Info("用户操作日志",
			zap.Any("user_id", userID),
			zap.Any("username", username),
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", statusCode),
			zap.String("ip", clientIP),
			zap.Duration("duration", latency),
		)
	}
}

// CasbinMiddleware Casbin权限验证中间件
func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户ID
		ident, exists := c.Get("ident")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "未授权访问",
			})
			c.Abort()
			return
		}

		// 获取请求方法和路径
		method := c.Request.Method
		path := c.FullPath()

		// 如果没有路径（可能是404），跳过权限检查
		if path == "" {
			c.Next()
			return
		}

		// 特殊处理：用户访问自己的信息时，允许通过
		// 检查是否是获取用户信息的请求
		if path == "/api/v1/users/profile" {
			// 用户可以访问自己的信息，允许通过
			c.Next()
			return
		}

		// 将用户ID转换为字符串
		var sub string
		switch v := ident.(type) {
		case string:
			sub = v
		case int:
			sub = fmt.Sprintf("%d", v)
		case int64:
			sub = fmt.Sprintf("%d", v)
		default:
			sub = fmt.Sprintf("%v", v)
		}
		obj := path
		act := method

		// 使用Casbin检查权限
		allowed, err := casbin.Enforcer.Enforce(sub, obj, act)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "权限检查失败",
			})
			c.Abort()
			return
		}

		if !allowed {
			// 检查请求是否是AJAX请求或API请求
			if c.GetHeader("X-Requested-With") == "XMLHttpRequest" || 
				strings.Contains(c.GetHeader("Accept"), "application/json") {
				// API请求，返回JSON错误
				c.JSON(http.StatusForbidden, gin.H{
					"error": "权限不足",
				})
			} else {
				// 页面请求，重定向到无权限页面
				c.Redirect(http.StatusFound, "/no-permission")
			}
			c.Abort()
			return
		}
		c.Next()
	}
}
