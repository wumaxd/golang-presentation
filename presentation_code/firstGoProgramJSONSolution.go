package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID      uint32
	Name    string
	Surname string
	gopher  bool
}

func main() {
	user := &User{
		ID:      1,
		Name:    "Alan",
		Surname: "Turing",
		gopher:  true,
	}

	userJSON, _ := json.Marshal(user)
	fmt.Println(string(userJSON))

	response := User{}
	json.Unmarshal([]byte(userJSON), &response)
	fmt.Println(response.Name)
}
