package main

import "fmt"

func main() {
	ds := [4]string{
		"phung thi huong",
		"ngo thi xuan",
		"cao thi hoa",
		"tran van mai",
	}
	fmt.Println(ds)
	a := ds[0:2]
	b := ds[1:4]
	fmt.Println(a, b)
	b[0] = "chua cap nhap"
	fmt.Println(a, b)
	fmt.Println(ds)
}
