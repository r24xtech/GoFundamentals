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