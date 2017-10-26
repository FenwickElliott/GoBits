package main

import "fmt"

func main() {
	// build array
	b := [5]int{1, 1, 0, 0, 1}
	fmt.Println(b)

	// slice it
	// nb s:= b[:] would slice whole array
	s := b[3:5]
	fmt.Println(s)

	// changes to slice are reflected in aray
	s[0] = 1
	fmt.Println(b)

	// slices and can be declared independently, go will build an appropriate array in the background and slice it
	strings := []string{"tom", "dick", "harry"}
	fmt.Println(strings)

	x := make([]bool, 1)
	fmt.Println(x)
}
