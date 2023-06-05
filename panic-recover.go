package main

import (
	"fmt"
	"strings"
)

// recover akan menjalankan program yg error
func endApp() {
	message := recover()
	if message != nil {
		//output pertama jika value parameter error true: Error dengan message Aplikasi Error
		fmt.Println("Error dengan message :", message)
	}
	fmt.Println("aplikasi selesai") //output kedua jika value error false
}

func runApp(error bool) { //value error diisi pada func main
	defer endApp() //loging akan diproses pertama jika output false atau sukses
	if error {
		panic("Aplikasi Error") //jika value error true
	}

	fmt.Println("Aplikasi berjalan") //output pertama jika value error false
}

func main() {

	runApp(false)
	fmt.Println(strings.Repeat("=", 80))
	runApp(true)
}
