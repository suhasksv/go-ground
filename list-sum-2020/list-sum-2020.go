package main

import "fmt"

var years = []int{2000, 20, 1980, 40, 9000, 49, 100}
var se []int
var i int

func main() {
	for _, f := range years {
		for _, k := range years {
			if notIn(f, k) == true {
				if f + k == 2020 {
					fmt.Printf("%d + %d = 2020\n", f, k)
					se = append(se, f, k)
				}
			}
			continue
		}
	}
}

func notIn(a, b int) bool {
	for _, i = range se {
		if a != i && b != i {
			return true
			continue
		}
		return false
	}
	return true
}



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