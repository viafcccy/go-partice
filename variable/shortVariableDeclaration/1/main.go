package main

import "fmt"

func test() (string, string) {
	return "newString", "newString"
}

func main() {
	var oldVal = "oldString"
	fmt.Printf("%v, %T\n", &oldVal, oldVal) // 0xc000010250, string
	newVal, oldVal := test()
	fmt.Printf("%v, %T\n", &oldVal, oldVal) // 0xc000010250, string
	fmt.Println(oldVal)                     // newString
	fmt.Println(newVal)                     // newString
}

// 尽量少的覆盖变量
