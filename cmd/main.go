package main

import (
    "log"
    "net/http"
    "os"

    "github.com/MaskLR/MaskLR-Go/config"
    "github.com/MaskLR/MaskLR-Go/pkg"
    "github.com/MaskLR/MaskLR-Go/routes"
)

func main() {
    // 加载配置
	cfg := config.Load()

	// 初始化日志
	pkg.InitLogger()

	// 初始化数据库
	err := pkg.InitDB(cfg)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
		os.Exit(1)
	}

	// 启动服务
	r := routes.SetupRouter()
	log.Println("服务器启动于端口:", cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, r))

}