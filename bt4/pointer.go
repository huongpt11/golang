package main

import "fmt"

func value(i int) {
	i = 5
	//fmt.Println("value before:", ival)
}
func valuep(ivalp *int) {
	*ivalp = 6
}
func main() {
	i := 7
	fmt.Println(i)
	value(i)
	fmt.Println(i)
	valuep(&i)
	fmt.Println(i)
	fmt.Println(&i)
}
