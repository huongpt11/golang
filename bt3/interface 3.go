package main

import (
	"fmt"
	"math"
)

type hinhvuong struct {
	canh float64
}
type hinhtron struct {
	bankinh float64
}
type tamgiac struct {
	a, b, c, h float64
}
type tinhtoan interface {
	chuvi() float64
	dientich() float64
}

func (hv hinhvuong) chuvi() float64 {
	return 4 * hv.canh
}
func (hv hinhvuong) dientich() float64 {
	return hv.canh * hv.canh
}
func (ht hinhtron) chuvi() float64 {
	return 2 * math.Pi * ht.bankinh
}
func (ht hinhtron) dientich() float64 {
	return math.Pi * ht.bankinh * ht.bankinh
}
func (tg tamgiac) chuvi() float64 {
	return tg.a + tg.b + tg.c
}
func (tg tamgiac) dientich() float64 {
	return (tg.a * tg.h) / 2
}
func yyy(tt tinhtoan) {
	fmt.Println(tt)
	fmt.Println(tt.chuvi())
	fmt.Println(tt.dientich())
}

func main() {
	hv := hinhvuong{4}
	ht := hinhtron{5}
	tg := tamgiac{5, 4, 3, 6}
	yyy(hv)
	yyy(ht)
	yyy(tg)

}
