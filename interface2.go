package main

import "fmt"

type M interface{}

func main() {
	var i M

	i = 42
	fmt.Printf("%T, %v", i, i)
}
