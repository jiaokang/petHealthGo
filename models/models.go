package models

import (
	"time"

	"gorm.io/gorm"
)

// 用户信息
type Users struct {
	gorm.Model
	NickName string `gorm:"size:64;comment:昵称"`
	Phone    string `gorm:"size:20;uniqueIndex;comment:主人手机号"`
	Email    string `gorm:"size:100;uniqueIndex;comment:主人邮箱"`
	Address  string `gorm:"type:text;comment:主人地址"`
	Pwd      string `gorm:"size:64;comment:密码"`
}

// 认证方式
type AuthMethods struct {
	gorm.Model
	UserId     uint   `gorm:"index;comment:用户ID"`
	AuthType   string `gorm:"size:20;comment:认证类型"`
	AuthValue  string `gorm:"size:100;comment:认证值"`
	Credential string `gorm:"size:100;comment:认证凭据"`
}

// 宠物信息
type Pets struct {
	gorm.Model
	Name     string    `gorm:"size:20;comment:宠物名称"`
	Breed    string    `gorm:"size:50;comment:宠物品种"`
	Brithday time.Time `gorm:"not null;comment:宠物品种"`
	UserId   uint      `gorm:"index;comment:用户ID"`
	Avatar   string    `gorm:"size:100;comment:宠物头像"`
}

// 宠物驱虫记录
type DewormingRecords struct {
	gorm.Model
	PetId       uint      `gorm:"index;comment:宠物ID"`
	RecordDate  time.Time `gorm:"not null;comment:记录日期"`
	Weight      float64   `gorm:"not null;comment:体重"`
	Medicine    string    `gorm:"size:100;comment:药物"`
	Temperature float64   `gorm:"not null;comment:体温"`
	Age         int       `gorm:"not null;comment:年龄"`
	HealthState string    `gorm:"size:100;comment:健康状态"`
	Remark      string    `gorm:"size:100;comment:备注"`
}

// 接种记录
type VaccinationRecords struct {
	gorm.Model
	PetId       uint      `gorm:"index;comment:宠物ID"`
	RecordDate  time.Time `gorm:"not null;comment:记录日期"`
	Weight      float64   `gorm:"not null;comment:体重"`
	Medicine    string    `gorm:"size:100;comment:药物"`
	Temperature float64   `gorm:"not null;comment:体温"`
	Age         int       `gorm:"not null;comment:年龄"`
	HealthState string    `gorm:"size:100;comment:健康状态"`
	Remark      string    `gorm:"size:100;comment:备注"`
}

// 任务排期
type Scheduleds struct {
	gorm.Model
	PetId        uint      `gorm:"index;comment:宠物ID"`
	UserId       uint      `gorm:"index;comment:用户ID"`
	TaskType     string    `gorm:"size:20;comment:任务类型"`
	ExpectDate   time.Time `gorm:"not null;comment:预计日期"`
	ExecuteDate  time.Time `gorm:"comment:执行日期"`
	ExecuteState bool      `gorm:"comment:执行状态"`
	NotiftyState bool      `gorm:"comment:通知状态"`
}

func (Users) TableName() string {
	return "users"
}

func (AuthMethods) TableName() string {
	return "auth_methods"
}

func (Pets) TableName() string {
	return "pets"
}

func (DewormingRecords) TableName() string {
	return "deworming_records"
}

func (VaccinationRecords) TableName() string {
	return "vaccination_records"
}
func (Scheduleds) TableName() string {
	return "scheduleds"
}
