package main

import (
	"fmt"
	"os"
)

func main() {
	// feb := "february"
	_31months := []string{"january", "march", "may", "july", "august", "october", "december"}
	_30months := []string{"april", "june", "september", "november"}

	var date, year int
	var month string
	fmt.Print("Enter date DD MM YYYY: ")
	fmt.Scan(&date, &month, &year)
	if date > 31 {
		fmt.Println("Please enter correct date 1 - 31")
		os.Exit(1)
	}
	for _, i := range _31months {
		if month == i {
			days := 31 - date
			if days < 0 {
				fmt.Println("Please enter correct date respective of the month entered")
				os.Exit(1)
			} else {
				fmt.Println("Days left in ", month, "are:", days)
			}
		} else if true == true {
			for _, j := range _30months {
				if month == j {
					days := 30 - date
					if days < 0 {
						fmt.Println("Plz enter correct date respective of the month entered")
						os.Exit(1)
					} else {
						fmt.Println("Days left in ", month, "are: ", days)
						os.Exit(0)
					}
				}
			}
		} else {
			if leapyear(year) {
				days := 29 - date
				if days < 0 {
					fmt.Println("Please enter correct date respective of the month entered")
					os.Exit(1)
				} else {
					fmt.Println("Days left in", month, "are: ", days)
					//	os.Exit(0)
				}
			} else {
				days := 28 - date
				if days < 0 {
					fmt.Println("Please enter correct date with respective of the month entered")
					os.Exit(1)
				} else {
					fmt.Println("Days left in", month, "are: ", days)
				}
			}
		}
	}
}

func leapyear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	} else {
		return false
	}
}
