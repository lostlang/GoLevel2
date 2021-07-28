package main

import (
	"fmt"
)

type Cat struct {
	name string
	age  int
}

type CatF struct {
	Cat
	pregnate bool
	partner  *CatM
}

type CatM struct {
	Cat
}

func (father CatM) breed(mother *CatF) {
	mother.pregnate = true
	mother.partner = &father
}

func main() {
	var Olga = CatF{}
	Olga.age = 15
	Olga.name = "Olga"

	var Boris = CatM{}
	Boris.age = 12
	Boris.name = "Boris"

	fmt.Println(Olga)

	Boris.breed(&Olga)

	fmt.Println(Olga)
}
