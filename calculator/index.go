package calculator

import "fmt"

type typeBinaryOperator func(a int, b int) int

type typeOpMap map[string]typeBinaryOperator

var operations typeOpMap

func Activity() {

	var (
		a int
		b int

		op string
	)
	operations = make(typeOpMap)
	operations["+"] = func(a int, b int) int {

		return a + b
	}

	operations["-"] = func(a int, b int) int {

		return a - b
	}
	operations["/"] = func(a int, b int) int {

		return a / b
	}

	operations["*"] = func(a int, b int) int {

		return a * b
	}

	fmt.Println("Your first number")

	fmt.Scanln(&a)

	fmt.Println("Your op")

	fmt.Scanln(&op)

	fmt.Println("Your second number")

	fmt.Scanln(&b)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[Recovered] Oops something went wrong with calculating...")
		}
	}()

	var result int = operations[op](a, b)

	fmt.Println("Result")
	fmt.Println(result)

}
