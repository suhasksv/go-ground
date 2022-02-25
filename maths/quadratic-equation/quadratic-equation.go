package main

import (
	"fmt"
	"math"
	"os"
)

var (
	a     float64
	b     float64
	c     float64
	d     float64
	alpha float64
	beta  float64
	i     int
	ask   string
)

func main() {
	fmt.Print("Enter the value of a: ")
	fmt.Scan(&a)

	fmt.Print("Enter the value of b: ")
	fmt.Scan(&b)

	fmt.Print("Enter the value of c: ")
	fmt.Scan(&c)
	i = dChk(a, b, c)
	alpha, beta = calcAlphaBeta(a, b, c)
	switchStatement(i, alpha, beta)
	callAgain()
}

func dChk(a, b, c float64) int {
	d = (b * b) - (4 * a * c)

	if d > 0 {
		i = 1
	} else if d == 0 {
		i = 2
	} else if d < 0 {
		i = 3
	}
	return i
}

func calcAlphaBeta(a, b, c float64) (float64, float64) {
	d = (b * b) - (4 * a * c)
	alpha = (-b + math.Sqrt(d)) / (2 * a)
	beta = (-b - math.Sqrt(d)) / (2 * a)
	return alpha, beta
}

func switchStatement(i int, alpha, beta float64) {
	switch i {
	case 1:
		fmt.Printf("\nRoots are Real\nThe value of alpha = %.2f and beta = %.2f\n", alpha, beta)
	case 2:
		fmt.Printf("\nRoots are Real and equal\nThe value of alpha and beta = %.2f\n", alpha)
	case 3:
		fmt.Printf("\nRoots are imaganary\nThe value of alpha = %.2f and beta = %.2f\n", alpha, beta)
	}
}

func callAgain() {
	fmt.Print("Do you want to calculate Roots of another quadratic equation? type 'y' to continue: ")
	fmt.Scan(&ask)
	if ask == "y" {
		main()
	} else {
		os.Exit(0)
	}
}
