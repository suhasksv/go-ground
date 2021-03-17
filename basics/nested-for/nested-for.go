package main

import "fmt"

func main() {
	var cnt int = 1
	for i := 10; i > 5; i-- {
		for j := 5; j > 2; j-- {
			fmt.Printf("%.2d This is a nested for loop \n", cnt)
			cnt++
		}
	}
	fmt.Println("")
}
