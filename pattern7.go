package main

import (
	"fmt"
)

type Animal interface {
	callMe() string
}

type Cat struct {
	name string
}

func (cat Cat) callMe() string {
	return "I'm furry fag " + cat.name
}

type Dog struct {
	name string
	woof string
}

func (dog Dog) callMe() string {
	return "I'm good boy " + dog.name + ", my favourite woof - " + dog.woof
}

func main() {
	fmt.Println()
	var Boris Cat = Cat{"Boris"}
	var Genri Dog = Dog{"Genri", "woooof"}

	var MyPets = []Animal{Boris, Genri}
	for _, pet := range MyPets {
		fmt.Println(pet.callMe())
	}
}
