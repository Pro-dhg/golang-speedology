package main

import "fmt"

func main() {

	var asd int = 22
	var qwe *int
	qwe = &asd
	fmt.Printf("aa 变量的地址是 %x", &asd)
	fmt.Println()
	fmt.Printf("ip 变量的指针地址是 %x", qwe)
	fmt.Println()
	fmt.Printf("*ip 变量的值是 %d", *qwe)
	fmt.Println()

	nilTest()

}

func nilTest() {
	var ptr *int
	fmt.Printf("ptr的值为 %x", ptr)
	fmt.Println()

	//空指针判断
	if ptr != nil {
		fmt.Println("不为空")
		fmt.Println(ptr)
	}
	if ptr == nil {
		fmt.Println("为空")
		fmt.Println(ptr)
	}
}
