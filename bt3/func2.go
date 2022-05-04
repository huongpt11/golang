package main

import (
	"fmt"
	"math"
)

func cal(a, b, c float64) float64 {
	return math.Sqrt(a) + b*c
}
func cals(a, b float64) float64 {
	return math.Dim(a, b)
}
func main() {
	cal := cal(4, 5, 6)
	fmt.Println(cal)
	cals := cals(8, 7)
	fmt.Println(cals)
}
