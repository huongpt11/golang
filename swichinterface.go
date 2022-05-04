package main

import (
	"fmt"
)

func test(t interface{}) { // khai bao interface trong
	switch v := t.(type) { // kiem tra kieu du lieu
	case int:
		fmt.Println(v)
	case string:
		fmt.Println(v)
	default:
		fmt.Println("i don't know!!")

	}

}
func main() {
	test(6)
	test("hello")
	test(true)
}
