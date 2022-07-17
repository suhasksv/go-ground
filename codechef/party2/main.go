package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		var friends, price, money int
		fmt.Scan(&friends, &price, &money)

		if friends * price <= money {
			fmt.Print("yes\n")
		} else {
			fmt.Print("no\n")
		}
	}
}
