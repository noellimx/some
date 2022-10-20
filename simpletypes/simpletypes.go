package simpletypes

import (
	"fmt"

	"example/greetings/utils"
)

// Page 8

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

/*


Create a Go application that has the following features.
1. Create a basic Go program using the standard template
2. Declare the following variables
● text of “The following is the account information.”
● first name “Luke”
● last name “Skywalkter”
● age of 20 yrs old
● weight of 73.0 kg
● height of 1.72 m
● remaining credits of $123.55
● account name of “admin”
● account password of “password”
● subscribed user as “true”
3. Use the package fmt to print out the values and types of the variables declared in part 2.

*/

func Activity() {

	var (
		dialog    string = "The following is the account information."
		nameFirst string = "Luke"
		nameLast  string = "Skywalker"

		age     int     = 20
		weight  float64 = 73.0
		height  float64 = 1.72
		credits float64 = 123.55

		accountName     string = "admin"
		accountPassword string = "password"

		isSubscribed bool = true
	)

	fmt.Printf("%s\n", dialog)
	fmt.Printf("Name: %s %s\n", nameFirst, nameLast)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Weight: %.2f\n", weight)
	fmt.Printf("Height: %.2f\n", height)
	fmt.Printf("Credits: %.2f\n", credits)
	fmt.Printf("Account name: %s\n", accountName)
	fmt.Printf("Account password: %s\n", accountPassword)
	fmt.Printf("Is Suscribed: %t\n", isSubscribed)
}
