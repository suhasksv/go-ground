package main

import (
	"fmt"
)

const color string = "blue"
var i int = 0

func main() {
	var guess string
	for i <= 5 {
		fmt.Print("Guess my favourite color: ")
		if _, err := fmt.Scanln(&guess); err != nil {
			fmt.Printf("%s", guess)
			fmt.Println("")
			return
		}
		if color == guess {
			fmt.Println("You guessed my favourite color")
			break
		} else {
			fmt.Printf("%s is not my favourite color, try again...😀", guess)
			continue
			i++
		}
	}
	fmt.Printf("%s is my favourite color", color)
}
