package main

import "fmt"

func main() {
	var numbers [5]int
	var numbers2 = [5]int{1, 2, 3, 4, 5}
	numbers3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println(numbers2)
	fmt.Println(numbers3)

	var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)
	balance[4] = 23.1
	fmt.Println(balance)

	fmt.Println("--------------")

	var n [10]int
	for i := 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j := 0; j < 10; j++ {
		fmt.Printf("%d %d", j, n[j])
		fmt.Println()
	}

}
