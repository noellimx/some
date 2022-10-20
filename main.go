package main

import (
	"example/greetings/simpletypes"
	"example/greetings/truefalse"
	"fmt"
)

var String1 = "hello"

var String2 string = "hello"

func WithinFunctionString() {
	String3 := "hello"
	fmt.Println(String3)
}

func main() {
	fmt.Println("Hello Go World")

	simpletypes.FooTyping()
	simpletypes.Activity()

	truefalse.Activity()

	fmt.Println("Bye Go World")

}
