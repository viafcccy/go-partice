package main

import "fmt"

// reference: https://iswade.github.io/translate/go_interface/

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof~~"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow~~"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "~~~~~~"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns~~"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}} // interface 是一种类型，[]interface{} 与 []string{} 是等价的；接口是两个字的宽度：示意图如(type, value)
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
