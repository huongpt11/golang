package main

import "fmt"

func main() {
	var a [4]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	fmt.Println(a[0], a[1], a[2], a[3])
	fmt.Println(a)
	m := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(m)
}
