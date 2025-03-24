package common

import (
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
