package auth

import (
	"context"
	"dynamic_heart_rates_detection/config"
	"dynamic_heart_rates_detection/model"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"net/mail"
	"net/smtp"
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
func SendEmail(email, verificationCode, userName string) error {
	// 发件人邮箱和密码
	from := "1543732388@qq.com"
	password := "oocqkxaqwvpqhgcd"

	// 收件人邮箱
	to := email

	// SMTP 服务器地址和端口
	smtpHost := "smtp.qq.com"
	smtpPort := 587

	// 设置邮件标头
	fromAddress := mail.Address{Name: "", Address: from}
	toAddress := mail.Address{Name: "", Address: to}

	// 设置昵称
	nickname := "南笙"
	nicknameEncoded := base64.StdEncoding.EncodeToString([]byte(nickname))
	fromHeader := fmt.Sprintf("=?utf-8?B?%s?=%s", nicknameEncoded, fromAddress.Address)

	// 邮件标头
	message := "Subject: Your Subject\r\n" +
		"From: " + fromHeader + "\r\n" +
		"To: " + toAddress.String() + "\r\n" +
		"\r\n" +
		"亲爱的" + userName + "：\n" + "您的验证码是：" + verificationCode + "。请在五分钟内完成注册"

	// 连接到 SMTP 服务器
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+fmt.Sprint(smtpPort), auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}

// VerifyVerificationCode - 验证邮箱验证码
func VerifyVerificationCode(verificationCode, email string, ctx context.Context) error {
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
