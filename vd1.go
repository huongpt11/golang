package main

import (
	"fmt"
	"math"
)

type verb struct {
	x, y float64
}

func (v *verb) them(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}
func (v *verb) tt() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func main() {
	v := verb{3, 4}
	fmt.Println("v:, tt:", v, v.tt())
	v.them(10)
	fmt.Println("v:,tt:", v, v.tt())
}
