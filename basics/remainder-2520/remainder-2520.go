package main

import "fmt"

func main() {
	const num = 2520
	for i := 1; i <= 9; i++ {
		fmt.Println(num%i)
	}
}
