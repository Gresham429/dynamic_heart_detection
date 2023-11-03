package middleware

import (
	"dynamic_heart_rates_detection/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserName string `json:"name"`
	Admin    bool   `json:"admin"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(user model.User) (string, error) {
	// 设置 claims
	claims := &JwtCustomClaims{
		UserName: user.UserName,
		Admin:    false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	// 用 claims 创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成 jwt 令牌
	jwt, err := token.SignedString([]byte("Gresham"))

	return jwt, err
}

func IsTokenExpired(tokenString string) (bool, error) {
	// 解析 JWT 令牌
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 这里添加用于验证签名的密钥，密钥必须和生成令牌时的密钥一致
		return []byte("Gresham"), nil
	})

	if err != nil {
		return false, err
	}

	// 检查令牌是否有效
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		// 检查令牌是否过期
		return claims.ExpiresAt.Before(time.Now()), nil
	}

	return false, err
}
