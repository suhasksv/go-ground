package main

import "fmt"

func main() {
	var prev = 0
	var curr = 1
	var nex = 0

	fmt.Print("Enter how many fibonacci numbers to be printed:")
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Print(nex, " ")
		prev = curr
		curr = nex
		nex = prev + curr
	}
	fmt.Println("")
}
