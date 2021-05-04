package main

import (
	"fmt"
)

type Car struct {
	Name string
	Year int
}

func (c Car) changeName(name string) {
	c.Name = name
}

func (c *Car) changeNamePointer(name string) {
	c.Name = name
}

func main() {
	car := Car{
		Name: "Fusca",
		Year: 1969,
	}

	car.changeName("Ferrari")
	// car name is Fusca, not change to Ferrari, what??
	fmt.Println(car.Name)

	// change value name reference
	car.changeNamePointer("Volks")
	// car name is Volks
	fmt.Println(car.Name)
}
