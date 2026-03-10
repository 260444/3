// Package utils 提供系统通用的工具函数。
//
// 该包包含 JWT token 生成和解析、上下文处理、验证码生成等功能。
package utils

import (
	"backend/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims 表示 JWT token 的声明信息。
//
// 该结构体包含用户 ID、用户名、角色标识符以及标准的 JWT 声明。
type Claims struct {
	UserID   uint   `json:"user_id"`  // 用户 ID
	Username string `json:"username"` // 用户名
	Ident    string `json:"ident"`    // 角色标识符
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT token。
//
// 该函数使用 HS256 签名算法生成 token，包含用户信息和过期时间。
//
// 参数：
//   - userID: 用户 ID
//   - username: 用户名
//   - ident: 角色标识符
//
// 返回：
//   - string: 生成的 JWT token 字符串
//   - error: 如果生成失败，返回错误
func GenerateToken(userID uint, username string, ident string) (string, error) {

	secret := config.GlobalConfig.JWT.Secret
	timeout := config.GlobalConfig.JWT.Timeout
	//logger.Logger.Info("生成JWT", zap.String("secret", secret), zap.String("timeout", timeout.String()))
	claims := &Claims{
		UserID:   userID,
		Username: username,
		Ident:    ident,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(timeout * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "admin-system",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseToken 解析 JWT token。
//
// 该函数验证 token 的签名和有效期，并提取用户信息。
//
// 参数：
//   - tokenString: JWT token 字符串
//
// 返回：
//   - *Claims: 包含用户信息的声明结构体
//   - error: 如果 token 无效、过期或签名错误，返回错误
func ParseToken(tokenString string) (*Claims, error) {

	secret := config.GlobalConfig.JWT.Secret

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的token")
}
