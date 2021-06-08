package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("HOME")
	fmt.Println(name)
}
