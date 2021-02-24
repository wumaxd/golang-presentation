package main

import "fmt"

type User struct {
	id      uint32
	name    string
	surname string
	gopher  bool
}

func main() {
	user := User{
		id:      1,
		name:    "Linus",
		surname: "Torvalds",
		gopher:  true,
	}

	fmt.Printf("%v", user)
}
