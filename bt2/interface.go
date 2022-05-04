package main

import (
	"fmt"
	"math"
)

type hinh interface {
	chuvi() float64
	dientich() float64
}
type hinhchunhat struct {
	chieudai, chieurong float64
}
type hinhtron struct {
	bankinh float64
}

func (hcn hinhchunhat) chuvi() float64 {
	return 2*hcn.chieudai + 2*hcn.chieurong
}
func (hcn hinhchunhat) dientich() float64 {
	return hcn.chieurong * hcn.chieudai
}
func (ht hinhtron) chuvi() float64 {
	return 2 * math.Pi * ht.bankinh
}
func (ht hinhtron) dientich() float64 {
	return math.Pi * ht.bankinh * ht.bankinh
}
func xx(h hinh) {
	fmt.Println(h)
	fmt.Println(h.chuvi())
	fmt.Println(h.dientich())
}
func main() {
	hcn := hinhchunhat{3, 4}
	ht := hinhtron{5}
	xx(hcn)
	xx(ht)
}
