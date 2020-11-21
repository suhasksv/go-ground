package main

import "fmt"

func main() {
	fmt.Print("Enter the year:")
	var year int
	fmt.Scan(&year)
	leapyear(year)
}

func leapyear(year int) {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Printf("%d is a leap year", year)
		fmt.Println("")
	} else {
		fmt.Printf("%d is not a leap year", year)
		fmt.Println("")
	}
}
