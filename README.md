# GolangFundamentals
Golang fundamentals.


```go
package main

import "fmt"

func main(){
	f()
	g()
}

func f(){
	i := 1
	for i <= 10 {
		if i % 2 == 0 {
			fmt.Println(i, "Even")
		} else if i % 2 != 0 && i % 3 == 0 {
			fmt.Println(i, "Odd and divisible by 3")
		} else {
			fmt.Println(i, "Odd")
		}
		i = i + 1
	}
}

func g(){
	for i := 11; i <= 20; i++ {
		switch i{
		case 11: fmt.Println("Eleven")
		case 12: fmt.Println("Twelve")
		case 13: fmt.Println("Thirteen")
		default: fmt.Println(i) 
		}
	}
}

```
<hr>

```go
// go run Main.go
// go build -- creates .exe file at same loc you were when you used 'go build'
// go install

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
// 10100 |==| Hello, world! ||==|| This string is constant
// Enter a number: 3
// 6
```

<hr>

**Arrays, Slices, and Maps**
```go
// Arrays
func main(){
	x := [5]float64{ 98, 93, 77, 82, 83}
	var total float64 = 0

	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
	g()
}

func g(){
	my_array := [5]float64{ 1, 2, 3, 4, 5}
	var total float64 = 0

	for i := 0; i < len(my_array); i++{
		total += my_array[i]
	}
	fmt.Println(total/float64(len(my_array)))
}
```
