package main

import (
	"fmt"
	"strings"
)

func main() {
	g := 1
	for g == 1 {
		fmt.Print("Enter an phrase to create an short form, Note - Use '-' in place of <space>: ")
		var word string
		fmt.Scan(&word)
		fmt.Println(word)
		fmt.Println(abbr(word))

		var ask string
		fmt.Print("Do you want to find Another? type y to continue...! :")
		
		fmt.Scan(&ask)	
		if strings.ToLower(ask) != "y" {
			g += 1
		}
	}
}

func abbr(phrase string) string {
	phrase1 := strings.Split(phrase, "-")
	var abbre string
	for _, words := range phrase1 {
		abbre += string(words[0])
	}
	return strings.ToUpper(abbre)
}
