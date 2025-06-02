package middleware

import (
	"net/http"
	"strings"

	"MaskLR-Go/internal/util"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 用于验证 JWT Token 的中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Header 里获取 Authorization 字段
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供授权信息"})
			c.Abort()
			return
		}

		// 支持 "Bearer token" 格式
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "授权格式错误"})
			c.Abort()
			return
		}

		// 验证 token
		claims, err := util.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 token"})
			c.Abort()
			return
		}

		// 将用户信息保存到 context 中，后续处理器可使用
		c.Set("user_id", claims.UserID)
		c.Set("nickname", claims.Nickname)

		c.Next()
	}
}
