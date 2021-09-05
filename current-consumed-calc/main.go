package main

import "fmt"

/*
const (
	EXT_POWER_SOURCE         int = 0
	SOLAR_POWER_SOURCE       int = 1
	IN_HOUSE_SOLAR_PLANT     int = 2
	IN_HOUSE_GENERATOR_PLANT int = 3
	WIND_POW_SOURCE          int = 4
)
*/

func main() {
	// Accepting units variable a input from the person
	var units int
	fmt.Print("Enter no. of units:")
	fmt.Scan(&units)
	// Accepting hours variable a input from the person
	var hours int
	fmt.Print("Enter no. of hours:")
	fmt.Scan(&hours)

	// Number of Power Sources
	powerSourceLabels := []string{"External Power", "Solar Power", "In-house Solar Power", "In-house Geneset", "Wind Power Souce"}
	unitPriceArray := []int{12, 4, 2, 6, 5} // Cost of each power source respectively to powerSourceLabels
	powerSourcesValues := []int{}           // Array that stores all the Prices of each power source

	fmt.Println("\n\nThe Cost for Power -")
	for i := 0; i < len(powerSourceLabels); i++ { // this for loop that appends the price to powerSourceValues Array and prints the final statement
		value := calc(unitPriceArray[i], units, hours)                                            // calculates price of each case
		powerSourcesValues = append(powerSourcesValues, value)                                    // appends to prices to powerSourceValues Array
		fmt.Printf("%s @%d per unit - Rs %.2d\n", powerSourceLabels[i], unitPriceArray[i], value) // Prints the final statement
	}

	fmt.Println("\n\nRange of Prices")
	fmt.Println("The minimum cost is: Rs", min(powerSourcesValues))
	fmt.Println("The maximum cost is: Rs", max(powerSourcesValues))
}

func calc(price, units, hours int) int {
	return (price * units) * hours
}

func min(arr []int) int {
	mini := arr[0]
	for i, j := range arr {
		if j < mini {
			mini = arr[i]
		}
		continue
	}
	return mini
}

func max(arr []int) int {
	maxi := arr[0]
	for i, j := range arr {
		if j > maxi {
			maxi = arr[i]
		}
		continue
	}
	return maxi
}
