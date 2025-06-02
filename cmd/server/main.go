package main

import (
	"log"
	"os"

	"MaskLR-Go/internal/config"
	"MaskLR-Go/internal/db"
	"MaskLR-Go/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 设置 Gin 为生产模式
	gin.SetMode(gin.ReleaseMode)

	// 打开日志文件
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

	// 创建 Gin 引擎并重定向日志
	r := gin.New()
	r.Use(gin.LoggerWithWriter(logFile), gin.Recovery())

	// 注册路由
	router.SetupRouter(r)

	// 启动 HTTP 服务
	addr := ":" + config.Conf.Port
	log.Printf("服务启动成功，监听地址 http://localhost%s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
