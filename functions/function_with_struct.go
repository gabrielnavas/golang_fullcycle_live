package main

import (
	"errors"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (person Person) sayTo(other Person) (string, error) {
	if other.Name == "Maria" {
		return "", errors.New(person.Name + " cannot to say with " + other.Name)
	}

	return "Hey " + other.Name + " :)", nil
}

func (person Person) walk() (string, error) {
	if person.Name == "Ronaldo" {
		return "", errors.New("Gabriel cannot to walk")
	}

	return "Hey Gabriel lets go to walk :)", nil
}

func main() {
	gabriel := Person{
		Name: "Gabriel",
		Age:  26,
	}
	otherPerson := Person{"Júlia", 22}
	walk, errWalk := gabriel.walk()
	if errWalk != nil {
		fmt.Println(errWalk)
	}
	fmt.Println(walk)
	sayTo, errSay := gabriel.sayTo(otherPerson)
	if errSay != nil {
		fmt.Println(errWalk)
	}
	fmt.Println(sayTo)
}
