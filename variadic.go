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
	res := add(1,2,3,4,5)
	fmt.Println(res)
}