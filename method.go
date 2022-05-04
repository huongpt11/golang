package main

import (
	"fmt"
	"math"
)

type add struct { // khai bao mot struct co ten la add gom cac truong la x,y co kieu du lieu la float64
	x, y float64
}

func (a add) Abs() float64 { //khai bao mot method co ten la Abs  treen mot struct la add(vat nhan- reciever)
	return math.Sqrt(a.x*a.x + a.y*a.y) // tra ve gia tri tinh toan
}
func main() {
	a := add{3, 4}       // khai bao gia tri tinh
	fmt.Println(a.Abs()) // in ket qua ra man hinh
}
