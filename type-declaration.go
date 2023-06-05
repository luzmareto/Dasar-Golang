package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var notKtpLuz NoKTP = "12321"
	var marriedStatus Married = true
	fmt.Println(notKtpLuz)
	fmt.Println(marriedStatus)
}
