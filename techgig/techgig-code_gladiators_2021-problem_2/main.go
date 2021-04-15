package main

import "fmt"

var arr []int

func main() {
	var n int
	var min = 0
	var max = 0

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&min, &max)
		calc(min, max)

		if max == 0 {
			if isPrime(min) == true {
				fmt.Println(min)
			}
		} else if len(arr) == 0 {
			fmt.Println(-1)
		} else {
			fmt.Println(maximum() - minimum())
		}
		min, max = 0, 0
		arr = nil
	}
}

func calc(min, max int) {
	for i := min; i <= max; i++ {
		if isPrime(i) {
			arr = append(arr, i)
		}
	}
}

func isPrime(a int) bool {
	var cnt = 0
	for i := 1; i <= a/2; i++ {
		if a % i == 0 {
			cnt++
		}
		continue
	}
	if cnt == 1 {
		return true
	}
	return false
}

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
