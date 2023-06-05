package main

import "fmt"

// 3 return yang tipedatanya string semua
func getFullName2() (firstName, middelName, lastName string) {
	//memberi value pada return
	firstName = "luz"
	middelName = "mar"
	lastName = "eto"

	return
}

func main() {
	//3 var yang funsginya sesuai urutan return
	a, b, c := getFullName2()
	fmt.Println(a, b, c) //luz mar eto
}
