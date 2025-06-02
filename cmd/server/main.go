package main

import (
	"log"

	"MaskLR-Go/internal/config"
	"MaskLR-Go/internal/db"
	"MaskLR-Go/internal/router"
)

func main() {
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
