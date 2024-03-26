package main

import (
	"fmt"
	"main.go/serve"
)

func main() {
	fmt.Println("后端项目启动")
	server := serve.NewServer("127.0.0.1", 8888)
	server.Start()
}
