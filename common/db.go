package common

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := "admin:admin2025@tcp(127.0.0.1:3306)/mask_lr?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("数据库无法访问: %v", err)
	}
}
