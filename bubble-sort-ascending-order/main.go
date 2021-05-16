package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{4, 2, 3, 1, 10, 2804, 9832945, 2, 2974, 2974, -24867, -6376, 356672, 2641, 2776, -2874, 29748}
	start := time.Now()
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				//swap
				t := arr[i]
				arr[i] = arr[j]
				arr[j] = t
			}
		}
	}
	fmt.Println(arr)
	elapsed := time.Since(start)
	fmt.Printf("page took %s", elapsed)
}
