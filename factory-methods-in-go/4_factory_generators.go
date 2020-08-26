package main

import "fmt"

// We can use factory functions to make factories with some defaults built in. For example, if we want to make
// a “dog factory” and a “horse factory” instead of a generic “animal factory”, we use a generator function:

type Animal struct {
	species string
	age     int
}

type AnimalHouse struct {
	name         string
	sizeInMeters int
}

type AnimalFactory struct {
	species   string
	houseName string
}

func (af AnimalFactory) NewAnimal(age int) Animal {
	return Animal{
		species: af.species,
		age:     age,
	}
}

func (af AnimalFactory) NewHouse(sizeInMeters int) AnimalHouse {
	return AnimalHouse{
		name:         af.houseName,
		sizeInMeters: sizeInMeters,
	}
}

func NewAnimalFactory(houseName string) func(species string) AnimalFactory {
	return func(species string) AnimalFactory {
		return AnimalFactory{
			species:   species,
			houseName: houseName,
		}
	}
}

func main() {
	newOne1 := NewAnimalFactory("kernel")
	dogFactory := newOne1("dog") // This creates the factory
	dog := dogFactory.NewAnimal(2)
	kernel := dogFactory.NewHouse(3)
	fmt.Print(dog)
	fmt.Print("\n", kernel)

	newOne2 := NewAnimalFactory("stable")
	horseFactory := newOne2("horse")
	horse := horseFactory.NewAnimal(12)
	stable := horseFactory.NewHouse(30)
	fmt.Print("\n", horse)
	fmt.Print("\n", stable)
}
