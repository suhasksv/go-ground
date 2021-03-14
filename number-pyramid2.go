package main

import "fmt"

func main() {
	fmt.Print("Enter number of rows:")
	var rows int
	fmt.Scan(&rows)
	mainLoop(rows)
}

func mainLoop(rows int) {
	var num int = 0
	for i := -1; i < rows; i++ {
		for j := -1; j <= i; j++ {
			fmt.Printf("%d ", num)
			num++
		}
		fmt.Println("")
		if num != 0 {
			num = 0
		}
	}
}
