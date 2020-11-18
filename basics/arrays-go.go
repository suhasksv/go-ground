package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(arr)

	arr = append(arr, 4, 5)
	fmt.Println(arr)

	r1 := arr[:]
	r2 := arr[1:]
	r3 := arr[:4]
	fmt.Println(r1, "\n", r2, "\n", r3)
	a := 10
	fmt.Print(a)
}
