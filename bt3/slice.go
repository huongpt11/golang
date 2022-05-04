package main

import "fmt"

func main() {
	a := make([]string, 3) // tao mang string co cap la 3,a rong
	fmt.Println(a)
	a[0] = "phung"
	a[1] = "thi"
	a[2] = "huong"
	fmt.Println(a)
	fmt.Println(a[2])
	fmt.Println("len a la:", len(a))
	a = append(a, "_ma sinh vien:") // them ptu vao mang a
	a = append(a, "B18DCVT221")
	fmt.Println(a)
	c := make([]string, len(a)) // tao mang c co do dai bang mang a
	copy(c, a)                  // sao chep mang a sang mang c
	fmt.Println(c)
	m := a[2:4] // mang m co gia tri tu v tri thu 2 den vi tri thu 3trong mang a
	fmt.Println(m)
	n := a[2:] // mang n co gia tri tu vi tri thu 2 tro di trong mang a
	fmt.Println(n)
	add := make([][]int, 4)  // tao mang da chieu add co cap toi da la 4<co the la 1 2 3 4>
	for i := 0; i < 4; i++ { //xet vong lap i chaay tu 1 den 3
		lenadd := i + 1
		add[i] = make([]int, lenadd) // xet cap cua tung mang trong add <vi du tai i=0 thi len=1, i=2 thi len =3>
		for j := 0; j < lenadd; j++ {
			add[i][j] = i + j
		}
	}
	fmt.Println(add)
}
