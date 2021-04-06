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
	//
	//var a int16 = 4
	//b := &a
	//fmt.Println(b)
	//*b = 7
	//fmt.Println(a)
	//
	//var arr = []uint{1, 2, 3, 4, 5, 6}
	//for i, j := range arr {
	//	fmt.Println(i, j)
	//}

	//for i, j := range "golang" {
	//	fmt.Printf("The index at %U is %d\n", j, i)
	//}

	var map1 = map[string]string{"Go": ".go", "Python": ".py", "Java": ".java", "Unix exec": ".exec", "App": ".app", "compiled c unix": ".so"}

	for i, j := range map1 {
		fmt.Printf("%s --> %s\n", i, j)
	}

	chn1 := make(chan int)
	go func() {
		chn1 <- 10
		chn1 <- 100
		chn1 <- 1000
		chn1 <- 10000
		close(chn1)
	}()

	for i := range chn1 {
		fmt.Println(i)
	}
	switch day := 4; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}
	var i int = 2
	switch {
	case i == 1:
		fmt.Println("Python")
	case i == 2:
		fmt.Println("Go")
	case i == 3:
		fmt.Println("Java")
	case i == 4:
		fmt.Println("C")
	case i == 5:
		fmt.Println("")
	}
}


/*
func foo() {
	years := []int{10, 100, 200, 2000, 20}
	for i := range years {
		fmt.Println(i)
	}
}
*/
