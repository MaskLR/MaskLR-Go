package repository

import (
	"database/sql"
	"errors"
	"github.com/MaskLR/MaskLR-Go/internal/model"
	"github.com/MaskLR/MaskLR-Go/pkg"
)

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	query := "SELECT id, username, password FROM users WHERE username = ? LIMIT 1"

	row := pkg.DB.QueryRow(query, username)
	err := row.Scan(&user.ID, &user.Username, &user.Password)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil // 用户不存在
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}
