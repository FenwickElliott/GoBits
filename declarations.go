package main

import "fmt"

// decaring a boolean
var a bool

// declaring an integer and setting it
var b = 9

// declaring strings
var c,d = "tom", "dick"

func main(){
	// 'short' assignments can only be made in functions
	e := true
	fmt.Println(a, b, c, d, e)
}