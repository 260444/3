package utils

import (
	"backend/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT Claims
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Ident    string `json:"ident"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT token
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

// ParseToken 解析JWT token
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
