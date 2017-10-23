package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	if true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	fmt.Print("Go runs on ")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		//simply print os
		fmt.Printf("%s", os)
	}

	// blank switch statments serve well for multiple if else evaluations
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}