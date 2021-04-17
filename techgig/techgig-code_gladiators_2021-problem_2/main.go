package main

import (
	"fmt"
	"math"
)

func main() {
	n := 0
	min := 0
	max := 0

	_, _ = fmt.Scan(&n)
	for j := 0; j < n; j++ {
		_, _ = fmt.Scan(&min, &max)

		k := -1
		l := -1

		for i := min; i <= max; i++ {
			if isPrime(i) {
				if k < 0 {
					k = i
				} else {
					l = i
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
	}
}

func isPrime(a int) bool {
	if a!=2 && a%2==0 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a % i == 0 {
			return false
		}
	}
	return true
}

/*
func minimum() int {
	min := arr[0]
	for i, j := range arr {
		if j < min {
			min = arr[i]
		}
		continue
	}
	return min
}

func maximum() int {
	max := arr[0]
	for i, j := range arr {
		if j > max {
			max = arr[i]
		}
		continue
	}
	return max
}
*/