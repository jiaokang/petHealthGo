package handles

import "github.com/gin-gonic/gin"

// 注册路由
func RegisterAuthRoutes(r *gin.Engine) {
	// 创建路由组，前缀为 /api/auth
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/register", nil)
		authGroup.POST("/login", nil)
		authGroup.GET("/get", nil)
	}
}
