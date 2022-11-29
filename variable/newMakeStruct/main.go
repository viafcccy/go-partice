package main

import "fmt"

// https://stackoverflow.com/questions/13244947/is-there-a-difference-between-new-and-regular-allocation
// 1. new(T) 返回T类型指针，T类型为任意类型，T为结构体时结构体内的指针类型（如：arr）已经初始化
// 2. make(T[slice | chan | map]) 返回值类型，初始化slice、map、chan专用
// 3. T[struct]{} 返回值类型，等于*new(T)
// 4. &T[struct]{} 返回指针类型，等于new(T)

type testStruct struct {
	list  []string
	count int64
}

func main() {
	methodA := new(testStruct)
	// methodB := make(testStruct) // error from compiler: invalid argument: cannot make testStruct; type must be slice, map, or channel
	methodC := testStruct{}
	methodD := &testStruct{}
	fmt.Println(methodA)
	fmt.Println(methodA.list)
	fmt.Println(methodC)
	fmt.Println(methodC.list)
	fmt.Println(methodD)
	fmt.Println(methodD.list)
	fmt.Println(*methodD)
}
