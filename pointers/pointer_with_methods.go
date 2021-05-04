package main

import "fmt"

func main() {
	name := "Gabriel"
	// value of the variable: Gabriel
	fmt.Println(name)
	// address of the variable in hex example: 0xc000010240
	fmt.Println(&name)

	var namePointer *string
	namePointer = &name

	// address of the variable in hex example: 0xc000010240
	fmt.Println(namePointer)

	// value of the variable: Gabriel
	fmt.Println(*namePointer)

	if namePointer == &name {
		fmt.Println("namePointer is equals name")
	}

	if *namePointer == name {
		fmt.Println("value of the namePointer is equals name")
	}
}
