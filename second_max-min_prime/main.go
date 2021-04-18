package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var min int
	var max int
	var arr []int

	fmt.Print("Enter no of test cases: ")
	_, _ = fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Enter minimum and maximum value: ")
		_, _ = fmt.Scan(&min, &max)

		for j := min; j <= max; j++ {
			if isPrime(j) {
				arr = append(arr, j)
			}
		}

		if len(arr) == 0 {
			fmt.Println(-1)
		} else if min == max {
			fmt.Println(0)
		} else {
			fmt.Println(arr[len(arr)-1] - arr[1])
		}
	}
}

func isPrime(a int) bool {
	if a != 2 && a%2 == 0 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

/*
for j := 0; j <= max; j++ {
			if isPrime(j) {
				if k < 0 {
					k = i
				} else {
					l = i
				}
			}
		}
for f := 0; f <= max; f++ {
			if isPrime(f) {
				if f > k {
					k = f
					break
				}
			}
		}

		for d := max; d > 0 ; d-- {
			if isPrime(d) {
				if d < l {
					l = d
					break
				}
			}
		}

		if k < 0 {
			fmt.Println(-1)
		} else if l < 0 {
			fmt.Println(0)
		} else {
			fmt.Println(l - k)
		}
*/
