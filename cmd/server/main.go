package main

import (
	"log"
	"os"

	"MaskLR-Go/internal/config"
	"MaskLR-Go/internal/db"
	"MaskLR-Go/internal/router"
)

func main() {

	// 打开 log 文件
	logFile, err := os.OpenFile("masklr.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("无法打开日志文件: %v", err)
	}
	defer logFile.Close()

	// 设置标准日志输出到文件
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 加载配置
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化数据库连接
	if err := db.InitMySQL(config.Conf.MySQL.SocketDSN); err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 注册路由
	r := router.SetupRouter()

	// 启动 HTTP 服务
	addr := ":" + config.Conf.Port
	log.Printf("服务启动成功，监听地址 http://localhost%s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
