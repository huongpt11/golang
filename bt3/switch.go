package main

import (
	"fmt"
	"time"
)

func main() {
	i := 6
	fmt.Println("Today is : ")
	switch i {
	case 2:
		fmt.Println("monday")
	case 3:
		fmt.Println("tuesday")
	case 4:
		fmt.Println("wesday")
	case 5:
		fmt.Println("thursday")
	case 6:
		fmt.Println("friday")
	case 7:
		fmt.Println("saturday")
	case 8:
		fmt.Println("sunday")

	}
	t := time.Now().Weekday()
	switch t {
	case time.Saturday, time.Sunday:
		fmt.Println("it's a weekend")
	default:
		fmt.Println("it's a weekday")

	}
	t1 := time.Now()
	switch {
	case t1.Hour() < 12:
		fmt.Println("morning")
	case t1.Hour() < 19:
		fmt.Println("afternoon")
	default:
		fmt.Println("EVENING")

	}
	compare := func(i interface{}) { // khai bao mot interface rong, no co the truyen vao bat ki kieu du lieu nao
		fmt.Printf("Type=%T, Value= %v", i, i)
		switch t := i.(type) { // la cu phap xet kieu gia tri co ban cua interface
		case bool: // bool gom hai gia tri la "true " or "false"
			fmt.Println("-->i'm a bool")
		case int:
			fmt.Println("-->i'm a int")
		case string:
			fmt.Println("-->i'm a string")
		default:
			fmt.Println("-->i don't know %T", t)
		}
	}
	compare(1)
	compare("hello word")
	compare("true")
}
