// Package response 统一响应格式和错误处理
package response

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIResponse 统一API响应结构
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// BusinessError 业务错误类型
type BusinessError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *BusinessError) Error() string {
	return e.Message
}

// ErrorCodes 预定义错误码
const (
	ErrCodeInvalidParams   = 40001
	ErrCodeUnauthorized    = 40101
	ErrCodeForbidden       = 40301
	ErrCodeNotFound        = 40401
	ErrCodeInternalError   = 50001
	ErrCodeDatabaseError   = 50002
	ErrCodeValidationError = 40002
)

// Predefined errors 预定义错误
var (
	ErrInvalidParams   = &BusinessError{Code: ErrCodeInvalidParams, Message: "参数错误"}
	ErrUnauthorized    = &BusinessError{Code: ErrCodeUnauthorized, Message: "未授权访问"}
	ErrForbidden       = &BusinessError{Code: ErrCodeForbidden, Message: "权限不足"}
	ErrNotFound        = &BusinessError{Code: ErrCodeNotFound, Message: "资源不存在"}
	ErrInternalError   = &BusinessError{Code: ErrCodeInternalError, Message: "服务器内部错误"}
	ErrDatabaseError   = &BusinessError{Code: ErrCodeDatabaseError, Message: "数据库操作失败"}
	ErrValidationError = &BusinessError{Code: ErrCodeValidationError, Message: "数据验证失败"}
)

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "操作成功",
		Data:    data,
	})
}

// SuccessWithMessage 带消息的成功响应
func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error 错误响应
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

// ErrorWithCode 带错误码的错误响应
func ErrorWithCode(c *gin.Context, code int, message string) {
	c.JSON(code, APIResponse{
		Success: false,
		Message: message,
		Error:   message,
	})
}

// ValidationError 参数验证错误
func ValidationError(c *gin.Context, field string, message string) {
	c.JSON(http.StatusBadRequest, APIResponse{
		Success: false,
		Message: "参数验证失败",
		Error:   fmt.Sprintf("%s: %s", field, message),
	})
}

// Unauthorized 未授权错误
func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, APIResponse{
		Success: false,
		Message: message,
		Error:   message,
	})
}

// Forbidden 权限不足错误
func Forbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, APIResponse{
		Success: false,
		Message: message,
		Error:   message,
	})
}

// NotFound 资源不存在错误
func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, APIResponse{
		Success: false,
		Message: message,
		Error:   message,
	})
}
