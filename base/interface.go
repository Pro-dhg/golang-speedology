package main

import "fmt"

type Phone interface {
	call()
	stop()
	start()
}

type IPhone struct {
}

func (iPhone IPhone) stop() {
	fmt.Println("stop")
}
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone ,I can call you")
}
func main() {
	phone := new(IPhone)
	phone.call()
	phone.stop()
}
