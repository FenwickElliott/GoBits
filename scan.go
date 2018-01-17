package main

import "fmt"

var response string

func main() {
	fmt.Print("User input: ")
	// Scan must know how many whitespace seporated peices of input were given
	fmt.Scan(&response)
	fmt.Println("Response was: ", response)
}
