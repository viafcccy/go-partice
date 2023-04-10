package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

func main() {
	var a = make([]int, 100)
	fmt.Println(a)
	fmt.Println(len(a[:10])) // 前闭后开 0-9
	fmt.Println(len(a[10:])) // 10-99
	i := 10
	// 先取出 a 的第 10 个和之后的元素 共 90 个，接到 1 后，再放到 a 的前十个之后
	a = append(a[:i], append([]int{1}, a[i:]...)...) // 在第 i 个位置插入 x
	fmt.Println(a)
	fmt.Println(len(a))
	a = append(a[:i], append([]int{1, 2, 3}, a[i:]...)...) // 在第 i 个位置插入切片
	fmt.Println(a)
	fmt.Println(len(a))

	var b = make([]int, 100)
	b = append(b[:i], []int{1}...) // 在第 i 个位置插入 x
	fmt.Println(b)
	fmt.Println(len(b))
	b = append(b[:i], []int{1, 2, 3}...) // 在第 i 个位置插入切片
	fmt.Println(b)
	fmt.Println(len(b))

	var a2 = []float64{4, 2, 5, 7, 2, 1, 88, 1}
	var a3 = 1 << 20
	fmt.Println(a2)
	fmt.Println(a3)
}

func SortFloat64FastV1(a []float64) {
	// 强制类型转换
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]

	// 以 int 方式给 float64 排序
	sort.Ints(b)
}

func SortFloat64FastV2(a []float64) {
	// 通过 reflect.SliceHeader 更新切片头部信息实现转换
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr

	// 以 int 方式给 float64 排序
	sort.Ints(c)
}
