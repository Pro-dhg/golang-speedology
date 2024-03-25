package main

import "fmt"

var c, d int
var ( //这种因式分解关键字的写法一般用于声明全局变量
	e int
	f bool
)

var g, h int = 1, 2
var i, j = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//k,l := 123,"hello"

func main() {
	k, l := 123, "hello"
	var a = true
	fmt.Println(a)
	a = false
	fmt.Println(a)

	b := true
	fmt.Println(b)

	var x, y, z = 1, 2, 3
	fmt.Println(x, y, z)

	fmt.Println(c, d, e, f, g, h, i, j, k, l)
	fmt.Println(&c)
	as := &c
	fmt.Println(as)

}
