package main

import "fmt"

type Apapun interface {
}

// interface kosong. tidak memiliki kontrak dan tipe data
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return 2
	} else {
		return "Ups"
	}
}

func main() {
	// memberi value pada kontrak interface kosong
	var data interface{} = Ups(1)
	fmt.Println(data) //1
}
