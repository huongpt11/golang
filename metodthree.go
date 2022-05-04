package main

import (
	"fmt"
	"math"
)

type para struct { // khia bao struct co ten para gom cac truong x,y co kieu float64
	x, y float64
}

func (p para) abs() float64 { //khai bao method co ten abs co reciever la para
	return math.Sqrt(p.x*p.x + p.y*p.y) // tra ve phep tinh
}
func (p *para) add(f float64) { // khai bao method co ten add co reciever la con tro tro toi para
	p.x = p.x * f
	p.y = p.y * f
}

func main() {
	p := para{3, 4}
	fmt.Println(p.abs()) // inra gia tri tai p , cau lenh danh cho method
	p.add(10)
	fmt.Println(p.abs()) // in ra gia tri sau khi qua method add
}
