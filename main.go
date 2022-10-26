package main

import (
	"example/greetings/calculator"
	"example/greetings/currency"
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

	fmt.Println("Bye Go Worlda")
}

func Implemented() {
	calculator.Activity1()
	currency.Activity()
	simpletypes.Activity()
	truefalse.Activity()
}
