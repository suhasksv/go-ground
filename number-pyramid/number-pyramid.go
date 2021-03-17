package main

import "fmt"

func main()  {
	fmt.Print("Enter number of rows:")
	var row int
	fmt.Scan(&row)

	loop(row)
}

func loop(row int)  {
	var num int = 1
	for i := 0; i <= row; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%.2d ", num)
			num++	
		}
		println("")
	}
}
