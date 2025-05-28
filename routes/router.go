package routes

import (
	"github.com/MaskLR/MaskLR-Go/internal/handler"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/login", handler.LoginHandler)
	return mux
}
