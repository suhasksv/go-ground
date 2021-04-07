package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64

	fmt.Print("Enter the base number: ")
	fmt.Scan(&a)

	fmt.Print("Enter the exponent number: ")
	fmt.Scan(&b)

	fmt.Printf("The exponent of %.2f and %.2f is %.2f", a, b, exponent(a, b))
	fmt.Println()
}

func exponent(a float64, b float64) float64 {
	var expo float64 = math.Pow(a, b)
	return expo
}
