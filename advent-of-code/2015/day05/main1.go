package main

import (
	"io/ioutil"
	"regexp"
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
		if goodstring(str) {
			cnt++
		}
	}
	println(cnt)
}

func badstrchar(str string) bool {
	return !regexp.MustCompile("ab|cd|pq|xy").MatchString(str)
}

func mustContainVowels(str string) bool {
	vowels := []string{"a", "e", "i", "o", "u"}
	cnt := 0
	for _, i := range vowels {
		cnt += strings.Count(str, i)
	}
	if cnt >= 3 {
		return true
	}
	return false
}

func repeatLetters(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			return true
		}
	}
	return false
}

func goodstring(str string) bool {
	if mustContainVowels(str) && repeatLetters(str) && badstrchar(str){
		return true
	}
	return false
}