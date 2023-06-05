package main

import "fmt"

// parameter name
func getGoodBye(name string) string { //harus ada return dengan tipedata string
	return "Good Bye " + name //name akan di isi pada func main
}

func main() {
	//var sayGoodBye valuenya adalah func getGoodBye
	sayGoodBye := getGoodBye

	//karna pada func getGoodBye ada return string, maka harus membuat var penampung lagi
	result := sayGoodBye("luz") //luz adalah value sekaligus return
	fmt.Println(result)
}
