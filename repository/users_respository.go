package repository

import (
	"errors"
	"petHealthTool/common"
	"petHealthTool/models"

	"gorm.io/gorm"
)

type UsersRepo struct {
}

// GetUserByName 根据用户名获取用户信息
func (u *UsersRepo) GetUserByName(name string) (models.Users, error) {
	var userModel models.Users
	result := common.DB.First(&userModel, "name = ?", name)
	return userModel, result.Error
}

func (u *UsersRepo) GetUserByEmail(email string) (*models.Users, error) {
	var userModel models.Users
	result := common.DB.First(&userModel, "email = ?", email)

	// 检查是否未找到记录
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound // 直接返回 gorm.ErrRecordNotFound
	}
	// 其他错误
	if result.Error != nil {
		return nil, result.Error
	}
	// 返回找到的记录
	return &userModel, nil
}

func (u *UsersRepo) CreateUser(user *models.Users) error {
	result := common.DB.Create(user)
	return result.Error
}
