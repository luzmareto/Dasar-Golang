package main

import "fmt"

func main() {
	/*
		variable const never be change and never error if they not using
	*/

	const firstName string = " luz"
	// this will error
	// firstName = "rio"

	const lastName string = " mareto"
	const value = 24

	//Multiple Contant
	const (
		namaDepan    = "luz"
		namaBelakang = "mareto"
		umur         = 24
	)

	fmt.Println(namaDepan)    //luz
	fmt.Println(namaBelakang) //mareto
	fmt.Println(umur)         //24
}
