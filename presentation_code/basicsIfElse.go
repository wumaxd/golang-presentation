package main

import "fmt"

func main() {
	var number = 5

	if number < 0 {
		fmt.Println("Number is negative")
	} else if number > 0 {
		fmt.Println("Number is positive")
	} else {
		fmt.Println("Number is zero")
	}
}
