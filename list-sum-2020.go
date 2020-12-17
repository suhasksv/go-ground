package main

import (
	"fmt"
	//"reflect"
)

// func itemExists(arrayType interface{}, item interface{}) bool {
// 	arr := reflect.ValueOf(arrayType)

// 	if arr.Kind() != reflect.Array {
// 		panic("Invalid data-type")
// 	}

// 	for i := 0; i < arr.Len(); i++ {
// 		if arr.Index(i).Interface() == item {
// 			return true
// 		}
// 	}

// 	return false
// }

func main()  {
	var years []int{2000, 20, 1980, 40, 9000, 49, 100}
	// se := []int{}
	for i := range years {
		for j := range years {
			// if itemExists(i, se) && itemExists(j, se) {
			// 	continue;
			// } else 
			if i + j == 2020 {
				fmt.Printf("%d + %d = 2020", i, j)
				fmt.Println()
				// se = append(se, i, j)
			} else {
				continue;
			}
		}
	} 
}
