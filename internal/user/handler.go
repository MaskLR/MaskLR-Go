package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHandler 用户注册接口
func RegisterHandler(c *gin.Context) {
	var req struct {
		Nickname string `json:"nickname" binding:"required,min=3,max=32"`
		Password string `json:"password" binding:"required,min=6,max=64"`
		Email    string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数格式错误"})
		return
	}

	ip := c.ClientIP()
	user, err := RegisterUser(req.Nickname, req.Password, req.Email, ip)
	if err != nil {
		if err == ErrUserExists {
			c.JSON(http.StatusConflict, gin.H{"error": "邮箱已存在"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
		"user": gin.H{
			"id":        user.ID,
			"nickname":  user.Nickname,
			"email":     user.Email,
			"createdAt": user.CreatedAt,
		},
	})
}

// LoginHandler 用户登录接口
func LoginHandler(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	ip := c.ClientIP()
	token, user, err := LoginUserByEmail(req.Email, req.Password, ip)
	if err != nil {
		switch err {
		case ErrUserNotFound, ErrInvalidPassword:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "邮箱或密码错误"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "登录失败"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   token,
		"user": gin.H{
			"id":          user.ID,
			"nickname":    user.Nickname,
			"email":       user.Email,
			"lastLoginAt": user.LastLogin,
			"lastLoginIP": user.LoginIP,
		},
	})
}
