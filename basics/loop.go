package main

import "fmt"

func main() {
	for_loop(32, 40)
}

func for_loop(i int, j int) {
	for f := i; f <= j, f++ {
		fmt.Println(f)
	}
}
