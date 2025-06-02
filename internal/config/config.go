package config

import (
	"os"
)

type Config struct {
	Port   string // 服务端口
	JWTKey string // JWT 密钥
	MySQL  MySQLConfig
}

type MySQLConfig struct {
	SocketDSN string // Unix socket 连接方式
	TCPHost   string // TCP 主机地址（用于 fallback）
	DBName    string // 数据库名
	User      string // 用户名
	Password  string // 密码
}

var Conf *Config

// 加载配置函数，从环境变量读取配置，如果没有则用默认值
func LoadConfig() error {
	Conf = &Config{
		Port:   getEnv("APP_PORT", "8080"),
		JWTKey: getEnv("JWT_SECRET", "replace_with_strong_jwt_key"),

		MySQL: MySQLConfig{
			SocketDSN: getEnv("MYSQL_SOCKET_DSN", "admin:admin2025@unix(/data/data/com.termux/files/usr/var/run/mysqld.sock)/masklr?parseTime=true"),
			TCPHost:   getEnv("MYSQL_TCP_HOST", "192.168.1.5:3306"),
			DBName:    getEnv("MYSQL_DB", "mask_lr"),
			User:      getEnv("MYSQL_USER", "admin"),
			Password:  getEnv("MYSQL_PASSWORD", "admin2025"),
		},
	}
	return nil
}

// 获取环境变量的值，如果不存在则返回默认值
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
