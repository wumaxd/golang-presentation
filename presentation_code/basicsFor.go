package main

import "fmt"

func main() {
	var sum = 0
	for i := 0; i <= 3; i++ { // HL
		sum += i
	}

	for sum < 100 { // HL
		sum += 1
	}
	fmt.Println("Summe:", sum)

	// for { break } // HL

	for index, value := range []string{"a", "b", "c", "d"} { // HL
		fmt.Printf("Index: %d - Value: %s\n", index, value)
	}
}
