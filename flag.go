package main

import (
	"flag"
	"fmt"
)

func main() {
	flagName := flag.String("name", "default", "description")
	flag.Parse()
	fmt.Println(*flagName)
	// Usage: go run main.go -name=whatever
}
