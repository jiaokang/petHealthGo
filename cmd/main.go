package main

import (
	"os"
	"petHealthTool/common"
	"petHealthTool/routes"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// 加载数据库配置文件
	cfg, err := common.LoadConfig("E:/petHealthGo/config.yml")
	if err != nil {
		panic("faild to load config.yml file")
	}

	// 设置日志级别
	logrus.SetLevel(logrus.DebugLevel)
	// 设置日志输出格式为 JSON
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 将日志输出到文件
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Failed to open log file:", err)
	}
	defer file.Close()
	logrus.SetOutput(file)
	logrus.Info("This log is written to a file")

	common.InitDB(cfg)

	router := gin.Default()

	routes.RegisterAuthRoutes(router)

	router.Run(":9000")

}
