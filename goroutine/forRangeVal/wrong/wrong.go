package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		vals      = []int{1, 2, 3, 4, 5}
		waitGroup sync.WaitGroup
	)

	for _, val := range vals {
		// val := val
		waitGroup.Add(1)
		go func() {
			fmt.Println(val)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}

/*
output:
5
5
5
5
5
*/
