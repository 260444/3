// Package response 提供统一的 HTTP 响应格式和错误处理。
//
// 该包定义了标准的 API 响应结构、业务错误类型以及常用的响应辅助函数。
// 所有 API 接口都应该使用该包提供的函数来返回响应。
package response

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIResponse 定义统一的 API 响应结构。
//
// 所有 API 接口的响应都应遵循此格式，包含成功状态、消息、数据和错误信息。
type APIResponse struct {
	Success bool        `json:"success"`         // 请求是否成功
	Message string      `json:"message"`         // 响应消息
	Data    interface{} `json:"data,omitempty"`  // 响应数据（可选）
	Error   string      `json:"error,omitempty"` // 错误信息（可选）
}

// BusinessError 定义业务错误类型。
//
// 该类型用于封装业务逻辑中的错误，包含错误码和错误消息。
type BusinessError struct {
	Code    int    `json:"code"`    // 业务错误码
	Message string `json:"message"` // 错误消息
}

// Error 实现 error 接口。
func (e *BusinessError) Error() string {
	return e.Message
}

// ErrorCodes 预定义错误码常量。
const (
	ErrCodeInvalidParams   = 40001 // 参数错误
	ErrCodeUnauthorized    = 40101 // 未授权
	ErrCodeForbidden       = 40301 // 禁止访问
	ErrCodeNotFound        = 40401 // 资源不存在
	ErrCodeInternalError   = 50001 // 内部错误
	ErrCodeDatabaseError   = 50002 // 数据库错误
	ErrCodeValidationError = 40002 // 验证错误
)

// Predefined errors 预定义错误实例。
var (
	ErrInvalidParams   = &BusinessError{Code: ErrCodeInvalidParams, Message: "参数错误"}
	ErrUnauthorized    = &BusinessError{Code: ErrCodeUnauthorized, Message: "未授权访问"}
	ErrForbidden       = &BusinessError{Code: ErrCodeForbidden, Message: "权限不足"}
	ErrNotFound        = &BusinessError{Code: ErrCodeNotFound, Message: "资源不存在"}
	ErrInternalError   = &BusinessError{Code: ErrCodeInternalError, Message: "服务器内部错误"}
	ErrDatabaseError   = &BusinessError{Code: ErrCodeDatabaseError, Message: "数据库操作失败"}
	ErrValidationError = &BusinessError{Code: ErrCodeValidationError, Message: "数据验证失败"}
)

// Success 返回成功的响应。
//
// 参数：
//   - c: Gin 上下文
//   - data: 响应数据
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "操作成功",
		Data:    data,
	})
}

// SuccessWithMessage 返回带自定义消息的成功响应。
//
// 参数：
//   - c: Gin 上下文
//   - message: 自定义消息
//   - data: 响应数据
func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error 根据错误类型返回相应的错误响应。
//
// 该函数会自动判断错误类型并返回相应的 HTTP 状态码和错误消息。
//
// 参数：
//   - c: Gin 上下文
//   - err: 错误实例
func Error(c *gin.Context, err error) {
	var businessErr *BusinessError

	// 判断是否为业务错误
	if errors.As(err, &businessErr) {
		switch businessErr.Code {
		case ErrCodeInvalidParams, ErrCodeValidationError:
			c.JSON(http.StatusBadRequest, APIResponse{
				Success: false,
				Message: businessErr.Message,
				Error:   businessErr.Message,
			})
		case ErrCodeUnauthorized:
			c.JSON(http.StatusUnauthorized, APIResponse{
				Success: false,
				Message: businessErr.Message,
				Error:   businessErr.Message,
			})
		case ErrCodeForbidden:
			c.JSON(http.StatusForbidden, APIResponse{
				Success: false,
				Message: businessErr.Message,
				Error:   businessErr.Message,
			})
		case ErrCodeNotFound:
			c.JSON(http.StatusNotFound, APIResponse{
				Success: false,
				Message: businessErr.Message,
				Error:   businessErr.Message,
			})
		default:
			c.JSON(http.StatusInternalServerError, APIResponse{
				Success: false,
				Message: businessErr.Message,
				Error:   businessErr.Message,
			})
		}
	} else {
		// 系统错误
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Message: "服务器内部错误",
			Error:   err.Error(),
		})
	}
}

// ErrorWithCode 返回带指定 HTTP 状态码的错误响应。
//
// 参数：
//   - c: Gin 上下文
//   - code: HTTP 状态码
//   - message: 错误消息
func ErrorWithCode(c *gin.Context, code int, message string) {
	c.JSON(code, APIResponse{
		Success: false,
		Message: message,
		Error:   message,
	})
}

// ValidationError 返回参数验证错误的响应。
//
// 参数：
//   - c: Gin 上下文
//   - field: 参数字段名
//   - message: 验证失败消息
func ValidationError(c *gin.Context, field string, message string) {
	c.JSON(http.StatusBadRequest, APIResponse{
		Success: false,
		Message: "参数验证失败",
		Error:   fmt.Sprintf("%s: %s", field, message),
	})
}

// Unauthorized 返回未授权错误的响应。
//
// 参数：
//   - c: Gin 上下文
//   - message: 错误消息
func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, APIResponse{
		Success: false,
		Message: message,
		Error:   message,
	})
}

// Forbidden 返回权限不足错误的响应。
//
// 参数：
//   - c: Gin 上下文
//   - message: 错误消息
func Forbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, APIResponse{
		Success: false,
		Message: message,
		Error:   message,
	})
}

// NotFound 返回资源不存在错误的响应。
//
// 参数：
//   - c: Gin 上下文
//   - message: 错误消息
func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, APIResponse{
		Success: false,
		Message: message,
		Error:   message,
	})
}
