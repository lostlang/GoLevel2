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
		fmt.Println(err)
		return
	}
	println("ok") // а не это
}
