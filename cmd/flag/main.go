package main

import (
	"flag"
	"fmt"
)

var (
	// 定义 flg
	varstr  = flag.String("str", "str", "input str")
	varint  = flag.Int("int", 0, "input int")
	varbool = flag.Bool("bool", false, "input bool")
)

func main() {
	// 解析 flag
	flag.Parse()
	// 使用
	fmt.Println(*varstr)
	fmt.Println(*varint)
	fmt.Println(*varbool)
}
