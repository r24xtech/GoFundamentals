# GolangFundamentals
Golang fundamentals.






```go
// go run Main.go
// go build ("stops at the folder Main.go is on") 
// (W10) creates .exe file at same loc you were when you used 'go build'
// go install
// bin/firstapp

package main

import "fmt"

var (
	a = 5
	b = 10
	c = 50
)

func main(){
	const z string = "This string is constant"
	var i int = 80
	e := 20
	var t int
	t = 10000
	var my_string string = "Hello, world!"
	
	fmt.Println(i+e+t, "|==|", my_string, "||==||", z)
	f()
}

func f(){
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
```
