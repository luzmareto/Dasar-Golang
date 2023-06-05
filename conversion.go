package main

import (
	"fmt"
)

func main() {
	var nilai32 int32 = 10000000

	// change type data nilai32
	var nilai64 int64 = int64(nilai32) //value nilai32 still same

	//when var cannot convert, value will down
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32) //10000000
	fmt.Println(nilai64) //10000000
	fmt.Println(nilai8)  // 128

}
