package model

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primary_key;unique;column:id"`
	UserName string `gorm:"unique;column:user_name"`
	Password string `gorm:"column:password"`

	// 可为空字段
	FullName    string `gorm:"column:full_name"`
	Email       string `gorm:"unique;column:email"`
	PhoneNumber string `gorm:"unique;column:phone_number"`
	Address     string `gorm:"column:address"`
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
