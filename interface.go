package main

import "fmt"

// interface is like contract
type HasName interface { //HasName have contrcat Getname() string
	GetName() string //string is return
}

// this func have parameter hasName with type data interface Hasname
func SayHello(hasName HasName) {
	//parameter hasName field value at func main and inside GetName()
	fmt.Println("hello", hasName.GetName()) //hello luz
}

type Person struct {
	Name string
}

// this func GetName() from InterfaceHasname
func (person Person) GetName() string { //return string
	return person.Name //give value from func main
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	//luz is struct Person
	var luz Person
	//give value for field name at struct Person
	luz.Name = "luz"

	//call func SayHello and give parameter from var Luz Person
	SayHello(luz)

	cat := Animal{
		Name: "push",
	}
	SayHello(cat) //hello push
}
