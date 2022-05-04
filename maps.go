package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["x1"] = 6
	m["x2"] = 11
	m["x3"] = 2000
	fmt.Println("map:", m)
	m1 := m["x2"]
	fmt.Println(m1)
	fmt.Println("len(m):", len(m))
	delete(m, "x3")
	fmt.Println(m)
	_, prs := m["x3"]
	fmt.Println(prs)
	n := map[string]int{"k1": 11, "k2": 20}
	fmt.Println(n)
}
