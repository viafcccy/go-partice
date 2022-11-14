package main

import (
	"fmt"
)

/*
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
*/

type StringObj struct {
	strings []string
}

func (o *StringObj) setStingSlice(strings []string) {
	o.strings = strings
}

func main() {
	var o StringObj
	var paramStrings = []string{"1", "2"}
	o.setStingSlice(paramStrings)
	fmt.Println(o)
	fmt.Println(paramStrings)
	o.strings[0] = "0" // 修改对象 o 时修改了，参数 paramStrings
	fmt.Println(paramStrings)
}

/*
output:
{[1 2]}
[1 2]
[0 2]
*/
