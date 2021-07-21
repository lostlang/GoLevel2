package main

import (
	"fmt"
)

func Foo() (err error) {
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
