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
	circuits := make(map[string][]*circuit)
	outputs := make(map[string][]int)

	for _, instruction := range instructions {
		circuit[input] := NewCircuit(instruction)
		for _, input := range circuit.inputWires {
			circuit[input] = append(circuits[input], circuit)
		}

		if circuit.hasoutput() {
			
		}
	}
}
