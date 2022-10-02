package main

import (
	"io/ioutil"
	"strings"
)

func validPassword(password string) bool {
	return incStr(password) && twoPair(password) && excludedChar(password)
}

func incStr(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i]+1 == password[i+1] && password[i]+2 == password[i+2] {
			return true
		}
	}
	return false
}

func excludedChar(password string) bool {
	return !strings.ContainsAny(password, "iol")
}

func twoPair(password string) bool {
	pairs := 0
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			pairs++
			i++
		}
	}
	return pairs >= 2
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}
	password := string(input)
	nextPass := 1

	for nextPass != 0 {
		incrementPassword(&password)
		if validPassword(password) {
			nextPass--
		}
	}
	println(password)
}

func incrementPassword(password *string) {
	newPassword := []byte(*password)
	for i := len(newPassword) - 1; i >= 0; i-- {
		if newPassword[i] != 'z' {
			newPassword[i]++
			break
		} else {
			newPassword[i] = 'a'
		}
	}
	*password = string(newPassword)
}
