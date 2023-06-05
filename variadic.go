package main

import "fmt"

// variable argument(variadic) bisa menginput banyak value dengan cara ...
func sumAll(numbers ...int) int { //reutrn int
	total := 5

	for _, value := range numbers {
		total += value //5 + menjulah tiap value dari parameter numbers
	}
	return total //total adalah int
}

func main() {
	total := sumAll(1, 2, 10, 123) //memasukan value untuk parameter numbers
	fmt.Println(total)             //141

	//memanggil parameter sebagai slice
	slice := []int{10, 20, 12, 202} //var slice type data int, value yg akan dimasukan ke parameter number
	total = sumAll(slice...)        //memanggil var total dengan value func sumAll dan menjadikan slice jadi parameter
	fmt.Println(total)              //249
}
