package pkg

import (
	"log"
	"os"
	"time"
)

func InitLogger() {
	// 强制设置本地时区为操作系统时区
	loc, err := time.LoadLocation("Local")
	if err == nil {
		time.Local = loc
	}

	logFile, err := os.OpenFile("masklr.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("无法创建日志文件，使用默认输出")
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags)
}
