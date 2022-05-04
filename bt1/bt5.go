package main

import "fmt"

type sinhvien struct {
	msv  string
	ten  string
	tuoi int
}

func main() {
	v := sinhvien{"b18dcvt221", "phung thi huong", 22}
	v.msv = "b18dcvt133"
	fmt.Println(v.msv)
}
