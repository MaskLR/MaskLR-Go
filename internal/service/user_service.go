package service

import (
	"errors"
	"github.com/MaskLR/MaskLR-Go/internal/repository"
)

func Authenticate(username, password string) (bool, error) {
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return false, err
	}

	if user.Password != password {
		return false, errors.New("密码错误")
	}

	return true, nil
}
