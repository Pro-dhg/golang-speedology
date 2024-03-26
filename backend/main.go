package main

import (
	"fmt"
)

func main() {
	fmt.Println("后端项目启动")
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
