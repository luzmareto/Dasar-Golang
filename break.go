package main

import "fmt"

func main() {
	//value i adalah 0. akan dicetak maksimal sampai 9
	for i := 0; i < 10; i++ {
		if i == 5 {
			break //berenti pada perulangan ke-5
		}
		fmt.Println("perulangan  ke", i) //0-4
	}
}
