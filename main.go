package main

import (
    "log"
    "net/http"
    "github.com/MaskLR/MaskLR-Go/internal"
)

func main() {
    http.HandleFunc("/", internal.IndexHandler)
    log.Println("MaskLR-Go 启动于 http://127.0.0.1:8088")
    log.Fatal(http.ListenAndServe("127.0.0.1:8088", nil))
}