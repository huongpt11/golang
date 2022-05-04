package main

import "fmt"

func add(x, y int) int {
	return x + y
}
func main() {
	defer fmt.Println(add(2, 3))
	fmt.Println("done")
}
