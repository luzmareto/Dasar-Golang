package main

import "fmt"

//defer adalah proses pemberentian tugas, dan jika programnya ada yang error, defer akan menjalankan program yg sukses

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp(value int) {
	defer logging()

	fmt.Println("Run Appp")
	result := 10 / value
	fmt.Println("result", result)
	// logging()
}

func main() {
	runApp(10)
}
