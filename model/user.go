package model

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primary_email;unique;column:id"`
	UserName string `gorm:"primary_email;unique;column:user_name"`
	Password string `gorm:"column:password"`

	// 可为空字段
	FullName string `gorm:"column:full_name,default:NULL"`
	Email    string `gorm:"column:email,default:NULL"`
	Address  string `gorm:"column:address,default:NULL"`
}

// Create - 创建用户
func CreateUser(user *User) error {
	result := DB.Create(user)
	return result.Error
}

// Read - 读取用户信息
func GetUserInfo(username string) (*User, error) {
	user := &User{}
	result := DB.Where("user_name = ?", username).First(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// Update - 更新用户信息
func UpdateUser(user *User) error {
	result := DB.Save(user)
	return result.Error
}

// Delete - 删除用户
func DeleteUser(username string) error {
	result := DB.Where("user_name = ?", username).Delete(&User{})
	return result.Error
}

// StoreVerificationCode - 存储 6 位的邮箱验证码(会覆盖前一次的验证码)
func StoreVerificationCode(email string, verificationCode string, ctx context.Context) error {
	err := RDB.Set(ctx, "verificationCode:"+email, verificationCode, 5*time.Minute).Err()

	if err != nil {
		return err
	}

	err = RDB.Set(ctx, "lastSentTime:"+email, time.Now().Unix(), 0).Err()

	return err
}

// CleanVerificationCode - 清除邮箱验证码信息
func CleanVerificationCode(email string, ctx context.Context) error {
	err := RDB.Del(ctx, "verificationCode:"+email).Err()

	if err != nil {
		return err
	}

	err = RDB.Del(ctx, "lastSentTime:"+email).Err()

	return err
}

func GetVerificationCode(email string, ctx context.Context) (string, error) {
	verificationCode, err := RDB.Get(ctx, "verificationCode:"+email).Result()
	return verificationCode, err
}

func GetLastSentTime(email string, ctx context.Context) (string, error) {
	lastSentTime, err := RDB.Get(ctx, "lastSentTime:"+email).Result()
	return lastSentTime, err
}
