package main

import (
	"encoding/json"
	"io/ioutil"
)

func findNum(input interface{}) []int {
	var numbers []int

	switch input := input.(type) {
	case []interface{}:
		for _, val := range input {
			numbers = append(numbers, findNum(val)...)
		}
	case map[string]interface{}:
		for _, val := range input {
			numbers = append(numbers, findNum(val)...)
		}
	case float64:
		numbers = append(numbers, int(input))
	}

	return numbers
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{}, 0)
	err1 := json.Unmarshal(input, &data)
	if err1 != nil {
		panic(err1)
	}

	sum := 0

	for _, num := range findNum(data) {
		sum += num
	}

	println(sum)
}
