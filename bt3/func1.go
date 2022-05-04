package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func adds(a, b, c int) int {
	return a - b + c
}
func main() {
	add := add(10, 5)
	fmt.Println(add)
	adds := adds(10, 4, 7)
	fmt.Println(adds)
}
