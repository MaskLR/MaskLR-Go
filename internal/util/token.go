package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 建议将这个密钥放到 config 或环境变量中
var jwtSecret = []byte("YourSuperSecretKey") // 替换成你的安全密钥

// Claims 是 JWT 的载荷结构
type Claims struct {
	UserID   uint64 `json:"user_id"`
	Nickname string `json:"nickname"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT token
func GenerateToken(userID uint64, nickname string) (string, error) {
	expiration := time.Now().Add(7 * 24 * time.Hour) // Token 有效期：7天
	claims := Claims{
		UserID:   userID,
		Nickname: nickname,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 验证并解析 JWT token
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
