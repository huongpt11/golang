package main

import (
	"fmt"
	"math"
)

type tinhtoan interface {
	dientich() float64
	chuvi() float64
}
type hinhvuong struct {
	canh float64
}
type hinhtron struct {
	bankinh float64
}

func (hv hinhvuong) dientich() float64 {
	return float64(hv.canh) * float64(hv.canh)
}
func (hv hinhvuong) chuvi() float64 {
	return 4 * hv.canh
}
func (ht hinhtron) dientich() float64 {
	return math.Pi * ht.bankinh * ht.bankinh
}
func (ht hinhtron) chuvi() float64 {
	return 2 * math.Pi * ht.bankinh
}
func x(tt tinhtoan) {
	fmt.Println(tt)
	fmt.Println(tt.dientich())
	fmt.Println(tt.chuvi())
}
func main() {
	canh := hinhvuong{4}
	bk := hinhtron{5}
	x(canh)
	x(bk)
}
