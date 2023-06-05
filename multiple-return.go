package main

import "fmt"

func getFullName() (string, string) { //2 return type data string
	return "luz", "mareto"
}

func main() {
	//2 variable penampung untuk return string pada fun getFullName()
	firstName, lastName := getFullName()
	//mencetak return
	fmt.Println(firstName, lastName) //luz mareto
}
