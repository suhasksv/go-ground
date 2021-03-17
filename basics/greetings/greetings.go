package main

import "fmt"

func main() {
	fmt.Print("Please enter your name.... :")
	var name string

	fmt.Scanln(&name)
	fmt.Printf("Hello %s \nI am Go from Google...! \nThanks for pinging me..", name)
	fmt.Println("")
}
