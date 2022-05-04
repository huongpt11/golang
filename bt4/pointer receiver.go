package main

import (
	"fmt"
	"math"
)

type low struct {
	x, y float64
}

func (l low) abs() float64 {
	return math.Sqrt(l.x * l.y)
}
func (l *low) them(f float64) {
	l.x = l.x * f
	l.y = l.y * f
}
func main() {
	low := low{9, 9}
	low.them(10)
	fmt.Println(low.abs())
}
