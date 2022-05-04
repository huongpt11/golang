package main

import "fmt"

func add() func() float64 {
	var x, y, z float64
	x = 7
	y = 6
	z = 3.5
	return func() float64 {
		return x*y - z
	}
}
func main() {
	out := add()
	fmt.Println(out())
}
