package main

import "fmt"

type Cat struct {
	name string
}

type CatProducer struct {
	count *int
	name  string
}

func (prod CatProducer) getCat() *Cat {
	*prod.count++
	return &Cat{prod.name + " " + fmt.Sprint(*prod.count)}
}

func main() {
	var breader1Cats = 0
	var breader2Cats = 100

	var breader1 = CatProducer{&breader1Cats, "Murzik"}
	var breader2 = CatProducer{&breader2Cats, "Peach"}

	var cat1 = breader1.getCat()
	var cat2 = breader2.getCat()
	var cat3 = breader1.getCat()
	var cat4 = breader1.getCat()

	fmt.Println(cat1, cat2, cat3, cat4)
}
