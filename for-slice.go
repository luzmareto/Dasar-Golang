package main

import (
	"fmt"
	"strings"
)

func main() {

	for counter := 1; counter <= 10; counter++ { //value count di cetak maksimal sampai 10
		fmt.Println("perulangan ke", counter)
	}
	//counter++ artinya mencetak sampai 11x. tapi di angka 11 sudah false, jadi kondisinya terpenuhi sampai angka 10

	fmt.Println(strings.Repeat("=", 80))

	//for range:
	slice := []string{"Luz", "Mareto", "tingganya di", "kepo lu"}

	//cara mencetak literasi dan value
	for i, value := range slice { //i akan mencetak urutan angka
		fmt.Println("Index", i, "=", value)
		/*
			output :
			Index 0 = Luz
			Index 1 = Mareto
			Index 2 = tingganya di
			Index 3 = kepo lu
		*/
	}

	fmt.Println(strings.Repeat("=", 80))

	// hanya mencetak value
	for _, value := range slice { //i akan mencetak urutan angka
		fmt.Println(value)
		/*
			Luz
			Mareto
			tingganya di
			kepo lu
		*/
	}

	fmt.Println(strings.Repeat("=", 80))

	// for range slice
	person := make(map[string]string)
	person["name"] = "luz"
	person["tittle"] = "programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
