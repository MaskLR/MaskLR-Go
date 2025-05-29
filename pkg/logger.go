package pkg

import (
	"log"
	"os"
)

func InitLogger() {
	logFile, err := os.OpenFile("masklr.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("无法创建日志文件，使用默认输出")
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags)
}
