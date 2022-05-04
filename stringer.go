package main

import "fmt"

type list struct {
	name string
	age  int
}

func (l list) String() string {
	return fmt.Sprintf("%v (%v year)", l.name, l.age)
}
func main() {
	a := list{"huong", 15}
	b := list{"hoa", 16}
	fmt.Println(a, b)
}
