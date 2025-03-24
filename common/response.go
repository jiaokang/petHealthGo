package common

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

func Fail(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func Error(c *gin.Context, code int, message string, err error) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    err.Error(),
	})
}

func SuccessWithCode(c *gin.Context, code int, data interface{}) {
	c.JSON(code, Response{
		Code:    code,
		Message: "success",
		Data:    data,
	})
}

func ErrorWithDetails(c *gin.Context, code int, message string, err error) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    err.Error(),
	})
}
