package main

import "fmt"

func main() {

	var j int
	for j < 5 {
		println(j)
		j++
		if j == 3 {
			break
		}
		println("continuing...")
	}

	for i := 0; i < 20; i++ {
		fmt.Printf("The number is: %d\n", i)
	}

	for i := 0; i <= 5; i++ {
		println(i)
	}

	//for true {
	//	i := 0
	//	i++
	//	println(i)
	//}

	slice := []int{1, 2, 3}

	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	knownPorts := map[string]int{"http": 8080, "https": 300}

	for n, p := range knownPorts {
		println(n, p)
	}
}
