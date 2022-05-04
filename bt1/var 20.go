package main

import "fmt"

func main() {
	fmt.Println("day so la:")
	defer fmt.Println("done")
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}
