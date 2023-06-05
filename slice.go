package main

import "fmt"

func main() {
	//slice is part of array. but, value slice can change and unessecary initialized

	//[...]string is use slice string without initial value
	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	//call index of 4 - 7
	var slice1 = months[4:7] //[mei juni juli]
	fmt.Println(slice1)

	//count value
	fmt.Println(len(slice1)) //3 (mei juni juli)

	//count from first index(4) till finish. count mei - desember
	fmt.Println(cap(slice1)) //9

	//change value index 5 wich is juni
	months[5] = "ubah"
	fmt.Println(slice1) //[mei ubah juli]

	//changes values slice1.
	slice1[0] = "gatau" //index 0 is mei. remember, this data from slice1
	fmt.Println(slice1) //[gatau ubah juli]
}
