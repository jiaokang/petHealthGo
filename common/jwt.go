package common

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
}
type CustomClaims struct {
	UserId uint `json:"userId"`
	jwt.RegisteredClaims
}

// NewJwt 创建一个jwt实例
func (j *Jwt) CreateToken(userId uint) (string, error) {
	securityKey := []byte("petpetpet")

	claims := CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "petHealthTool",
			Subject:   "user token",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        "userId",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(securityKey)
}

// ParseToken 解析并验证 JWT Token
func (j *Jwt) ParseToken(tokenString string) (*CustomClaims, error) {
	securityKey := []byte("petpetpet")

	// 解析 Token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return securityKey, nil
	})

	// 处理解析错误
	if err != nil {
		return nil, err
	}

	// 验证 Token 是否有效
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
