package main

import (
	"fmt"
	"math"
)

type cal struct {
	a, b float64
}

func abs(c cal) float64 {
	return math.Sqrt(c.a*c.a + c.b*c.b)
}
func main() {
	c := cal{3, 4}
	fmt.Println(abs(c))
}
