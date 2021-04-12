package main

import "fmt"

var (
	t int
	x int
	y int
	n int
)

func main() {
	fmt.Print("Enter number of cases: ")
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		fmt.Scan(&x, &y, &n)
		w := 0
		days := 1
		for n - w > x {
			w += x
			w -= y
			days++
			continue
		}
		fmt.Println(days)
	}
}
