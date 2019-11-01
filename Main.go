package main

import "fmt"

func main(){
	f()
	g()
}

func f(){
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}
}

func g(){
	for i := 6; i <= 10; i++ {
		fmt.Println(i)
	}
}