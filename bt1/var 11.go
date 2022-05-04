package main

import (
	"fmt"
	"math"
)

func tinh(x, y, max float64) float64 {
	if v := math.Dim(x, y); v < max {
		return v
	}
	return max
}
func main() {
	fmt.Println(
		tinh(6, 7, 10),
		tinh(1, 2, 1))
}
