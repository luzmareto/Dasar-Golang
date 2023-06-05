package main

import "fmt"

func sayHello() {
	fmt.Println("hello world")
}

func main() {
	//var i akan mencetak sebanyak 5x
	for i := 0; i < 5; i++ {
		sayHello() //yang dicetak adalah hello wolrd dari func sayHello
	}
}
