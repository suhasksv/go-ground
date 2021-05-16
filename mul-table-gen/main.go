package main

import "fmt"
func main() {
	var num int
	fmt.Print("Enter an integer to generate the table: ")
	fmt.Scan(&num)

	calcTable(num)
}

func calcTable(num int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num * i)
	}
}
