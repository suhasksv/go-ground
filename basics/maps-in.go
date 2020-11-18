package main

import ("fmt")

func main() {
	n := map[string]int{"golang": 100, "python": 90}
	fmt.Println(n)
	fmt.Println(n["golang"])
	fmt.Println(n["python"])

	n["python"] = 95
	fmt.Println(n["python"])
}
