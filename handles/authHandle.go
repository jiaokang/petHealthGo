package handles

import (
	"crypto/md5"
	"errors"
	"fmt"
	"petHealthTool/common"
	"petHealthTool/models"
	"petHealthTool/repository"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthHandle struct {
}

// 账号密码登录
func (a *AuthHandle) LoginByPwd(req *common.LoginByPwdReq) (map[string]string, error) {
	logrus.Info("loginByPwdReq:", req)
	usersRepo := &repository.UsersRepo{}
	user, err := usersRepo.GetUserByName(req.EmailOrPhone)
	if err != nil {
		logrus.Error("loginByPwd failed, err:", err)
		return nil, fmt.Errorf("user is not exist")
	}
	md5Pwd := fmt.Sprintf("%x", md5.Sum([]byte(req.Password)))
	if user.Pwd != md5Pwd {
		return nil, fmt.Errorf("password is invalid")
	}
	// 生成token
	jwt := &common.Jwt{}
	token, err := jwt.CreateToken(user.ID)
	if err != nil {
		logrus.Error("loginByPwd failed, err:", err)
		return nil, fmt.Errorf("create token failed")
	}
	return map[string]string{
		"nickName": user.NickName,
		"token":    token,
	}, nil
}

// 邮箱登录
func (a *AuthHandle) LoginByEmail(req *common.LoginByEmailReq, c *gin.Context) {
	logrus.Info("loginByEmailReq:", req)

	// 获取配置和 Redis 客户端
	cfg := common.GetConfig()
	redisClient := common.GetRedisClient(cfg.Redis.Host, cfg.Redis.Pass, cfg.Redis.Db)

	// 验证验证码
	if err := validateVerificationCode(redisClient, req.Email, req.VerifyCode); err != nil {
		logrus.Error("verification code error:", err)
		common.Fail(c, 400, err.Error())
		return
	}

	// 查询或创建用户
	usersRepo := &repository.UsersRepo{}
	user, err := usersRepo.GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user, err = createUser(usersRepo, req.Email)
			if err != nil {
				logrus.Error("create user failed:", err)
				common.Fail(c, 500, "注册用户失败")
				return
			}
		} else {
			logrus.Error("get user failed:", err)
			common.Fail(c, 500, "系统错误")
			return
		}
	}

	// 生成 Token
	token, err := generateToken(user.ID)
	if err != nil {
		logrus.Error("create token failed:", err)
		common.Fail(c, 500, "生成token失败")
		return
	}

	// 返回成功响应
	common.Success(c, map[string]string{
		"nickName": user.NickName,
		"token":    token,
	})
}

// 验证验证码
func validateVerificationCode(redisClient *common.RedisClient, email, verifyCode string) error {
	storedCode, err := redisClient.Get(email)
	if err != nil {
		if err == redis.Nil {
			return fmt.Errorf("验证码已过期")
		}
		return fmt.Errorf("系统错误")
	}
	if storedCode != verifyCode {
		return fmt.Errorf("验证码错误")
	}
	redisClient.Del(email)
	return nil
}

// 创建用户
func createUser(usersRepo *repository.UsersRepo, email string) (*models.Users, error) {
	user := &models.Users{
		NickName: email[:strings.Index(email, "@")], // 使用邮箱前缀作为昵称
		Phone:    "",                                // 手机号为空
		Email:    email,
		Address:  "",
		Pwd:      "",
	}
	if err := usersRepo.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

// 生成 Token
func generateToken(userID uint) (string, error) {
	jwt := &common.Jwt{}
	return jwt.CreateToken(userID)
}
