package main

import "fmt"

func main() {
	var x, y int = 3, 4
	var z float64 = float64(x*4 - y)
	var t uint = uint(z)
	fmt.Println(x, y, t)
}
