package main

import (
	"fmt"
	"math"
)

type addm struct {
	x, y float64
}

func absd(v addm) float64 {
	return math.Sqrt(v.y*v.x - v.y)
}
func main() {
	v := addm{5, 6}
	fmt.Println(absd(v))
}
