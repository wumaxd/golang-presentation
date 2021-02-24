package main

import "fmt"

func myDiv(a, b int) (*int, error) {
	if b == 0 {
		var err error = fmt.Errorf("You cannot divide by zero!") // HL
		return nil, err                                          // HL
	} else {
		var result int = a / b
		return &result, nil
	}
}

func main() {
	result, err := myDiv(5, 0) // HL
	if err != nil {            // HL
		fmt.Println(err)
	} else {
		fmt.Printf("The result is %d", result)
	}
}
