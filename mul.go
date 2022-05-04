package main

import "fmt"

func value() (int, int) {
	return 7, 3
}
func main() {
	a, b := value()
	fmt.Println(a)
	fmt.Println(b)
}
