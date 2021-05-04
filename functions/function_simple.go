package main

import (
	"errors"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func walk(person Person) (string, error) {
	if person.Name == "Gabril" {
		return "", errors.New("Gabriel cannot to walk")
	}

	return "Hey Gabriel lets go to walk :)", nil
}

func main() {
	gabriel := Person{
		Name: "Gabriel",
		Age:  26,
	}
	walk, err := walk(gabriel)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(walk)
}
