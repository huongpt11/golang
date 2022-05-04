package main

import "fmt"

func main() {
	a := [6]int{2, 3, 4, 5, 6, 7}
	var s []int = a[2:4]
	fmt.Println(s)
}
