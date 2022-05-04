package main

import "fmt"

func adds() func() int {
	var x, y int
	x = 6
	y = 3
	return func() int {

		return x + y
	}
}
func main() {
	fun := adds()
	fmt.Println(fun())
}
