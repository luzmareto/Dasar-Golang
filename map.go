package main

import (
	"fmt"
	"strings"
)

func main() {
	//mapping is like array

	// short declaration
	person := map[string]string{
		"name":    "luz",
		"address": "kepo",
	}

	// adding index
	person["tittle"] = "programmer"

	//dont print like this
	fmt.Println(person) //map[address:kepo name:luz]

	// do this. make sure there is typo
	fmt.Println(person["name"])    //luz
	fmt.Println(person["address"]) //kepo
	fmt.Println(person["tittle"])  //programmer

	fmt.Println(strings.Repeat("=", 80))

	//long declaration
	var book map[string]string = make(map[string]string)
	book["tiitle"] = "belajar go"
	book["author"] = "luz"
	book["ups"] = "salah"

	delete(book, "ups") //delete index ups

	fmt.Println(book)
}
