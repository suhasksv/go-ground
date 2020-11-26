package main

import "fmt"

func main()  {
	fmt.Print("Enter number of rows:")
	var rowz int
	fmt.Scan(&rowz)
	loopmain(rowz)
}

func factorialmain(numb int)  {
	var g int = 1
	for i := 1; i <= numb; i++ {
		g *= i
	}
	fmt.Print(g)
}

func loopmain(rowz int)  {
	var num int = 1
	for i := 0; i <= rowz; i++ {
		for j := 0; j <= i; j++ {
			var g int = 1
			for f := 1; f <= num; f++ {
				g *= f
			}
			fmt.Printf("%.3d ", g)
			num++
		}
		fmt.Println("")
	}
}
