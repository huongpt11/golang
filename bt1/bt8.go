package main

import "fmt"

type ds struct {
	stt  int
	diem int
}

func main() {
	v1 := ds{1, 10}
	v2 := ds{stt: 2}
	v3 := ds{}
	p := &ds{1, 10}
	fmt.Println(v1, v2, v3, p)
}
