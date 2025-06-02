package user

import (
	"database/sql"
	"errors"
	"time"

	"MaskLR-Go/internal/db"
	"MaskLR-Go/internal/util"
)

var (
	ErrUserExists      = errors.New("用户已存在")
	ErrUserNotFound    = errors.New("用户不存在")
	ErrInvalidPassword = errors.New("密码错误")
)

// RegisterUser 注册新用户
func RegisterUser(nickname, password, email, ip string) (*User, error) {
	// 1. 检查用户名是否已存在
	var exists int
	err := db.DB.Get(&exists, "SELECT COUNT(*) FROM users WHERE nickname = ?", nickname)
	if err != nil {
		return nil, err
	}
	if exists > 0 {
		return nil, ErrUserExists
	}

	// 2. 哈希密码
	hash, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// 3. 插入新用户
	now := time.Now()
	res, err := db.DB.Exec(`
		INSERT INTO users (nickname, password_hash, email, register_ip, login_ip, last_login, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		nickname, hash, email, ip, ip, now, now, now)
	if err != nil {
		return nil, err
	}

	// 4. 获取新用户 ID
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &User{
		ID:           uint64(id),
		Nickname:     nickname,
		PasswordHash: hash,
		Email:        email,
		RegisterIP:   ip,
		LoginIP:      ip,
		LastLogin:    now,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}

// LoginUser 用户登录，验证密码并更新最后登录时间/IP，返回 token
func LoginUserByEmail(email, password, ip string) (string, *User, error) {
	var u User
	err := db.DB.Get(&u, "SELECT * FROM users WHERE email = ?", email)
	if err == sql.ErrNoRows {
		return "", nil, ErrUserNotFound
	} else if err != nil {
		return "", nil, err
	}

	if !util.CheckPasswordHash(password, u.PasswordHash) {
		return "", nil, ErrInvalidPassword
	}

	// 更新登录时间与 IP
	now := time.Now()
	_, err = db.DB.Exec(`
		UPDATE users SET login_ip = ?, last_login = ?, updated_at = ? WHERE id = ?`,
		ip, now, now, u.ID)
	if err != nil {
		return "", nil, err
	}

	// 更新结构体信息
	u.LoginIP = ip
	u.LastLogin = now
	u.UpdatedAt = now

	// 生成 JWT token
	token, err := util.GenerateToken(u.ID, u.Nickname)
	if err != nil {
		return "", nil, err
	}

	return token, &u, nil
}
