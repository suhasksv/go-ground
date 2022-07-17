package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var initial, capacity, water, time int
		fmt.Scan(&initial, &capacity, &water, &time)

		j := initial + (time * water)
		switch {
		case j == capacity:
			fmt.Print("filled\n")
		case j < capacity:
			fmt.Print("unfilled\n")
		case j > capacity:
			fmt.Print("overflow\n")
		}
	}
}
