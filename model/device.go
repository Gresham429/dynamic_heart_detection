package model

import (
	"errors"

	"gorm.io/gorm"
)

type Device struct {
	ID        uint   `json:"id" gorm:"primary_key;unique;column:id"`
	Url       string `json:"url" gorm:"colum:url"`
	Position  string `json:"position" gorm:"colum:position"`
	Connected bool   `json:"connected" gorm:"colum:connected"`
}

func CreateDevice(device *Device) error {
	result := DB.Create(device)
	return result.Error
}

// Read - 读取设备信息
func GetDeviceInfo(url string) (*Device, error) {
	device := &Device{}
	result := DB.Where("url = ?", url).First(device)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return device, nil
}

// Update - 更新设备信息
func UpdateDevice(Device *Device) error {
	result := DB.Save(Device)
	return result.Error
}

// Delete - 删除设备
func DeleteDevice(url string) error {
	result := DB.Where("url = ?", url).Delete(&Device{})
	return result.Error
}

// ReadAll - 查询所有设备信息
func GetAllDevices() ([]Device, error) {
	var devices []Device

	// 执行查询，并将结果存储在 data 切片中
	err := DB.Find(&devices).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return devices, nil
}
