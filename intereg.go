package main

import "fmt"

func main() {
	var m interface{} = 6 // khai bao interface trong gan gia tri bang 6
	s := m.(int)          // giao dien m giu kieu int , gan gia tri cua m vao s
	fmt.Println(s)        // in ra gia tri cua s
	t, ok := m.(int)      // kiem tra m co thuoc kieu int khong, neu co thi t la gia tri co ban va ok la true, neu nguoc lai thi ok la false va s co gia tri 0
	fmt.Println(t, ok)
	f, ok := m.(string) // kiem tra f co thuoc kieu int hay k
	fmt.Println(f, ok)

}
