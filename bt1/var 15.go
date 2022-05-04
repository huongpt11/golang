package main

import (
	"fmt"
	"math/rand"
)

func main() {
	print("bao gio di chi the?\n ")
	switch x := rand.Intn(8); x {
	case 1:
		fmt.Println("6h nha")
	case 2:
		fmt.Println("7h nha")
	case 3:
		fmt.Println("8h nha")
	case 4:
		fmt.Println("9h nha")
	default:
		fmt.Println("chiu nha")

	}
}
