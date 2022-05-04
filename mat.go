package main

import "fmt"

func main() {
	var a [3]string
	a[0] = "phung"
	a[1] = "thi"
	a[2] = "huong"
	fmt.Println(a[0], a[1], a[2])
	fmt.Println(a)
	//chuoiso:=[6]{}
	//s := "aa"
	//var s2 string
	i := 6
	with_pointer(&i)
	fmt.Printf("outer i=%d\n", i)

	fmt.Println("-----------------")
	with_var(i)
	fmt.Printf("outer i=%d\n", i)
}

func with_pointer(i *int) { //pass by reference
	*i = 10
	fmt.Printf("in func i=%d\n", *i)
}
func with_var(i int) { // pass by value
	i = i + 1
	fmt.Printf("in func i=%d\n", i)
}
