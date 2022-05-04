package main

import "fmt"

type list struct {
	name string
	age  int
}

func addlist(name string) *list {
	p := list{name: name}
	p.age = 10
	return &p
}
func main() {
	fmt.Println(list{"huong", 20})
	fmt.Println(list{"hoa", 22})
	fmt.Println(&list{"hung", 9})
	fmt.Println(addlist("hai"))
	s := list{"cao", 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
}
