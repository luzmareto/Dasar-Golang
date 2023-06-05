package main

import "fmt"

func main() {

	//jika kondisi terpenuhi, maka bisa dicetak
	//jika kondisi tidak terpenuhi, tidak bisa dicetak

	var name = "Luz" //kondisi

	if name == "irfan" {
		fmt.Println("Hello Luz")
	} else if name == "Luz" { //terpenuhi
		fmt.Println("hai kamu")
	} else if name == "luz mar" {
		fmt.Println("hai mas")
	} else {
		fmt.Println("hai, kenalan dong")
	}

	//menghitung jumlah karakter string pada variable
	if length := len(name); length > 5 { //apakah value name lebih dari 5 karakter?
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Sudah benar") //kondisi yang terpenuhi
	}
}
