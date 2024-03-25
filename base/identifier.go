package main

import (
	"fmt"
	"unsafe"
)

const (
	a  = "abc"
	b  = len(a)
	cc = unsafe.Sizeof(a)
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	area = LENGTH * WIDTH
	fmt.Printf("面积为 %d", area)
	fmt.Println()
	fmt.Println(a, b, cc)

	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
