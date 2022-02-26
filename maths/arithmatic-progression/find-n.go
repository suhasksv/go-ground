package main

import (
	"fmt"
	"os"
)

func main() {
	// Formula n = ((an - a) / d) + 1

	fmt.Println("To Find an")
	fmt.Print("Enter the value of first term OR a:")
	var a int
	fmt.Scan(&a)

	fmt.Print("Enter the value of common difference OR d:")
	var d int
	fmt.Scan(&d)

	fmt.Print("Enter the value of resultant term OR an:")
	var an int
	fmt.Scan(&an)

	fmt.Println(calcs(a, d, an))
	callagain()
}

func calcs(a, d, an int) int {
	n := ((an - a) / d) + 1
	return n
}

func callagain() {
	fmt.Print("Do you want to calculate Roots of another quadratic equation? type 'y' to continue: ")
	var ask string
	fmt.Scan(&ask)
	if ask == "y" {
		main()
	} else {
		os.Exit(0)
	}
}
