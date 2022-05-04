package main

import "fmt"

type list struct {
	ten  string
	tuoi int
}

func main() {
	v := list{"phung thi huong", 20}
	p := &v
	p.tuoi = 10
	//fmt.Println("ten:", p.ten)
	//fmt.Println("tuoi", p.tuoi)
	fmt.Println(v)
}
