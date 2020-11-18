package main

import "fmt"

func main() {
	fmt.Print("Enter the radius of the circle:")
	var r float64
	fmt.Scan(&r)
	const pi float64 = 3.14
	area := pi * r * r
	cir := 2 * pi * r
	fmt.Printf("The area of the circle is %.2f \nThe circumference is %.2f", area, cir)
	fmt.Println("")
}
