package main

import (
	"fmt"

	"example.com/fibonacci"
)

func main() {
	// TODO: Scan Keyboard input
	n, err := fmt.Scanln()

	if err != nil {
		panic("An unexpected error occured")
	}

	fmt.Println(fibonacci.Fibonacci(n))
}
