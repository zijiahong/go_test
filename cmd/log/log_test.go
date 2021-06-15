package main

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	// log 输出文件
	file, err := os.OpenFile("errors1.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFlags(log.Llongfile | log.LstdFlags)
	log.SetPrefix("Info ")

	// 使用
	log.Println("testing ")

}
