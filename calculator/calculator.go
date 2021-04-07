package main

import "fmt"

var (
	num1 float64
	num2 float64
	op string
)



func main() {
	fmt.Println("Hello welcome to Quantum Tricks Go Infinity calculator!")

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)
	num1 = float64(num1)

	for true {
		fmt.Print("Enter the operation that you want to perform (addition: add, subtraction: sub, multiplication: mul, division: div): ")
		fmt.Scan(&op)

		fmt.Print("Enter second number: ")
		fmt.Scan(&num2)

		if op == "add" {
			fmt.Println(add(num1, num2))
		} else if op == "sub" {
			fmt.Println(sub(num1, num2))
		} else if op == "mul" {
			fmt.Println(mul(num1, num2))
		} else if op == "div" {
			fmt.Println(div(num1, num2))
		} else {
			fmt.Println("Operation is not supported at QT Infinity Go Calculator")
		}
	}
}

func add(a float64, b float64) float64 {
	return a + b
}

func sub(a float64, b float64) float64 {
	return a - b
}

func mul(a float64, b float64) float64 {
	return a * b
}

func div(a float64, b float64) float64 {
	return a / b
}
