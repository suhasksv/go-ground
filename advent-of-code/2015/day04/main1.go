package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"regexp"
)

func main() {
	input := "iwrupvqb"
	cnt := 0
	regp := regexp.MustCompile("\\A00000")
	mine := ""

	for !regp.MatchString(mine) {
		cnt++
		hashgen := md5.New()
		_, err := io.WriteString(hashgen, input)
		if err != nil {
			return
		}
		_, err = io.WriteString(hashgen, fmt.Sprintf("%d", cnt))
		if err != nil {
			return
		}
		mine = fmt.Sprintf("%x", hashgen.Sum(nil))
	}
	println(cnt)
}
