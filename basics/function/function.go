package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Go....!")
	port := 8080
	_, err := startWebServer(port)
	fmt.Println(port, err)
}

func startWebServer(port int) (int, error)  {
	fmt.Println("Starting server...!")
	// a bunch of code
	fmt.Println("Server started at port :", port)
	return port, nil
}
