package main

import "fmt"

func main() {
	pala("malayalam")
}

func pala(inputstr string) {
	var newstr string
	var reversedstr string

	var i string
	for i = range inputstr {
		if i != " " {
			newstr += i
			reversedstr += i
		}
	}

	if newstr == reversedstr {
		fmt.Println("true")
	}
	fmt.Println("false")
}
