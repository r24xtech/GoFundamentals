package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	// To create a dynamic array
	arr := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		// Scans a line from Stdin(console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
			arr = append(arr, text)
		} else {
			break
		}
	}
	// Use collected inputs
	fmt.Println(arr)
}

/*
 Three ways of taking input
   1. fmt.Scanln(&input)
   2. reader.ReadString()
   3. scanner.Scan()

   Here we recommend using bufio.NewScanner
*/