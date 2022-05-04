package main

import "fmt"

func x(m ...int) {
	fmt.Print(m, "")
	sum := 0
	for _, m := range m {
		sum += m
	}
	fmt.Println(sum)
}
func main() {
	x(1, 2, 4)
	x(2, 3, 4)
	m := []int{1, 2, 3, 4}
	x(m...)
}
