package main

import "fmt"

/*
return digunakan jika dalam sebuah func ada deklarasi data kembalian
value return bebas, asalkan type datanya sesuai dengan data kembalian pada func
*/

func getHello(name string) string {
	if name == "" {
		return "hello Bro"
	} else {
		return "hello " + name
	}
}

func main() {
	result := getHello("jefri") //jefri adalah return dari func getHello
	fmt.Println(result)         //hello jefri

	//string kosong adalah return dari func getHello
	fmt.Println(getHello("")) //hello bro
}
