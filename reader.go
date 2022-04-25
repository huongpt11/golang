package main
// loai tham so
import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello,Reader! ") // doc du lieu tu string

	b := make([]byte, 8) // tao bo dem truyen tai co do dai 8 byte
	for {
		// mang co do dai max la 8 nen lan dau no se doc den ki tu 'R'
		// lan thu 2 se doc cac ki tu con lai la 6 ki tu
		n, err := r.Read(b) // n doc den len(b) trong b, err kiem tra error la nil hay khong
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n]) // in ra cac ki tu da duoc doc
		if err == io.EOF {
			break
		}
	}
}
