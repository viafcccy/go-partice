package main

import "fmt"

func main() {
	// 一个数组内包含了五个空数组
	var times [5][0]int
	for _, val := range times {
		fmt.Println(val)
		fmt.Println("hello")
	}
	// var times [5][0]int
	// for range times {
	// 	fmt.Println("hello")
	// }
	// output:
	// 	[]
	// hello
	// []
	// hello
	// []
	// hello
	// []
	// hello
	// []
	// hello
}
