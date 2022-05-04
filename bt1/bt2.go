package main

import (
	"fmt"
)

func main() {
	x, y := 3, 4
	p := &x
	fmt.Println(*p)
	*p = 10
	fmt.Println(x)
	p = &y
	*p = *p * 2
	fmt.Println(y)

}
