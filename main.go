package main

import "fmt"

var String1 = "hello"

var String2 string = "hello"

func WithinFunctionString() {
	String3 := "hello"
	fmt.Println(String3)
}

func main() {
	fmt.Println("Hello Go World")
}
