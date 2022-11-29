//testpointer.go
package main

import (
	"fmt"
)

var p *int = new(int) // 创建空指针且赋值

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Printf("%p, %T\n", p, p) //output: 0xc000018090, *int
	fmt.Println(*p)              // 0 全局变量
}

func main() {
	fmt.Printf("%p, %T\n", p, p) //output: 0xc000018090, *int
	p, err := foo()
	fmt.Printf("%p, %T\n", p, p) //output: 0xc000018098, *int
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Printf("%p, %T\n", p, p) //output: 0xc000018098, *int
	fmt.Println(*p)              // 5 局部变量
}
