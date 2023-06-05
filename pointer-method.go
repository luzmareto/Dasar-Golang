package main

import "fmt"

type Man struct {
	Name string
}

// method
func (man *Man) Married() { //man diisi pada func main
	man.Name = "MR. " + man.Name //MR. adalah kata pertama pada output, man.name adalah kata ke 2
}

func main() {
	luz := Man{"luz"} //memanggil struct man + memberi value pada field name
	luz.Married()     //memanggil method

	fmt.Println(luz.Name) //MR. luz
	fmt.Println(luz)      //{MR. luz}
}
