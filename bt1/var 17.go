package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("dang la buoi sang")
	case t.Hour() < 17:
		fmt.Println("dang la buoi chieu")
	default:
		fmt.Println("BUOI TOI ROI BE")
	}
}
