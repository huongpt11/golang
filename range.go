package main

import "fmt"

func main() {
	var p = []int{1, 2, 4, 8, 16, 32, 64, 128}
	sum := int(0)
	for _, v := range p {
		fmt.Printf("%d\n", v)
		sum += v
	}
	fmt.Println(sum)
}
