package main

import "fmt"

// For example, the struct for a “Person”, along with a “Greet” method would look like this:

type Person interface {
	Greet()
}

type person struct {
	Name string
	Age  int
}

func (p person) Greet() {
	fmt.Printf("Hi! My name is %s", p.Name)
}

// We can now make use of factory functions to create new instances of Person
// The simplest, and most commonly used factory is a function that takes in some arguments, and returns an
// instance of Person:

// Here, NewPerson returns an interface, and not the person struct itself
func NewPerson(name string, age int) Person {
	return &person{
		Name: name,
		Age:  age,
	}
}

// Factory functions are a better alternative to initializing instances using something like p := &Person{}
// because, the function signature ensures that everyone will supply the required attributes.

// For example, one can easily forget to initialize the Age attribute when using struct initialization.
// The function signature of NewPerson(name string, age int) ensures that both the name and age are supplied
// in order to construct a new instance of Person

// Interface factories: Factory functions can return interfaces instead of structs. Interfaces allow you to
// define behavior without exposing internal implementation.

func main() {
	MyInfo := NewPerson("ibrahim", 21)
	MyInfo.Greet()
}
