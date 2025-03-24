package handles

import (
	"crypto/md5"
	"fmt"
	"petHealthTool/common"
	"petHealthTool/repository"

	"github.com/sirupsen/logrus"
)

type AuthHandle struct {
}

func (a *AuthHandle) LoginByPwd(req *common.LoginByPwdReq) (map[string]string, error) {
	logrus.Info("loginByPwdReq:", req)
	usersRepo := &repository.UsersRepo{}
	user, err := usersRepo.GetUserByName(req.Username)
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
		"name":     user.Name,
		"nickName": user.NickName,
		"token":    token,
	}, nil

}
