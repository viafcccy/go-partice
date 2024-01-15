package main

import (
	"fmt"
)

func main() {
	// forever := make(chan bool)

	// // 使用 goroutine 运行阻塞函数
	// go blockForever(forever)

	// 继续其他操作...
	fmt.Println("The main function continues to execute...")

	// 等待用户输入，防止程序立即退出
	var input string
	fmt.Scanln(&input)
	fmt.Println("Exited")
}

// blockForever 将会阻塞，因为它在等待从通道接收值
func blockForever(forever chan bool) {
	fmt.Println("Blocking function started...")
	<-forever
	fmt.Println("This line will never execute.")
}
