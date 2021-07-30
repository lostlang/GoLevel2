package main

import "fmt"

type BaseClass interface {
	getName() string
}

type CatHero struct {
	name   string
	damage int
	BaseClass
}

func (cat *CatHero) swapToWarior() *CatHero {
	cat.BaseClass = Warior{}
	return cat
}

func (cat *CatHero) swapToMage() *CatHero {
	cat.BaseClass = Mage{}
	return cat
}

type Warior struct{}

func (Warior) getName() string {
	return "I'am warior"
}

type Mage struct{}

func (Mage) getName() string {
	return "I'am mage"
}

func main() {
	var Boris = CatHero{}
	Boris.swapToWarior()
	fmt.Println(Boris.getName())
	Boris.swapToMage()
	fmt.Println(Boris.getName())
}
