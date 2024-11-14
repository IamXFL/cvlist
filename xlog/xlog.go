package xlog

import (
	"log"
	"os"
	"time"
)

func init() {
	currDate := time.Now().Format("2006-01-02")
	file, err := os.OpenFile(currDate+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// 将日志输出重定向到文件
	log.SetOutput(file)

	// 记录日志
	Info("This is a log message")
}

func Info(str string) {
	str = "[INFO] " + str
	log.Println(str)
}

func InfoF(str string, param ...any) {
	str = "[INFO] " + str
	log.Printf(str, param...)
}

func Error(str string) {
	str = "[Error] " + str
	log.Println(str)
}

func ErrorF(str string, param ...any) {
	str = "[Error] " + str
	log.Printf(str, param...)
}
