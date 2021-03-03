package main

import "fmt"

func f(from chan string) {
	for i := 0; i < 3; i++ {
		from <- fmt.Sprint(i)
	}
}

func main() {
	messages := make(chan string)
	go f(messages)

	msg := <-messages
	fmt.Println("First:", msg)
	msg = <-messages
	fmt.Println("Second:", msg)
	msg = <-messages
	fmt.Println("Third:", msg)
}
