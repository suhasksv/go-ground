package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var Na, Nb, Nc int
		fmt.Scan(&Na, &Nb, &Nc)

		k := Nb + Nc

		switch {
		case Na > k:
			fmt.Print("yes\n")
		case Na < k:
			fmt.Print("no\n")
		}
	}
}
