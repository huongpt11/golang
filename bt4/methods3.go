package main

import "fmt"

type verb struct {
	x, y float64
}

func (v *verb) scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}
func scalex(v *verb, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}
func main() {
	v := verb{3, 4}
	v.scale(2)
	scalex(&v, 10)

	p := &verb{4, 3}
	p.scale(3)
	scalex(p, 8)
	fmt.Println(v, p)
}
