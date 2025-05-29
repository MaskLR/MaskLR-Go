package config

import (
	"os"
)

type Config struct {
	AppPort string
	DBHost  string
	DBPort  string
	DBSocket string
	DBUser  string
	DBPass  string
	DBName  string
}

func Load() Config {
	return Config{
		AppPort: getEnv("APP_PORT", "8080"),
		DBHost:  getEnv("DB_HOST", "127.0.0.1"),
		DBPort:  getEnv("DB_PORT", "3306"),
		DBSocket:  getEnv("DB_SOCKET", "/data/data/com.termux/files/usr/var/run/mysqld.sock"),
		DBUser:  getEnv("DB_USER", "admin"),
		DBPass:  getEnv("DB_PASS", "admin2025"),
		DBName:  getEnv("DB_NAME", "mask_lr"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
