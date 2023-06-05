package main

import (
	"fmt"
	"strings"
)

func main() {
	// array start from zero

	var names [3]string

	names[0] = "luz" //array 0 = 1
	names[1] = "mar" //array 1 = 2
	names[2] = "eto" //array 2 = 3

	fmt.Println(names[0]) //luz
	fmt.Println(names[1]) //mar
	fmt.Println(names[2]) //eto

	fmt.Println(strings.Repeat("=", 80))

	//short array
	var values = [3]int{
		90,
		95,
		80,
	}

	//print all values
	fmt.Println(values) // [90 95 80]

	//print specifik values
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
}
