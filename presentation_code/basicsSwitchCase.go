package main

import "fmt"

func main() {
	var mood = "🙂"

	switch mood {
	case "🙂", "😀":
		fmt.Println("I'm happy")
	case "😞":
		fmt.Println("I'm disappointed")
	case "😢":
		fmt.Println("I'm sad")
	default:
		fmt.Println("No mood")
	}
}
