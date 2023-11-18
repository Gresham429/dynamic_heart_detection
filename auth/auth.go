package auth

import (
	"context"
	"dynamic_heart_rates_detection/config"
	"dynamic_heart_rates_detection/model"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	UserName string `json:"name"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}

// 生成 JWT Token
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

// 生成 6 位邮箱验证码
func GenerateVerificationCode() string {
	rand.NewSource(time.Now().Unix())
	return fmt.Sprintf("%6d", rand.Intn(1000000))
}

// 发送邮箱验证码
func SendEmail(email, verificationCode string) {
	// 在这里实现发送邮件的逻辑，使用你选择的邮件服务或库
	fmt.Printf("发送邮件至 %s，验证码为: %s\n", email, verificationCode)
}

// VerifyVerificationCode - 验证邮箱验证码
func VerifyVerificationCode(verificationCode string, email string, ctx context.Context) error {
	// 验证验证码
	storedCode, err := model.GetVerificationCode(email, ctx)
	if err != nil {
		return errors.New("验证码无效或已过期。")
	}

	if storedCode != verificationCode {
		return errors.New("验证码不匹配。")
	}

	// 验证通过，可以进行登录操作
	// 清除 Redis 中的验证码信息
	err = model.CleanVerificationCode(email, ctx)
	if err != nil {
		return err
	}

	return nil
}
