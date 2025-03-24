package common

// 注册参数
type LoginByPwdReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
