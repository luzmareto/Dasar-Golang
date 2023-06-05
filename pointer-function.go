package main

import "fmt"

// struct Address diisi pada func main
type Address struct {
	City, Province, Country string
}

// parameter address adalah pointer struct address yang fungsinya akan menduplikat field
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "indonesia" //default field country
}

func main() {
	//alamat adalah var penampung untuk memanggil struct Address
	var alamat = Address{
		City:     "subang",
		Province: "jawa barat",
		Country:  "",
	}

	//var ini menduplikat field struct addres, dan value fieldnya dari alamat
	var alamatPointer *Address = &alamat

	//memanggil func ChangeCountryToIndonesia dan memberi value parameter address menggunakan var alamat pointer
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat) //{subang jawa barat indonesia}
}
