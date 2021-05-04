package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	gabriel := Person{
		Name: "Gabriel",
		Age:  26,
	}
	maria := Person{"Maria", 19}

	fmt.Println(gabriel)
	fmt.Println(maria)
}
