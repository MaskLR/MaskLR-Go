package router

import (
	"MaskLR-Go/internal/middleware"
	"MaskLR-Go/internal/user"

	"github.com/gin-gonic/gin"
)

// SetupRouter 注册路由
func SetupRouter(r *gin.Engine) {
	// 用户相关接口（公开）
	r.POST("/register", user.RegisterHandler)
	r.POST("/login", user.LoginHandler)

	// 用户相关接口（需要登录）
	auth := r.Group("/user")
	auth.Use(middleware.AuthMiddleware())
	{
		// 示例接口（已注释）
		// auth.GET("/profile", user.ProfileHandler)
	}
}
