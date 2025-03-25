package main

import (
	"os"
	"petHealthTool/common"
	"petHealthTool/routes"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// 加载配置文件
	_, err := common.LoadConfig("config.yml")
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}
	cfg := common.GetConfig()

	// 初始化数据库
	common.InitDB(cfg)
	// 获取 Redis 客户端实例（单例模式）
	redisClient := common.GetRedisClient(cfg.Redis.Host, cfg.Redis.Pass, cfg.Redis.Db)
	defer redisClient.Close()

	// 设置日志级别
	logrus.SetLevel(logrus.DebugLevel)
	// 设置日志输出格式为 JSON
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	router := gin.Default()

	routes.RegisterAuthRoutes(router)

	router.Run(":9000")

}
