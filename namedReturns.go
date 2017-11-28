package main

import "fmt"

func name(a, b, c string) (x, y, z string) {
	x = a
	y = b
	z = c
	return
}

func main() {
	x, y, z := name("alpha", "bravo", "charlie")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
