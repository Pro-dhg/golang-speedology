package main

import "fmt"

func main() {
	fmt.Println(max(1, 2))
	fmt.Println(swap("world", "hello"))
}

func max(a, b int) int {
	var result int
	if a > b {
		result = a
	} else {
		result = b
	}

	return result
}

func swap(a, b string) (string, string) {
	return b, a
}
