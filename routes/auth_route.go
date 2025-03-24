package routes

import (
	"petHealthTool/common"
	"petHealthTool/handles"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 注册路由
func RegisterAuthRoutes(r *gin.Engine) {
	// 创建路由组，前缀为 /api/auth
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/loginByPwd", loginByPwd)
	}
}

// 账号密码登录
func loginByPwd(c *gin.Context) {
	var loginByPwdReq common.LoginByPwdReq
	if err := c.ShouldBindJSON(&loginByPwdReq); err != nil {
		common.Fail(c, 400, "参数错误")
		logrus.Error("loginByPwd failed, err:", err)
		return
	}
	authHandle := &handles.AuthHandle{}
	result, err := authHandle.LoginByPwd(&loginByPwdReq)
	if err != nil {
		common.Fail(c, 400, err.Error())
		logrus.Error("loginByPwd failed, err:", err)
		return
	}
	common.Success(c, result)
}
