package internal

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func TestDB() string {
    dsn := "admin:admin2025@tcp(127.0.0.1:3306)/masklr_go?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return "数据库连接失败: " + err.Error()
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        return "数据库无法 Ping: " + err.Error()
    }
    return "数据库连接成功"
}
