package main

import (
	"fmt"
	"os"
)

func main() {
	// Formula an = a + (n-1)d

	fmt.Println("To Find an")
	fmt.Print("Enter the value of first term OR a:")
	var a int
	fmt.Scan(&a)

	fmt.Print("Enter the value of common difference OR d:")
	var d int
	fmt.Scan(&d)

	fmt.Print("Enter the value of number of terms OR n:")
	var n int
	fmt.Scan(&n)

	fmt.Println(calc(a, d, n))
	callAgain()
}

func calc(a, d, n int) int {
	an := a + (n-1)*d
	return an
}

func callAgain() {
	fmt.Print("Do you want to calculate Roots of another quadratic equation? type 'y' to continue: ")
	var ask string
	fmt.Scan(&ask)
	if ask == "y" {
		main()
	} else {
		os.Exit(0)
	}
}
