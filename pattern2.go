package main

import "fmt"

type Cat struct {
	name string
}

type CatProducer struct {
	count int
	name  string
}

func (prod *CatProducer) getCat() *Cat {
	prod.count++
	return &Cat{prod.name + " " + fmt.Sprint(prod.count)}
}

func (prod *CatProducer) setCount(num int) {
	prod.count = num
}

func (prod *CatProducer) setName(s string) {
	prod.name = s
}

func main() {
	var breader1 = CatProducer{}
	breader1.setCount(0)
	breader1.setName("Murzik")

	var cat1 = breader1.getCat()
	var cat2 = breader1.getCat()

	fmt.Println(cat1, cat2)
}
