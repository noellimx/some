package truefalse

import (
	"fmt"
	"log"
)

// Page 11
var min int = 1
var max int = 100

func Activity() {

	var guess int
	fmt.Println("Enter integer value")

	fmt.Scanln(&guess)

	fmt.Printf("%d\n", guess)

	if !(min <= guess && guess <= max) {
		errMsg := fmt.Sprintf("Invalid guess. Please input number in range [%d,%d].", min, max)

		log.Fatalln(errMsg)
	}
}
