package main

import (
	"errors"
	"fmt"
)

// interface error digunakan untuk return yang sukses dan gagal
func Pembagi(nilai int, pemabagi int) (int, error) {
	if pemabagi == 0 {
		return 0, errors.New("pembagi tidak booleh 0") //error
	} else {
		result := nilai / pemabagi
		return result, nil //no errors or return is int
	}
}

func main() {
	hasil, err := Pembagi(28, 0) //karna parameter pembaginya 0, maka returnya adalah error
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error()) //Error pembagi tidak booleh 0
	}
}
