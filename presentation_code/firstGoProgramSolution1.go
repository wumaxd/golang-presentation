package main

import "fmt"

type User struct {
	id      uint32
	name    string
	surname string
	gopher  bool
}

func main() {
	var user1 User

	user1.id = 1
	user1.name = "Linus"
	user1.surname = "Torvalds"
	user1.gopher = true

	fmt.Println(user1)
}
