package main

import "fmt"

func compare[T comparable](s []T, x T) int { // khai bao ham compare
	for i, v := range s { // vofng for chay het mang s
		if v == x { //neu v==x
			return i // thi tra ve gia tri i
		}
	}
	return -1 // neu khong tim thay tra ve gia tri -1

}
func main() {
	si := []int{2, 4, 6, 8, 10, 12, 14}
	fmt.Println(compare(si, 12))
	sm := []string{"i", "love", "riu", "rui"}
	fmt.Println(compare(sm, "ri"))
}
