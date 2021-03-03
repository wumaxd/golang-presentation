package main

import (
	"fmt"
	"time"
)

func main() {
	resultsChan := make(chan string)

	select {
	case result := <-resultsChan:
		fmt.Println(result)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}
}
