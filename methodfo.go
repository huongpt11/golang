package main

// method ket hop pointer and func
import (
	"fmt"
	"math"
)

type par struct { // khia bao struct par co cac truong x ,y
	x, y float64
}

func abs(p par) float64 { // khia bao function
	return math.Sqrt(p.x*p.x + p.y*p.y) // ket qua tra ve
}
func adds(p *par, f float64) { // khai bao function co ket hop con tro
	p.x = p.x * f // thay doi gia tri x,y
	p.y = p.y * f
}
func main() {
	p := par{3, 4}
	fmt.Println(abs(p))
	adds(&p, 10) // doi voi function can phai tro toi dia chi cua p de thay doi gia tri
	fmt.Println(abs(p))
}
