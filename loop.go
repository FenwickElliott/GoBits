package main

import "fmt"

func main() {
	// 3 clauses
	for i := 0; i < 10; i++ {
		fmt.Println(i % 3)
	}

	fmt.Println("-------------")

	// 1 clause
	j := 10
	for j > 0 {
		fmt.Println(j)
		j--
	}

	fmt.Println("-------------")

	// 0 clauses
	k := 10
	for {
		fmt.Println(k)
		if k == 0 {
			break
		} else {
			k--
		}
	}

	// ranging
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for l, v := range pow {
		fmt.Printf("2**%d = %d\n", l, v)
	}
}
