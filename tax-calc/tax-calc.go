package main

import (
	"fmt"
	"os"
)

var (
	salary float64
	tax float64 = 0
	ask string
	a float64 = 250000
)

func main() {
	fmt.Print("Enter your annual Salary: ")
	tax = 0
	fmt.Scan(&salary)
	if salary <= 250000 {
		fmt.Println("Hooray you no need to pay any tax...! Enjoy your Life...!")
		callagain()
	}
	fmt.Printf("Your Salary or total income is %.2f\nTax charged by the government is %.2f and your saving are %.2f\n\n", salary, switchStatements(), salary - tax)
	callagain()
}

func _250KTo500K(salary float64) float64 {
	salary -= a
	tax += a * 0.05
	return tax
}

func _500KTo750K(salary float64) float64 {
	var b = _250KTo500K(salary)
	salary -= a
	tax += b + a * 0.10
	return tax
}

func _750KTo1000K(salary float64) float64 {
	var b = _500KTo750K(salary)
	var a float64 = 250000
	salary -= a
	tax += b + a * 0.15
	return tax
}

func _1000KTo1500K(salary float64) float64 {
	var b = _750KTo1000K(salary)
	salary -= a
	tax += b + a * 0.20
	return tax
}

func _1500KGreater(salary float64) float64 {
	var b = _750KTo1000K(salary)
	salary -= a
	tax += b + salary * 0.30
	return tax
}

func switchStatements() float64 {
	switch {
	case 250001 <= salary && salary <= 500000:
		return _250KTo500K(salary)
	case 500001 <= salary && salary <= 750000:
		return _500KTo750K(salary)
	case 750001 <= salary && salary <= 1000000:
		return _750KTo1000K(salary)
	case 1000001 <= salary && salary <= 1500000:
		return _1000KTo1500K(salary)
	case 1500001 <= salary:
		return _1500KGreater(salary)
	}
	return 0
}

func callagain() {
	fmt.Print("Do you want calculate another time? Type 'y' to continue: ")
	fmt.Scan(&ask)
	if ask == "y" {
		main()
	} else {
		os.Exit(0)
	}
}
