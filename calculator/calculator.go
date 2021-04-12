package main

import (
	"fmt"
	"os"
)

var (
	num1 float64
	num2 float64
	op string
)



func main() {
	fmt.Println("Welcome to Quantum Tricks Go Infinity calculator!")

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	for true {
		fmt.Print("Enter the operation that you want to perform (addition: add, subtraction: sub, multiplication: mul, division: div): ")
		fmt.Scan(&op)

		if op == "exit" || op == "break" || op == "quit" {
			fmt.Printf("Your final answer is %.2f", num1)
			fmt.Println("Thanks for choosing and using Quantum Tricks Go Infinity Calculator")
			os.Exit(0)
		} else if op != "add" || op != "sub" || op != "mul" || op != "div" {
			fmt.Printf("%s is not yet supported at QT Go infinity calculator\nPlease try another operation\n", op)
			continue
		} else {
			fmt.Print("Enter second number: ")
			fmt.Scan(&num2)
			num1 = switchSt(op)
			fmt.Println(num1)
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

func switchSt(op string) float64 {
	switch op {
	case "add":
		return add(num1, num2)
	case "sub":
		return sub(num1, num2)
	case "mul":
		return mul(num1, num2)
	case "div":
		return div(num1, num2)
	}
	return 0
}
