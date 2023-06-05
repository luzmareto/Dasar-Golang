package main

import (
	"fmt"
	"strings"
)

func main() {
	//the operation is like  + - * /

	var result = 10 + 10
	fmt.Println(result) //20

	var a = 100
	var b = 20
	var c = a - b
	fmt.Println(c) //80

	fmt.Println(strings.Repeat("=", 80))

	var i = 10
	i += 10        //10 + 10
	fmt.Println(i) //20

	i++            //20 + 1
	fmt.Println(i) //21
}
