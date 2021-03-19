package main

import "fmt"

func main() {

	swap_thr_var()
	swap_add()
	swap_mul()
}

func swap_thr_var() {
	var temp float32
	var a float32
	var b float32

	fmt.Print("Enter the value of a: ")
	fmt.Scan(&a)

	fmt.Print("Enter the value of b: ")
	fmt.Scan(&b)

	temp = a
	a = b
	b = temp

	fmt.Printf("The value of a is %.2f and the value of b is %.2f", a, b)
	fmt.Println("")
}

func swap_add() {
	var a float32
	var b float32

	fmt.Print("Enter the value of a: ")
	fmt.Scan(&a)

	fmt.Print("Enter the value of b: ")
	fmt.Scan(&b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("The value of a is %.2f and the value of b is %.2f", a, b)
	fmt.Println("")
}

func swap_mul() {
	var a float32
	var b float32

	fmt.Print("Enter the value of a: ")
	fmt.Scan(&a)

	fmt.Print("Enter the value of b: ")
	fmt.Scan(&b)

	a = a * b
	b = a / b
	a = a / b

	fmt.Printf("The value of a is %.2f and the value of b is %.2f", a, b)
	fmt.Println("")
}
