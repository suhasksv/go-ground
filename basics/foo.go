package main // Ever Go source code starts with package

import (
	"fmt"
	"math" // Go imports a module i.e, "fmt"
)

/*
The main function, when part of the package main, identifires the entry
point of the application
*/
func main() {
	a := 10.0

	a = math.Pow(a, 2)
	
	var b = "Tuesday"

	//var x int

	x := int(15)

	var str string

	str = "I am a random string"

	fmt.Printf("hello %.2f th of the day is %s of the month %d in the %s\n", a, b, x, str)
	fmt.Println("world")
	fmt.Print("hello", a, "th of the day is", b, "\n")
	fmt.Println("world")
	foo()
}

func foo() {
	years := []int{10, 100, 200, 2000, 20};
	for i := range years {
		fmt.Println(i)
	}
}

