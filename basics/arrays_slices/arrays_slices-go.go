package main

import "fmt"

func main() {
	// Three(3) element arrays, long syntax
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	// fmt.Println(arr[2])

	// Three(3) element arrays, short syntax
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	for i, j := range arr2 {
		fmt.Printf("%d The value of the Array is : %d", i, j)
		fmt.Println("")
	}

	sliceArray()
}

func sliceArray() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:]

	fmt.Printf("arr : %d, slice %d", arr, slice)
	fmt.Println("")

	slice2 := []float32{1, 2, 3, 4}
	fmt.Println(slice2)

	slice2 = append(slice2, 5) // append method - To add an element to an existing Array
	fmt.Println(slice2) // Output - Notice that an element is added

	// Indexing an Array
	ss2 := slice2[1:3]
	ss3 := slice2[3:]
	ss4 := slice2[:3]
	fmt.Println(ss2, ss3, ss4)
}
