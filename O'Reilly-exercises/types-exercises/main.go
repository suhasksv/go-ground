/*
Problem Statement -
Although overpowered for the task, you can use Go as a calculator. Write a program that computes 32,132 × 42,452
and prints it to the terminal (use the * operator for multiplication).
*/
package main

import "fmt"

func main() {
	var i int
	var f int
	fmt.Print("Enter initial range: ")
	fmt.Scan(&i)
	fmt.Print("Enter final range: ")
	fmt.Scan(&f)

	for j := i; j <= f; j++ {
		switch {
		case j%4 == 0 && j%5 == 0:
			fmt.Printf("%d --> GoPy\n", j)
		case j%4 == 0:
			fmt.Printf("%d --> Go\n", j)
		case j%5 == 0:
			fmt.Printf("%d --> Py\n", j)
		}
	}
}
