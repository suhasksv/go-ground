package main

import "fmt"

func main() {
	fmt.Print("Enter an integer to find factorial:")
	var num int
	fmt.Scan(&num)

	var g int = 1
	for i := 1; i <= num; i++ {
		g *= i
	}
	fmt.Printf("%d is the factorial of %d", g, num)
	fmt.Println("")
}
