package routes

import (
	"encoding/json"
	"fmt"
	"petHealthTool/common"
	"petHealthTool/handles"
	"time"

	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 注册路由
func RegisterAuthRoutes(r *gin.Engine) {
	// 创建路由组，前缀为 /api/auth
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/sendEmailVerifyCode", sendEmailVerifyCode)
		authGroup.POST("/loginByPwd", loginByPwd)
		authGroup.POST("/loginByEmail", loginByEmail)

	}
}

// 账号密码登录
func loginByPwd(c *gin.Context) {

}

// 发送邮箱验证码
func sendEmailVerifyCode(c *gin.Context) {
	rawData, err := c.GetRawData()
	if err != nil {
		common.Fail(c, 400, "参数错误")
		logrus.Error("loginByPwd failed, err:", err)
		return
	}
	// 解析JSON数据
	var data map[string]interface{}
	err = json.Unmarshal(rawData, &data)
	if err != nil {
		common.Fail(c, 400, "参数错误")
		logrus.Error("loginByPwd failed, err:", err)
		return
	}
	// 获取参数
	email, ok := data["email"].(string)
	if !ok {
		common.Fail(c, 400, "参数错误")
		logrus.Error("loginByPwd failed, err:", err)
		return
	}
	emaliHandle := &handles.EmailHandle{}
	// 获取邮箱验证码
	// 生成六位随机数
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成 6 位随机数
	randomNumber := rand.Intn(900000) + 100000 // 范围: 100000 ~ 999999

	// 将随机数转换为字符串
	randomString := fmt.Sprintf("%06d", randomNumber)
	err = emaliHandle.SendVerifyCode(email, randomString)
	if err != nil {
		common.Fail(c, 500, "发送验证码失败")
		logrus.Error("loginByPwd failed, err:", err)
		return
	}
	cfg := common.GetConfig()
	redisClient := common.GetRedisClient(cfg.Redis.Host, cfg.Redis.Pass, cfg.Redis.Db)
	redisClient.Set(email, randomString, time.Minute*5)
	common.Success(c, "发送验证码成功")
}

// 邮箱登录
func loginByEmail(c *gin.Context) {
	var loginByEmailReq common.LoginByEmailReq
	if err := c.ShouldBindJSON(&loginByEmailReq); err != nil {
		common.Fail(c, 400, "参数错误")
		logrus.Error("loginByEmail failed, err:", err)
		return
	}
	authHandle := &handles.AuthHandle{}
	authHandle.LoginByEmail(&loginByEmailReq, c)
}
