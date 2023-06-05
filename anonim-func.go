package main

import "fmt"

// BlackList adalah alias func dengan parameter string dan return bool
type BlackList func(string) bool

func registerUser(name string, blackList BlackList) { //blacklist valuenya adalah alias Blacklist
	//melakukan return boolean

	//jika value pada parameter name
	if blackList(name) {
		fmt.Println("you are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func main() {
	//blackList adalah anonim func yg isinya seperti alias blacklis
	blackList := func(name string) bool {
		return name == "admin" //jika var ini dipanggil, maka setiap var name admin akan dinyatakan you are blocked
	}

	//memanggil func registerUser, memberi value pada parameter name lalu disaring ke var blacklist lalu melakukan pengecekan di bari 12
	registerUser("admin", blackList) //you are blocked admin
	registerUser("luz", blackList)   //welcome luz

	//setiap value root akan mengeluarkan output you are blocked
	registerUser("root", func(name string) bool {
		return name == "root"
	})

	registerUser("luz", func(name string) bool {
		return name == "root"
	})
}
