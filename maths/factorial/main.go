/*
package main

import "fmt"

func main() {
	var f int
	var sum big.Int = 1
	fmt.Print("Enter number to calculate factorial:")
	fmt.Scan(&f)
	for i := 1; i <= f; i++ {
		sum *= big.Int(i)
	}
	fmt.Println("Factorial of", f, "is", sum)
}
*/
package main

import (
	"fmt"
	"math/big"
)

func main() {
	var f int64
	fmt.Print("Enter a number(+ve) to calculate factorial: ")
	fmt.Scan(&f)
	var fact = new(big.Int)
	fact.MulRange(1, f)
	fmt.Println(fact)

	str := fact.String()
	sum := 0
	for _, val := range str {
		sum += int(val - '0')
	}
	fmt.Println(sum)
}
