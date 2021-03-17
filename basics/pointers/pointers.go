package main

import "fmt"

func main() {
	var g *string = new(string)
	*g = "golang"
	fmt.Println(*g)
	foo()
}
func foo() {
	g := "google"
	prma := &g
	fmt.Println(prma, *prma)

	g = "python"
	fmt.Println(prma, *prma)

	var boo *int = new(int)
	fmt.Println(boo)

	*boo = 10
	fmt.Println(*boo)

	prma1 := &boo
	fmt.Println(prma1, *prma1, "\n", prma, *prma)
}
