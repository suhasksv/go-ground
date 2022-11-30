package main

import (
	"fmt"
	"main/maths/pandc/factorial"
)

func main() {
	var ask string
	var n, r int
	print("What do you want to find? : 1. permutation 2. combination :")
	fmt.Scan(&ask)

	print("Enter n and r values with space: ")
	fmt.Scan(&n, &r)
	if ask == "1" || ask == "permutation" {
		println(permutation(n, r))
	} else if ask == "2" || ask == "combination" {
		println(combinations(n, r))
	}
}

func permutation(n, r int) int {
	return factorial.Factorial(n) / factorial.Factorial(n-r)
}

func combinations(n, r int) int {
	return factorial.Factorial(n) / (factorial.Factorial(r) * factorial.Factorial(n-r))
}
