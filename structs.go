package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	x, y float64
}

func (r *Rectangle) area() float64 {
	return r.x * r.y
}

func (r *Rectangle) diagonal() float64 {

	return math.Sqrt(math.Pow(r.x,2) + math.Pow(r.y,2))
}

func main(){
	r := Rectangle{3,4}
	fmt.Println(r.area())
	fmt.Println(r.diagonal())
}