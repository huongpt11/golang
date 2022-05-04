package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return fact(n-1) + n
}
func main() {
	fmt.Println(fact(7))
	var five func(n int) int
	five = func(n int) int {
		if n < 5 {
			return n
		}
		return five(n-1) * five(n-2)
	}
	fmt.Println(five(10))
}
