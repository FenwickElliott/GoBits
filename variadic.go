package main

import "fmt"

func add(n ... int) int {
	res := 0

	for _, v := range n {
		res += v
	}

	return res
}

func main() {
	num := add(1,2,3,4,5)
	fmt.Println(num)

	fib := []int{1,1,2,3,5,8,13,21}
	fibSum := add(fib...)
	fmt.Println(fibSum)
}