package main

import "fmt"

func main() {
	fmt.Println(factorial(8))
	fmt.Println(fibonacci(8))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial( n - 1 )
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci( n - 1 ) + fibonacci( n - 2 )
	}
}