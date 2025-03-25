package common

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Mail struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
	} `yaml:"mail"`
	Redis struct {
		Host string `yaml:"host"`
		Pass string `yaml:"pass"`
		Db   int    `yaml:"db"`
	} `yaml:"redis"`
}

var cfg *Config // 全局变量，存储配置

// LoadConfig 读取并解析配置文件
func LoadConfig(filename string) (*Config, error) {
	// 读取配置文件
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// 解析配置文件
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	// 将配置存储到全局变量
	cfg = &config
	return cfg, nil
}

// GetConfig 获取全局配置
func GetConfig() *Config {
	if cfg == nil {
		panic("Config is not loaded. Please call LoadConfig first.")
	}
	return cfg
}
