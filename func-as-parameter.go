package main

import "fmt"

// membuat alias dengan type data func
type Filter func(string) string //Filter adalah alias yg isinya parameter string dan harus ada return string

// func sayHelloFilter dengan parameter name type data string dan parameter filter dengan typedata alias Filter
func sayHelloFilter(name string, filter Filter) {
	//var penampung untuk parameter filter
	nameFiltered := filter(name) //name adalah calon parameter yg akan diisi string

	//mencetak hello + value name yang akan diisi pada func main
	fmt.Println("hello", nameFiltered)
}

// melakukan filtering pada kata tertentu
func spamFilter(name string) string { //parameter name harus melakukan return string
	//jika value name pada func main adalah anjing, tolong disensor menjadi Anj*ng
	if name == "Anjing" {
		return "Anj*ng"
	} else {
		return name //name adalah string
	}
}

func main() {
	/*
		CARA KERJA :
		1. memanggil func sayHelloFilter dan melakukan penyaringan value
		2. luz dikirim ke func sayHelloFilter, lalu dimasukan ke nameFiltered dan dicetak para parameter name di alias filter
		3. Setelah itu, value "luz" akan di cek melalui func spamFilter, dan dia lulus sensor
		4. namun pada value "Anjing tidak lulus sensor, maka outputnya akan menjadi Anj*ng"
	*/

	sayHelloFilter("Luz", spamFilter)    //hello Luz
	sayHelloFilter("Anjing", spamFilter) //hello Anj*ng
}
