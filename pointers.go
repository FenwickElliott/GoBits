package main

import "fmt"

func main(){
	i := 81
	fmt.Println(i)
	// & access memory address of variable
	j := &i
	fmt.Println(j)
	// * access the variable at that address
	*j++
	fmt.Println(*j)
	fmt.Println(i)
}