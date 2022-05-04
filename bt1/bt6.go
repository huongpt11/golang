package main

import "fmt"

type dtb struct {
	msv string
	ten string
	dtb float64
}

func main() {
	var p dtb
	p.msv = "b18dcv221"
	p.ten = "phung thi huong"
	p.dtb = 8.8
	fmt.Println("ma sinh vien la:", p.msv)
	fmt.Println("ten:", p.ten)
	fmt.Println("dtb la:", p.dtb)

}
