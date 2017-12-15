package main

import "fmt"

func main() {
	inline()
	returningFunc()
}
func inline() {
	i := 0
	increment := func() {
		i++
	}
	fmt.Println(i)
	increment()
	fmt.Println(i)
}

func returningFunc() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i +=2
		return
	}
}