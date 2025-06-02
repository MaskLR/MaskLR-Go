package router

import (
	"github.com/gin-gonic/gin"

	"MaskLR-Go/internal/middleware"
	"MaskLR-Go/internal/user"
)

// SetupRouter 初始化 Gin 路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 用户相关接口（公开）
	r.POST("/register", user.RegisterHandler)
	r.POST("/login", user.LoginHandler)

	// 用户相关接口（需要登录）
	auth := r.Group("/user")
	auth.Use(middleware.AuthMiddleware())
	//	{
	//		auth.GET("/profile", user.ProfileHandler) // 示例接口，返回当前登录用户信息
	// 更多受保护接口可在这里添加
	//	}

	return r
}
