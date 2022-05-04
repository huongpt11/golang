package main

import "fmt"

func main() {
	ch := make(chan int, 2) // tao 1 channel co bo dem la 2
	ch <- 1                 // gui cac gia tri 1,2 toi kenh ch
	ch <- 2                 // neu qua bo nho chuong trinh se loi
	fmt.Println(<-ch)       // in ra cac gtri tai kenh ch
	fmt.Println(<-ch)
}
