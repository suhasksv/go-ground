package main

import "fmt"

func main()  {
	fmt.Println(pala("malayalam"))
}

func pala(input_str string) {
	var new_str string
	var reversed_str string
	var i string

	for i = range input_str {
		if i != " " {
			new_str +=  i
			reversed_str += i
		}
	}

	if new_str == reversed_str {
		return "true"
	}
	return "false"
}
