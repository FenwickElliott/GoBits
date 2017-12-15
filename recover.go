package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("I forgot to get milk")
	//unreachable code...
}