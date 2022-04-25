package main

import "fmt"

type input struct { //  khai bao interface gom hai truong la month  va year//
	month int
	year  int
}

func (i input) inlich() {
	var day int
	switch i.month { // xet nam nhuan//
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
	case 2:
		if i.year%4 == 0 && i.year%100 != 0 || i.year%400 == 0 {
			day = 29
		} else {
			day = 28
		}
	default:
		day = 30

	}
	d := 1 // ngay dau tien cua thang//
	a := (14 - i.month) / 12
	y := i.year - a
	m := i.month + 12*a - 2
	dow := (d + y + y/4 - y/100 + y/400 + (31*m)/12) % 7 // cong thuc zeller,tinh thu cua ngay dau tien trong thang
	fmt.Printf("\n%3s%3s%3s%3s%3s%3s%3s\n", "S", "M", "T", "W", "T", "F", "S")
	for i := 0; i < 7; i++ {
		if i < dow {
			fmt.Printf("   ")
		} else {
			fmt.Printf("%3d", d)
			dow++
			d++
			break // thoat vong lap//
		}
	}

	for {

		fmt.Printf("%3d", d) // inn  ra nhung ngay tiep theo//
		dow++                // tang gia tri bien xet//
		if dow == 7 {
			fmt.Printf("\n")
			dow = 0 // tra ve gia  tri  0 roi xet tiep//
		}
		if d == day {
			fmt.Printf("\n")
			break
		} else {
			d++
		}
	}

}
func main() {
	i := input{7, 4} //  nhap  vao  thang , nam de xet//
	fmt.Printf("calender of %v %v is:", i.month, i.year)
	i.inlich()
}
