package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		return
	}

	split := strings.Split(string(input), "\n")
	cnt := 0

	for _, str := range split {
		if goodstr(str) {
			cnt++
		}
	}
	println(cnt)
}

func repeatbtw(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if str[i] == str[i+2] {
			return true
		}
	}
	return false
}

func _2pairs(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if strings.Count(str, str[i:i+2]) >= 2 {
			return true
		}
	}
	return false
}

func goodstr(str string) bool {
	if repeatbtw(str) && _2pairs(str) {
		return true
	}
	return false
}