package user

import "time"

type User struct {
	ID           uint64    `db:"id" json:"id"`
	Nickname     string    `db:"nickname" json:"nickname"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"-"`
	Phone        *string   `db:"phone"`      // 可为空的字段要用指针
	AvatarURL    *string   `db:"avatar_url"` // 可为空
	RegisterIP   string    `db:"register_ip" json:"register_ip"`
	LoginIP      string    `db:"login_ip" json:"login_ip"`
	LastLogin    time.Time `db:"last_login" json:"last_login"`
	IsActive     bool      `db:"is_active" json:"is_active"`
	Role         string    `db:"role" json:"role"` // 角色字段
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
