package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func main(){
	fmt.Println("7+8 =", add(7,8))
}