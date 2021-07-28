package main

import (
	"fmt"
	"math/rand"
	"time"
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

func (mother CatF) giveBirth() interface{} {
	if mother.pregnate {
		defer func() {
			mother.pregnate = false
			mother.partner = nil
		}()
		rand.Seed(time.Now().Unix())
		if rand.Intn(2) == 0 {
			var child = CatF{}
			child.name = mother.name + " II"
			return child
		} else {
			var child = CatM{}
			child.name = mother.partner.name + " II"
			return child
		}
	} else {
		return nil
	}
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

	var newCat = Olga.giveBirth()

	fmt.Println(newCat)
}
