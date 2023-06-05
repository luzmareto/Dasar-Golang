package main

import (
	"fmt"
	"strings"
)

// value akan diisi padafunc main
type Address struct {
	City, Province, Country string
}

func main() {

	var address1 Address = Address{"subang", "jawa barat", "indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	*address2 = Address{"malang", "jawa timur", "indonesia"}

	//karena semua var pointer mengacu ke address 1. maka outputnya dari semua variabel adalah sama
	fmt.Println(address1) //{malang jawa timur indonesia}
	fmt.Println(address2) //&{malang jawa timur indonesia}
	fmt.Println(address3) //&{malang jawa timur indonesia}

	fmt.Println(strings.Repeat("=", 80))

	//mengambil format struct address
	var address4 Address = Address{} //outputnya kosong, karena fieldnya belum diisi

	// memberi value struct address 1 per 1
	address4.City = "Depok"
	fmt.Println(address4)
}
