package db

import (
	"fmt"
	"log"
	"time"

	"MaskLR-Go/internal/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

// InitMySQL 尝试 Socket 优先连接，失败 fallback 到 TCP
func InitMySQL(socketDSN string) error {
	var err error

	log.Println("尝试通过 Unix Socket 连接 MySQL...")
	DB, err = sqlx.Open("mysql", socketDSN)
	if err == nil && pingDB(DB) == nil {
		log.Println("✅ 成功通过 Socket 连接 MySQL")
		return nil
	}
	log.Printf("❌ Socket 连接失败：%v", err)

	tcpDSN := buildTCPDSN()
	log.Println("尝试通过 TCP 连接 MySQL...")
	DB, err = sqlx.Open("mysql", tcpDSN)
	if err != nil {
		return fmt.Errorf("打开 TCP 连接失败: %w", err)
	}
	if pingErr := pingDB(DB); pingErr != nil {
		return fmt.Errorf("TCP ping 失败: %w", pingErr)
	}
	log.Println("✅ 成功通过 TCP 连接 MySQL")
	return nil
}

func pingDB(db *sqlx.DB) error {
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return db.Ping()
}

func buildTCPDSN() string {
	c := config.Conf.MySQL
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		c.User,
		c.Password,
		c.TCPHost,
		c.DBName,
	)
}
