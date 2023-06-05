package main

import (
	"fmt"
	"strings"
)

func main() {

	//STRING

	// long declaration
	var name string
	name = "luz"
	fmt.Println(name)

	//semi long declaration
	var name1 = "mar"
	fmt.Println(name1)

	//short declaration
	name2 := "eti"
	fmt.Println(name2)

	//change value from variable
	name2 = "eto"
	fmt.Println(name2)

	fmt.Println(strings.Repeat("=", 80))

	// INT

	// Long Declaration
	var number3 int
	number3 = 100
	fmt.Println(number3)

	// Semi Short Declaration
	var number4 = 30
	fmt.Println(number4)

	// Semi Short Declaration
	var number5 = 200
	fmt.Println(number5)

	//changes values
	number4 = 4000
	fmt.Println(number4)

	fmt.Println(strings.Repeat("=", 80))

	//Declaration Multiple Variable
	var (
		firstName = "Luz"
		lastName  = "Mareto"
		age       = 24
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)

	// change value from Declaration Multiple Variable
	firstName = "apa"
	fmt.Println(firstName)
}
