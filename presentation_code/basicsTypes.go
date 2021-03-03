package main

import "fmt"

func main() {
	var myIntVar int = 23
	var myFloatVar float64 = 3.14
	var myString string = "Hello World"
	var myBool bool = true
	var myComplexVar complex128 = 5 + 3i

	fmt.Println("Integer:", myIntVar)
	fmt.Println("Float:", myFloatVar)
	fmt.Println("String:", myString)
	fmt.Println("My bool is", myBool)
	fmt.Println("My complex number is", myComplexVar)
}
