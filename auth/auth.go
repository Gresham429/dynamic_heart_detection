package auth

import (
	"dynamic_heart_rates_detection/config"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	UserName string `json:"name"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}

func GenerateJWTToken(username string) (string, error) {
	// 设置 claims
	claims := &JwtCustomClaims{
		UserName: username,
		Admin:    false,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	// 用 claims 创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成 jwt 令牌
	tokenEncode, err := token.SignedString([]byte(config.JsonConfiguration.JwtSecret))

	return tokenEncode, err
}
