package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Dog struct {
	name string
}

type Cat struct {
	name string
}

type Batler struct {
	Dog
	Cat
}

func (f Dog) woof() string {
	return f.name + " woof"
}

func (f Cat) meow() string {
	return f.name + " meow"
}

func (f Batler) fight() string {
	rand.Seed(time.Now().Unix())
	switch rand.Intn(3) {
	case 0:
		return "Win " + f.Dog.name
	case 1:
		return "Win " + f.Cat.name
	default:
		return "No winer"
	}
}

func main() {
	var dog = Dog{"Korgi"}
	var cat = Cat{"Nina"}
	var f = Batler{dog, cat}
	fmt.Println(f.woof())
	fmt.Println(f.meow())
	fmt.Println(f.fight())
}
