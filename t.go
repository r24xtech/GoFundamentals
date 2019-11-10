package main

import "fmt"

func zero(xPtr *int) {
  *xPtr = 0
}

func main() {
  x := 5
  y := new(int)

  zero(&x)
  fmt.Println(x)
  fmt.Println(y)
}
