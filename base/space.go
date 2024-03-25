package main

import "fmt"

const name string = "duange"

func main() {
	var x int
	const Pi float64 = 3.14159265358979323846

	fmt.Println(x)
	fmt.Println(Pi)
	fmt.Println(name)
	result := 1 + 1
	fmt.Println(result)

	var stockCode = 123
	var endDate = "2024-03-25"
	var url = "Code=%d endDate=%s"
	target := fmt.Sprintf(url, stockCode, endDate)
	fmt.Println(target)
	fmt.Printf(url, stockCode, endDate)
}
