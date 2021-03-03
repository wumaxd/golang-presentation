package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go f("goroutine")

	f("direct")

	fmt.Println("done")
}
