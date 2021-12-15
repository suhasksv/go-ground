package main

import "fmt"

func main() {
	arr := []int{2, 2004, 2847, 92, -453, 2901, 20985, -2947}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				// swap
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println(err)
}
