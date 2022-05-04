package main

import (
	"fmt"
	"math"
)

type myf float64

func (f myf) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func main() {
	f := myf(-math.Sqrt(64))
	fmt.Println(f.abs())
}
