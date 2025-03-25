package common

// 密码
type LoginByPwdReq struct {
	EmailOrPhone string `json:"emailOrPhone"`
	Password     string `json:"password"`
}

// 邮箱登录参数
type LoginByEmailReq struct {
	Email      string `json:"email"`
	VerifyCode string `json:"verifyCode"`
}
