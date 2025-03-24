package repository

import (
	"petHealthTool/common"
	"petHealthTool/models"
)

type UsersRepo struct {
}

// GetUserByName 根据用户名获取用户信息
func (u *UsersRepo) GetUserByName(name string) (models.Users, error) {
	var userModel models.Users
	result := common.DB.First(&userModel, "name = ?", name)
	return userModel, result.Error
}
