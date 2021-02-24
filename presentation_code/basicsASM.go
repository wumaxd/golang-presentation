package main

import "fmt"

func main() {
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array:", b)

	s := make([]string, 2)
	fmt.Println("empty slice:", s)
	s[0] = "a"
	s[1] = "b"
	s = append(s, "e", "f")
	fmt.Println("slice:", s)

	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)
	fmt.Println("map entry:", m["k1"])
}
