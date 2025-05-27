package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MaskLR/MaskLR-Go/handler"
)

func main() {
	http.HandleFunc("/register", handler.RegisterHandler)
	http.HandleFunc("/login", handler.LoginHandler)

	fmt.Println("服务启动端口:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
