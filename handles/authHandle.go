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
	}
	md5Pwd := fmt.Sprintf("%x", md5.Sum([]byte(req.Password)))
	if user.Pwd != md5Pwd {
		return nil, fmt.Errorf("password is invalid")
	}
	return map[string]string{
		"name":     user.Name,
		"nickName": user.NickName,
		"token":    "123456",
	}, nil

}
