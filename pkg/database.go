package pkg

import (
	"fmt"
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/MaskLR/MaskLR-Go/config"

)

var DB *sql.DB

func InitDB(cfg config.Config) error {
	dsn := fmt.Sprintf("%s:%s@unix(%s)/%s?parseTime=True&charset=utf8mb4",
		cfg.DBUser, cfg.DBPass, cfg.DBSocket, cfg.DBName)

	var err error
	DB, err = sql.Open("mysql",dsn)
	if err != nil {
		return fmt.Errorf("数据库连接失败: %v", err)
	}

	log.Println("数据库连接成功")
	return nil
}
