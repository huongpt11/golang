package main

import (
	"fmt"
	"math"
)

func add(x, lim float64) float64 {
	if v := math.Abs(2 * x); v < lim {
		return v
	}
	return lim

}
func main() {
	fmt.Println(
		add(6, 20),
		add(9, 10),
	)
}
