package main

import "fmt"

func main() {
	var arr = []int{10, 20, 400, 300, 80000, 100, 200, 90, 1}
	fmt.Println("Given array:", arr)
	fmt.Println("Second Highest number in the given array:", _2max(arr))
	fmt.Println("Second Lowest number in the given array:" , _2min(arr))
}

func _2max(arr []int) int {
	max1 := 0
	max := 0
	
	for _, i := range arr {
		if i > max {	
			if max1 < max {
				max1 = max
			}
			max = i
		} else if i > max1 {
			max1 = i
		}
	}
	return max1
}

func _2min(arr []int) int {
	var min1 = arr[0]
	min := arr[0]

	for _, i := range arr {
		if i < min {
			if min1 > min {
				min1 = min
			}
			min = i
		} else if i < min1 {
			min1 = i
		}
	}
	return min1
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
