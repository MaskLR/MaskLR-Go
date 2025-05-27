package main

import (
    "bufio"
    "log"
    "net/http"
    "os"
    "strings"
    "github.com/MaskLR/MaskLR-Go/internal"
)

func main() {
    // 启动 HTTP 服务
    go func() {
        http.HandleFunc("/", internal.IndexHandler)
        log.Println("MaskLR-Go 启动于 http://127.0.0.1:8080")
        log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
    }()

    // 监听输入 exit 退出
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        if strings.TrimSpace(scanner.Text()) == "exit" {
            log.Println("收到 exit 指令，程序退出。")
            os.Exit(0)
        }
    }
}