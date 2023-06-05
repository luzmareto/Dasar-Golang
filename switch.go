package main

import "fmt"

func main() {

	//switch case lebih sederhana dari if. biasanya digunakan untuk 1 variable.

	name := "Luz"

	switch name {
	case "Luz":
		fmt.Println("Hello Luz")
		fmt.Println("Hello Luz")
	case "joko":
		fmt.Println("joko")
		fmt.Println("joko")
	default:
		fmt.Println("hai kamu")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("kepanjangan namamu")
	case length < 2:
		fmt.Println("boleh lah")
	default:
		fmt.Println("ok juga")
	}
}
