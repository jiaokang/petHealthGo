package main

import (
	"petHealthTool/common"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载数据库配置文件
	cfg, err := common.LoadConfig("config.yml")
	if err != nil {
		panic("faild to load config.yml file")
	}

	common.InitDB(cfg)

	router := gin.Default()

	router.Run(":9000")

}
