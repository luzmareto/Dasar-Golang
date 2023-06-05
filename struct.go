package main

import "fmt"

// recommended use upperCase in name and field struct
type Customer struct {
	Name, Address string
	Age           int
}

// func sayHi have struct Customer and parameter name
func (customer Customer) sayHi(name string) { //parameter name field at func main
	//irvan from fun main. luzmareto from field var Luz Customer in func main
	fmt.Println("Hello", name, "My name is", customer.Name) //Hello Irvan My name is Luzmareto
}

func (c Customer) sayHuu() {
	//luzmareto is from field Customer c.Name and have value from var luz Customer
	fmt.Println("Huuuuu from", c.Name) //Huuuuu from Luzmareto
}

func main() {
	//var luz is type data struct Customer
	var luz Customer

	//field struct Customer
	luz.Name = "Luzmareto"
	luz.Address = "Indonesia"
	luz.Age = 24

	luz.sayHi("Irvan")

	luz.sayHuu()

	// fmt.Println(luz)         //{Luzmareto Indonesia 24}
	// fmt.Println(luz.Name)    //Luzmareto
	// fmt.Println(luz.Address) //indonesia
	// fmt.Println(luz.Age)     //24
}
