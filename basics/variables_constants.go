package main

import "fmt"

const (
	f = 10
)

const (
	first  = iota
	second = iota
	third = iota + 3
	forth = 10 >> iota
	fifth = 4 << iota
	sixth = iota + 12
	seventh = iota - 13
	eighth = iota ^ 3
)

const (
	a = iota + 1
	b
	c
	d
	e
)

func main() {
	// declearing a variable var <variable_name> <data_type> = <value>(optional)
	variable()
	fmt.Println()
	constant()
	iotaBlock()
	spePow()
}

func variable() {
	var x int = 10
	var y float32 = 43.5
	var i, j, f = 11, 13.6, "go"
	var ran string = "golang"

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("i, j, f:", i, j, f)
	fmt.Println("ran:", ran)
}

func constant() {
	const x = 10
	fmt.Println("x:", x)
	fmt.Println("x + 10:", x + 1.6)
}

func iotaBlock() {
	iotaconst := []int{first, second, third, forth, fifth, sixth, seventh, eighth}
	for i, j := range iotaconst {
		fmt.Printf("%d iota value = %d", i, j)
		fmt.Println("")
	}
}

func spePow() {
	iotaconst := []int{a, b, c, d, e}

	for i, j := range iotaconst {
		fmt.Printf("%d value: %d", i, j)
		fmt.Println("")
	}
}
