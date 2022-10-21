package main

import (
	"example/greetings/currency"
	"example/greetings/simpletypes"
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

	// truefalse.Activity()

	// calculator.Activity1()

	currency.Activity()
	fmt.Println("Bye Go Worlda")
}
