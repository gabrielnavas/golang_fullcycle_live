package main

import "fmt"

type Name string
type Names []Name

func main() {
	var names Names
	var name1 Name = "Maria"
	var name2 Name = "Gabriel"

	names = append(names, name1)
	names = append(names, name2)

	for _, n := range names {
		fmt.Println(n)
	}

}
