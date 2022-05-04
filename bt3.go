package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func main() {
	var s []int
	printSlice(s)
	s = append(s, 0, 1, 2, 3)
	printSlice(s)
	s = append(s, 6, 7, 8, 9)
	printSlice(s)

}
