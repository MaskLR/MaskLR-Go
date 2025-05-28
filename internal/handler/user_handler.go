package handler

import (
	"encoding/json"
	"net/http"
	"github.com/MaskLR/MaskLR-Go/internal/service"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	ok, err := service.Authenticate(req.Username, req.Password)
	if err != nil || !ok {
		http.Error(w, "登录失败", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "登录成功",
	})
}
