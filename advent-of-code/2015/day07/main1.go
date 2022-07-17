package main

import (
	"io/ioutil"
	"strings"
)

type circuit struct {
	outwire string
	inwire  []string
	values  []int
	gate    string
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	instructions := strings.Split(string(input), "\n")

}
