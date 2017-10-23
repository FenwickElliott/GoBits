package main

import "fmt"

func main(){
	// 3 clauses
	for i := 0; i < 10; i++{
		fmt.Println(i % 3)
	}

	fmt.Println("-------------")

	// 1 clause
	j:= 10
	for j > 0 {
		fmt.Println(j)
		j--
	}

	fmt.Println("-------------")

	// 0 clauses
	k := 10
	for {
		fmt.Println(k)
		if ( k == 0 ) {
			return
		} else {
			k--
		}
	}
}