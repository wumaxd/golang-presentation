package main

import (
	"fmt"
	"sort"
)

type User struct {
	id      uint32
	name    string
	surname string
	gopher  bool
}

type byID []User

func (a byID) Len() int           { return len(a) }
func (a byID) Less(i, j int) bool { return a[i].id < a[j].id }
func (a byID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	var user1 User

	user1.id = 5
	user1.name = "Linukk"
	user1.surname = "Torvalds"
	user1.gopher = true

	user := User{
		id:      1,
		name:    "Linus",
		surname: "Torvalds",
		gopher:  true,
	}

	user5 := User{
		id:      3,
		name:    "De",
		surname: "Torvalds",
		gopher:  true,
	}

	users := []User{user1, user, user5}

	fmt.Println("Unsorted:", users)

	sort.Sort(byID(users))
	fmt.Println("Sorted", users)

}
