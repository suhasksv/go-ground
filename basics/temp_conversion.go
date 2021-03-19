package main

import (
	"fmt"
)

var (
	f float64
	c float64
	ask string
	ask2 string
)

func main() {
	for true {
		fmt.Print("Enter 'c-f' to convert from Celsius to Fahrenheit and 'f-c' to convert from Fahrenheit to Celsius: ")
		fmt.Scan(&ask)
		if ask == "c-f" {
			h, j := temp_c_to_f()
			fmt.Printf("%.2f˚C = %.2f˚F", h, j)
			fmt.Println()
		} else if ask == "f-c" {
			a, b := temp_f_to_c()
			fmt.Printf("%.2f˚F = %.2f˚C", a, b)
			fmt.Println()
		} else {
			main()
		}

		fmt.Print("\nDo you want to calculate again? enter 'y' to continue: ")
		fmt.Scan(&ask2)
		if ask2 != "y" {
			break
		}
		fmt.Println()
		continue
	}
}

func temp_c_to_f() (float64, float64) {
	fmt.Print("Enter temperature in Celsius: ")
	fmt.Scan(&c)

	f = (c * 9.0 / 5.0) + 32.0
	return c, f
}

func temp_f_to_c() (float64, float64) {
	fmt.Print("Enter temperature in Fahrenheit: ")
	fmt.Scan(&f)

	c = (5.0 * f - 32.0) / 9.0
	return f, c
}
