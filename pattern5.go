package main

import "fmt"

type Cat struct {
	age int
}

func (cat *Cat) Grow() *Cat {
	cat.age++
	return cat
}

func main() {
	var Barsik = Cat{}
	fmt.Println(Barsik)
	Barsik.Grow().Grow().Grow().Grow().Grow().Grow()
	fmt.Println(Barsik)
}
