package main

import "fmt"

var String1 = "hello"

var String2 string = "hello"

func WithinFunctionString() {
	String3 := "hello"
	fmt.Println(String3)
}

func FooTyping() {

	int1 := 24

	fmt.Printf("[int1:type]%T\n", int1)

}
func main() {
	fmt.Println("Hello Go World")

	FooTyping()

	fmt.Println("Bye Go World")

}
