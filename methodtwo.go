package main

import (
	"math"
)

type myfloat float64

func (f myfloat) abs() float64 { // khai bao method co ten abs() voi vat nhan-reciever la myfloat
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func main() {
	f := myfloat(math.Sqrt(2))
	println(f.abs())
}
