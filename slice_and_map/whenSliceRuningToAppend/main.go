package main

import "fmt"

func appendWhenRunning() {
	list := []int64{0, 1, 2, 3, 4, 5}
	for _, val := range list {
		if val > 5 && val < 9 {
			list = append(list, val+1)
		}
		fmt.Println(val)
	}
}

func main() {
	appendWhenRunning()
}

// out put
// 0
// 1
// 2
// 3
// 4
// 5
