package main

import "fmt"

func main() {
	print("day la ngon: ")
	switch x := 2; x {
	case 1:
		fmt.Println("ngon cai")
	case 2:
		fmt.Println("ngon 2")
	case 3:
		fmt.Println("ngon 3")
	case 4:
		fmt.Println("ngon 4")
	case 5:
		fmt.Println("ngon 5")
	default:
		fmt.Println("so ngon khong dung")
	}
}
