package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var g *string = new(string)
	//*g = "golang"
	//fmt.Println(*g)
	// foo()

	name := "Go"
	port := 8080
	ptr := &port

	fmt.Println("Name is '", name, "' and is a", reflect.TypeOf(name), "type")
	fmt.Println("The port is '", port, "' and is of", reflect.TypeOf(port), "type")
	fmt.Println("The 'memory address' of the port is", &port, "and is of", reflect.TypeOf(&port), "type")
	fmt.Println("The 'memory address' of the name is", &name, "and is of", reflect.TypeOf(&name), "type\n\n")

	fmt.Println("ptr =", ptr)
	fmt.Println("De-initialed ptr =", *ptr)
}
/*
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
*/