package main

import "fmt"

func main() {
	fmt.Print("Enter the length of rectange:")
	var l float64
	fmt.Scan(&l)

	fmt.Print("Enter the breadth of the rectangle:")
	var b float64
	fmt.Scan(&b)

	var area float64 = l * b
	fmt.Printf("The area of rectangle is: %.2f", area)
	fmt.Println("")
}
