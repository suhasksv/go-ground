package main // Ever Go source code starts with package
import "fmt"

/*
The main function, when part of the package main, identifires the entry
point of the application
*/
func main() {
	/*
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


	var firstName *string = new(string)
	*firstName = "Golang"
	fmt.Println(*firstName)

	a := 20
	b := 10

	fmt.Printf("Result of a + b = %d\n", a + b)
	fmt.Printf("Result of a - b = %d\n", a - b)
	fmt.Printf("Result of a * b = %d\n", a * b)
	fmt.Printf("Result of a / b = %d\n", a / b)
	fmt.Printf("Result of a % b = %d", a % b)

	fmt.Println(a % b)
	a := 20
	b := 10
	fmt.Printf("a == b %t\n", a == b)
	fmt.Printf("a != b %t\n", a != b)
	fmt.Printf("a > b %t\n", a > b)
	fmt.Printf("a < b %t\n", a < b)
	fmt.Printf("a >= b %t\n", a >= b)
	fmt.Printf("a < b %t", a <= b)
	fmt.Println()
	 */
	//const a int = 20
	//const b int = 10
	//
	//fmt.Printf("Result of a & b = %d\n", a & b)
	//fmt.Printf("Result of a | b = %d\n", a | b)
	//fmt.Printf("Result of a ^ b = %d\n", a ^ b)
	//fmt.Printf("Result of a >> b = %d\n", a >> b)
	//fmt.Printf("Result of a << b = %d\n", a << b)
	//fmt.Printf("Result of a &^ b = %d", a &^ b)

	//var a int32
	//var b int32 = 10
	//
	//a = b
	//fmt.Println(a)
	//
	//a += b
	//fmt.Println(a)
	//
	//a -= b
	//fmt.Println(a)
	//
	//a *= b
	//fmt.Println(a)
	//
	//a /= b
	//fmt.Println(a)
	//
	//a %= b
	//fmt.Println(a)

	var a int16 = 4
	b := &a
	fmt.Println(b)
	*b = 7
	fmt.Println(a)


}
/*
func foo() {
	years := []int{10, 100, 200, 2000, 20}
	for i := range years {
		fmt.Println(i)
	}
}
*/
