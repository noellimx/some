package simpletypes

import (
	"fmt"

	"example/greetings/utils"
)

func FooTyping() {

	int1 := 24
	typeInt1Have := fmt.Sprintf("%T", int1)

	utils.Assert(typeInt1Have == "int")

	intLong1 := int64(int1)
	typeIntLong1Have := fmt.Sprintf("%T", intLong1)

	utils.Assert(typeIntLong1Have == "int64")

	var initIntExpects int

	utils.Assert(initIntExpects == 0)

	var age int // warning: should combine declaration and assignment
	age = 10

	utils.Assert(age == 10)

	const constage int = 1 // cannot reassign value.

	var string1 string = "some"
	typeString1Have := fmt.Sprintf("%T", string1)
	utils.Assert(typeString1Have == "string")

}
