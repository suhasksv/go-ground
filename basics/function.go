package main

import "fmt"

func main() {
	println("Hello Go....!")
	port := 8080
	var is bool = startWebServer(port)
	println(is)
}

func startWebServer(port int) bool {
	fmt.Println("Starting server...!")
	// a bunch of code
	fmt.Println("Server started at port :", port)
	return true
}
