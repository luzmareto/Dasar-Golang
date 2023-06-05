package main

import "fmt"

func main() {
	var name1 = "luz"
	var name2 = "mareto"

	//comparation value name1 and name2
	var result bool = name1 == name2 //is value ame1 and name2 ?
	fmt.Println(result)              //false

	var result1 bool = name2 < name1 //is character name2 more than name1 ?
	fmt.Println(result1)             //false
}
