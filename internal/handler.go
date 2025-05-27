package internal

import (
    "fmt"
    "net/http"
    "github.com/MaskLR/MaskLR-Go/db"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    msg := db.TestDB()
    fmt.Fprintf(w, "你好！数据库连接测试结果：%s\n", msg)
}
