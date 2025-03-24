package common

import (
	"fmt"
	"os"
	"petHealthTool/models"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func LoadConfig(filename string) (*Config, error) {
	// 读取配置文件
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// 解析配置文件
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

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
