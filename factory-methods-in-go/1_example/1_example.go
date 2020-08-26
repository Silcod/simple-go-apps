package main

import "fmt"

type Unicorn struct {
	Name   string
	Age    int
	weight int // Note this field is not exportable
}

func main() {
	// When creating an instance of a type, we need to decide if we want a `Pointer` or the actual `Value`
	// To create a `Value` we use the var keyword with a type:
	var larry Unicorn
	larry.Name = "Larry"
	larry.Age = 21
	fmt.Print(larry)

	// To create a `Pointer`
	nick := new(Unicorn)
	nick.Name = "Nick"
	nick.Age = 22
	nick.weight = 45
	fmt.Print("\n", nick)
}
