package main

import (
	"fmt"
	"time"
)

func main() {
	// 函数在一个新创建的goroutine中运行
	go spinner(100 * time.Millisecond)

	// main goroutine
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

	// 主函数返回。主函数返回时，所有的goroutine都会被直接打断，程序退出。
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
