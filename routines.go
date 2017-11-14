package main

import (
	"fmt"
	"strings"
)

func main() {

	first := make(chan string)
	second := make(chan string)
	third := make(chan string)

	go getone(first)
	go gettwo(second)
	go getthree(third)

	one := <-first
	two := <-second
	three := <-third

	res := strings.Join([]string{one, two, three}, " ")
	fmt.Println(res)
}

func getone(first chan string) {
	first <- "I"
}

func gettwo(second chan string) {
	second <- "Love"
}

func getthree(third chan string) {
	third <- "Goats"
}
