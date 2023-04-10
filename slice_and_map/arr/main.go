package main

import "fmt"

func main() {
	var arrA = [1]int{1}
	var arrB = [2]int{2}
	// arrB = arrA
	fmt.Println(arrB)
	fmt.Println(arrA)
	// error: cannot use arrA (variable of type [1]int) as [2]int value in assignment
	// 数组类型包含个数这个信息
}
