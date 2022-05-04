package main

import (
  "fmt"
  "math/cmplx"
)

func main() {
  var (
    x bool       = false
    y uint32     = 22
    z complex128 = cmplx.Sqrt(-5 + 12i)
  )
  fmt.Println("type: %T", "value:%v\n", x, x)
  fmt.Println("type: %T", "value:%v\n", y, y)
  fmt.Println("type: %T", "value:%v\n", z, z)

}
