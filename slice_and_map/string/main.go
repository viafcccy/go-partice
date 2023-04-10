package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

// unicode: 字符集，使用固定的 4 个 byte 表示地球几乎全部字符，底层为 []rune = []uint32
// utf-8 / utf-16：unicode 字符集的编码格式 string
// 字节数组 []byte = []uint8：对应 unicode 字符集中的每一个字节

func main() {
	testString := "你"
	fmt.Println(len(testString))
	forOnString(testString, func(i int, r rune) {
		fmt.Println(i)
		fmt.Println(r)
	})

	stringFrom := bytes2str([]byte{97})
	fmt.Println(stringFrom)
}

// 将 string 拆解为底层 rune
func forOnString(s string, forBody func(i int, r rune)) {
	for i := 0; len(s) > 0; {
		r, size := utf8.DecodeRuneInString(s)
		forBody(i, r)
		s = s[size:]
		i += size
	}
}

// 通过构造底层数据 bytes 转 str
func bytes2str(s []byte) (p string) {
	data := make([]byte, len(s))
	for i, c := range s {
		data[i] = c
	}
	fmt.Println(data)

	// reflect.StringHeader 为 string 底层结构体
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&p))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(s)
	fmt.Println(hdr)

	return p
}
