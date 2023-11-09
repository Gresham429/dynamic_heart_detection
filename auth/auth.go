package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	UserName string `json:"name"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}

func GetJwtClaims(c echo.Context) (claims JwtCustomClaims, err error) {
	authorization := strings.Split(c.Request().Header.Get("Authorization"), " ")

	if len(authorization) < 2 {
		return JwtCustomClaims{}, errors.New("请求头不合法")
	}

	if authorization[0] != "Bearer" {
		return JwtCustomClaims{}, errors.New("请求头不合法")
	}

	jwt_token := authorization[1]
	claims = JwtCustomClaims{}

	_, err = jwt.ParseWithClaims(jwt_token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("Gresham"), nil
	})

	return claims, err
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
	jwt, err := token.SignedString([]byte("Gresham"))

	return jwt, err
}
