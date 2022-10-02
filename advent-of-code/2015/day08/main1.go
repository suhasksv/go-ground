package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	line := strings.Split(string(input), "\n")

	codeChar := 0
	codelen := 0

	for _, lines := range line {
		codelen += len(lines)
		s, _ := strconv.Unquote(lines)
		codeChar += len(s)
	}
	println(codelen - codeChar)
}
