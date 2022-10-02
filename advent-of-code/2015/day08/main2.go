package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	codelen := 0
	codeChar := 0

	for sc.Scan() {
		line := sc.Text()

		codelen += len(line)
		codeChar += len(strconv.Quote(line))
	}
	println(codeChar - codelen)
}
