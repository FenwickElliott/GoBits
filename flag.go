package main

import (
	"flag"
	"fmt"
)

func main() {
	flagName := flag.String("name", "default", "description")
	flagBool := flag.Bool("bool", false, "boolean")
	flagInt := flag.Int("int", 0, "integer")
	flag.Parse()
	fmt.Println(*flagName)
	fmt.Println(*flagBool)
	fmt.Println(*flagInt)
	// Usage: go run main.go -name=whatever -bool=true -int=81
}
