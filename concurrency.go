package main

import (
	"fmt"
	"time"
)

func main() {
	go blue3()
	go red5()

	time.Sleep(5000 * time.Millisecond)
	// time.Sleep(100 * time.Millisecond)
}

func red5() {
	for i := 0; i < 6; i++ {
		fmt.Println("red")
		time.Sleep(800 * time.Millisecond)
	}
}

func blue3() {
	for i := 0; i < 4; i++ {
		fmt.Println("blue")
		time.Sleep(1000 * time.Millisecond)
	}
}
