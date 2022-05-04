package main

import (
	"fmt"
)

func main() {
	fmt.Print("hoc luc ki nay cua ban la:\n")
	switch x := 10; x {
	case 3:
		fmt.Println("yeu")
	case 6:
		fmt.Println("trung binh")
	case 8:
		fmt.Println("gioi")
	case 10:
		fmt.Println("xuat sac")
	default:
		fmt.Println("dang cap nhat du lieu")
	}

}
