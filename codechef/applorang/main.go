package main

import "fmt"

func main() {
	var rup int
	fmt.Scan(&rup)

	var apple, orange int
	fmt.Scan(&apple, &orange)

	if apple+orange <= rup {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
