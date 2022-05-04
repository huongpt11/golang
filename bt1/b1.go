package main

import "fmt"

func main() {
	i, j := 10, 40  // khai bao gia tri cua i j
	p := &i         //tao con tro p tro toi i
	fmt.Println(*p) // inra gia tri cua con tro p tro toi i la 10
	*p = 15         // cai gia tri moi cho i qua con tro p
	fmt.Println(i)  // in ra gia tri mooi cho i
	p = &j          /// tro toi con tro j
	*p = *p / 2     // tinh toan qua con tro
	fmt.Println(j)  // in ra gia tri moi cua con tro j
}
