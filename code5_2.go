package main

import "fmt"

type customError struct {
	msg error
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	err := test()
	if err != nil {
		fmt.Printf("%T", err)
		return
	}
	println("ok") // а не это
}
