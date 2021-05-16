package main

import (
	"fmt"
	"math"
)

func main() {
	var num int64
	fmt.Print("Enter an integer to find if it is prime or not: ")
	fmt.Scan(&num)
	if prime(num) {
		fmt.Println(num, "is a prime number")
	} else {
	fmt.Println(num, "is a not prime number")
	}
}

func prime(num int64) bool {
	if num != 2 && num % 2 == 0 {
		return false
	}
	
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if int(num) % i == 0 {
			return false
		}
	}
	return true
}
