package internal

import (
    "fmt"
    "net/http"
    "github.com/MaskLR/MaskLR-Go/internal"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    msg := internal.TestDB()
    fmt.Fprintf(w, "你好！数据库连接测试结果：%s\n", msg)
}
