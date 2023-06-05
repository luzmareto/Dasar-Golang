package main

import (
	"fmt"
	"strings"
)

func main() {

	//skip angka genap
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("perulangan ke", i) //perulangan ke 1,3,5,7,9
	}

	fmt.Println(strings.Repeat("=", 80))

	//skip angka ganjil
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println("perulangan ke", i) //perulangan ke 0,2,4,6,8
	}

}
