package main

import (
	"io"
	"log"
	"os"
)

var (
	Info *log.Logger
)

func init() {
	logFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}
	// 设置 log 输出
	log.SetOutput(io.MultiWriter(os.Stderr, logFile))
	// 设置 log 前缀
	log.SetPrefix("log ")
	// 设置 log 输出包含的信息
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
func main() {
	log.Println("testing")
}
