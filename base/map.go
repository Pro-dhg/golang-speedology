package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["apple"] = 1
	m["banana"] = 2
	m["orange"] = 3

	delete(m, "apple")

	for k, v := range m {
		fmt.Printf("key:%s value:%d \n", k, v)
	}
}
