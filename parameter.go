package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println("Hai", firstName, lastName) //value parameter first dan last diisi pada func main
}

func main() {
	//memanggil func sayHello dan memberi value pada 2 parameter
	sayHello("luz", "mareto") //Hai luz mareto
}
