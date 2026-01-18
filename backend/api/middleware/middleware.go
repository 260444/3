package middleware

import (
	"backend/pkg/casbin"
	"backend/pkg/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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

// LoggerToFile 日志中间件
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
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

// OperationLogMiddleware 操作日志中间件
func OperationLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录操作日志的逻辑
		// 在请求处理后记录日志
		c.Next()
	}
}

// CasbinMiddleware Casbin权限验证中间件
func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户ID
		userID, exists := c.Get("userID")
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

		// 将用户ID转换为字符串
		sub := fmt.Sprintf("user_%v", userID)
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
			c.JSON(http.StatusForbidden, gin.H{
				"error": "权限不足",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
