package main

import (
	"fmt"

	"example.com/fibonacci"
)

func main() {
	var n int
	fmt.Println("Fib Sequence calculator")
	fmt.Print("Please enter a number: ")
	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		panic(err)
	}

	fmt.Println(fibonacci.Fibonacci(n))
}
