package main

import "fmt"

func main() {
	var a [5]int // khai bao mang a co 5 phan tu kieu int, a rong
	fmt.Println(a)
	a[3] = 100 // set a taji vi tri 3 co gia tri la 100
	fmt.Println(a)
	fmt.Println("a3 co gia tri la:", a[3])
	fmt.Println("len a:", len(a)) // in ra do dai cua mang a
	b := [5]int{1, 2, 3, 4, 5}    // khai bao mang b co 5 phan tu
	fmt.Println(b)
	c := b[2:4] // c gom cac phan tu nam o vi tri thu 2 den thu 3 trong mang b
	fmt.Println(c)
	var d [2][4]int          // khai bao mang d hai chieu , moi mang gom co 4 phan tu
	for i := 0; i < 2; i++ { //vong lap for xac dinh cac phan tu cua mang hai chieu
		for j := 0; j < 4; j++ {
			d[i][j] = j - i
		}
	}
	fmt.Println(d)
}
