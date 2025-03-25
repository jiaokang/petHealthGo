package common

import (
	"fmt"
	"petHealthTool/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("inital database error, failed to connect database")
	}
	// 自动迁移表结构
	DB.AutoMigrate(&models.Users{}, &models.AuthMethods{}, &models.Pets{}, &models.VaccinationRecords{}, &models.Scheduleds{}, &models.DewormingRecords{})
}
