package main

import "fmt"

func main() {
	scores := map[string]int{
		"tom":   12,
		"dick":  8,
		"harry": 100,
	}
	fmt.Println(scores)
	scores["charles"] = 1000
	fmt.Println(scores)
	fmt.Println(scores["charles"])

	v, ok := scores["charles"]
	if ok {
		fmt.Printf("charles was present. His score was %d\n", v)
	} else {
		fmt.Println("charles was not pressent")
	}

	delete(scores, "charles")

	v, ok = scores["charles"]
	if ok {
		fmt.Printf("charles was present. His score was %d\n", v)
	} else {
		fmt.Println("charles was not pressent")
	}
}
