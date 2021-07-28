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

type Visitor struct{}

func (v Visitor) CatAge(target Cat) string {
	if target.age > 34 {
		return "Call to Guinness book of records"
	} else {
		return fmt.Sprint(target.age * 8)
	}
}

func main() {
	var MyCats = make([]Cat, 10)
	rand.Seed(time.Hour.Microseconds())
	for i := range MyCats {
		MyCats[i].age = rand.Intn(30) + 10
	}
	var v Visitor
	for _, i := range MyCats {
		fmt.Println(v.CatAge(i))
	}
}
