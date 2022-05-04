package main

import (
	"fmt"
	"math"
)

func pow(x, y, max float64) float64 {

	if v := math.Pow(x, y); v < max {
		return v
	} else {
		fmt.Printf("%g>=%g\n", v, max)
	}
	return max

}
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
