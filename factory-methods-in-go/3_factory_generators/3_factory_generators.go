package main

import "fmt"

// factory generators are incredibly useful, when you want to construct instances of different structs or
// interfaces that are not mutually exclusive, or if you want multiple factories with different defaults.

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

// NewAnimal is an `Animal` factory that fixes the species as the value of its `AnimalFactory` instance
func (af AnimalFactory) NewAnimal(age int) Animal {
	return Animal{
		species: af.species,
		age:     age,
	}
}

// NewHouse is an `AnimalHouse` factory that fixes the name as the value of its `AnimalFactory` instance
func (af AnimalFactory) NewHouse(sizeInMeters int) AnimalHouse {
	return AnimalHouse{
		name:         af.houseName,
		sizeInMeters: sizeInMeters,
	}
}

func main() {
	dogFactory := AnimalFactory{
		species:   "dog",
		houseName: "kernel",
	}
	dog := dogFactory.NewAnimal(2)
	kernel := dogFactory.NewHouse(3)
	fmt.Print(dog)
	fmt.Print("\n", kernel)

	horseFactory := AnimalFactory{
		species:   "horse",
		houseName: "stable",
	}
	horse := horseFactory.NewAnimal(12)
	stable := horseFactory.NewHouse(30)
	fmt.Print("\n", horse)
	fmt.Print("\n", stable)
}
