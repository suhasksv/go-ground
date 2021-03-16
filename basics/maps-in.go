package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"foo": 1, "boo":3}
	fmt.Println(m["foo"])
	fmt.Println(m)

	m["boo"] = 2
	fmt.Println(m["boo"])
	fmt.Println(m)

	delete(m, "boo")
	fmt.Println(m)
}
