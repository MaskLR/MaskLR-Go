package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MaskLR/MaskLR-Go/model"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "只支持 POST", http.StatusMethodNotAllowed)
		return
	}

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Username == "" || user.Password == "" {
		http.Error(w, "无效请求", http.StatusBadRequest)
		return
	}

	_, err = main.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		http.Error(w, "注册失败", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "注册成功")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "只支持 POST", http.StatusMethodNotAllowed)
		return
	}

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "无效请求", http.StatusBadRequest)
		return
	}

	row := main.DB.QueryRow("SELECT id FROM users WHERE username = ? AND password = ?", user.Username, user.Password)
	if err := row.Scan(&user.ID); err != nil {
		http.Error(w, "用户名或密码错误", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "登录成功，用户ID: %d", user.ID)
}
