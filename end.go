package main

import "fmt"

type input struct {
	month int
	year  int
}

// func (i input) namnhuan(day int) {
// 	switch i.month {
// 	case 1, 3, 5, 7, 8, 10, 12:
// 		day = 31
// 	case 2:
// 		if i.year%4 == 0 && i.year%100 != 0 || i.year%400 == 0 {
// 			day = 29
// 		} else {
// 			day = 28
// 		}
// 	default:
// 		day = 30

// 	}
// }
func (i input) inlich(day int) {
	switch i.month {
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
	d := 1 // ngay dau tien bat dau tinh tu 1/1/1
	a := (14 - i.month) / 12
	y := i.year - a
	m := i.month + 12*a - 2
	dow := (d + y + y/4 - y/100 + y/400 + (31*m)/12) % 7 // cong thuc zeller,tinh thu cua ngay dau tien trong thang
	fmt.Printf("\n%3s%3s%3s%3s%3s%3s%3s\n", "S", "M", "T", "W", "T", "F", "S")
	for i := 0; i < 7; i++ {
		if i < dow {
			fmt.Printf("  ")
		} else {
			fmt.Printf("%3d", d)
			d++
			fmt.Printf(" ")
		}

		for i := 0; i < 7; i++ {
			if d <= day {
				fmt.Printf("%3d", d)
				d++
				fmt.Printf("\n")
			} else {
				fmt.Printf("  ")
			}
		}

	}
}
func main() {
	i := input{1, 1}
	fmt.Printf("calender of %v %v is:", i.month, i.year)
	i.inlich(31)
}
