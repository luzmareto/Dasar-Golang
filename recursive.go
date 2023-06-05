package main

import "fmt"

//recursive adalah func yang memanggil dirinya sendiri di dalam func-nya

// func ini memiliki parameter value dan return int
func factorialRecursive(value int) int { //isi value akan diinput dari func main
	if value == 1 { //jika isi value adalah 1
		return 1 //maka perulangan akan berenti
	} else {
		//5 * 4 * 3 *2 * 1. perkali
		return value * factorialRecursive(value-1) //memanggil func diri sendiri
	}
}

func main() {

	recursive := factorialRecursive(5) //10 adalah isi dari parameter vale
	fmt.Println(recursive)             //3628800
}
