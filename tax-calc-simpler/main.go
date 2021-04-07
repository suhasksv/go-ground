package main

import (
	"fmt"
	"os"
)

var (
	orisalary float64
	salary float64
	ask string
)

func main() {
	fmt.Print("Enter your annual Salary: ")
	fmt.Scan(&orisalary)
	salary = orisalary
	if salary <= 250000 {
		fmt.Println("Hooray you no need to pay any tax...! Enjoy your Life...!")
		callAgain()
	}
	fmt.Printf("Your Salary or total income is %.2f\nTax charged by the government is %.2f and your saving are %.2f\n\n", orisalary, taxCalc(), orisalary - taxCalc())
	callAgain()
}

func taxCalc() float64 {
	slab := idSlab(salary)
	var a float64 = 250000
	switch slab {
	case 0.05:
		salary -= a
		return a * slab
	case 0.10:
		salary -= 2 * a
		return a * 0.05 + a * slab
	case 0.15:
		salary -= 3 * a
		return a * 0.05 + a * 0.10 + a * slab
	case 0.20:
		salary -= 4 * a
		return a * 0.05 + a * 0.10 + a * 0.15 + a * slab
	case 0.30:
		salary -= 5 * a
		return a * 0.05 + a * 0.10 + a * 0.15 + a * 0.20 + a * slab
	}
	return 0
}

func idSlab(salary float64) float64 {
	switch {
	case 250001 <= salary && salary <= 500000:
		return 0.05
	case 500001 <= salary && salary <= 750000:
		return 0.10
	case 750001 <= salary && salary <= 1000000:
		return 0.15
	case 1000001 <= salary && salary <= 1500000:
		return 0.20
	case 1500001 <= salary:
		return 0.30
	}

	return 0
}

func callAgain() {
	fmt.Print("Do you want calculate another time? Type 'y' to continue: ")
	fmt.Scan(&ask)
	if ask == "y" {
		main()
	} else {
		os.Exit(0)
	}
}
