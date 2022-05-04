package main

import (
	"fmt"
	"math"
)

type know struct {
	x, y, z float64
}

func tt(k know) float64 {
	return math.Sqrt(k.z * k.x * k.y)
}
func ttt(k know, f float64) {
	k.x = k.x * f
	k.y = k.y * f
	k.z = k.z * f
}
func main() {
	k := know{3, 4, 5}
	ttt(k, 15)
	fmt.Println(tt(k))
}
