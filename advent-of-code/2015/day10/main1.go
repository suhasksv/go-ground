/*
package main

import (
	"io/ioutil"
	"strconv"
)

func main() {
	lines, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 40; i++ {
		var num []byte

		for j := 0; j < len(lines); j++ {
			nums := lines[j]

			s := j
			for j < len(lines)-1 && lines[j+1] == num {
				j++
			}
			end := j

			length := end - s + 1

			num = append(num, strconv.Itoa(length)[0])
			num = append(num, nums[j])
		}
		lines = num
	}
	println(len(lines))
}
*/

package main

import (
	"io/ioutil"
	"strconv"
	"fmt"
)

func main() {
	numbers, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	for i := 0; i < 40; i++ {
		var newNumbers []byte

		for j := 0; j < len(numbers); j++ {
			num := numbers[j]

			start := j
			for j < len(numbers)-1 && numbers[j+1] == num {
				j++
			}
			end := j

			length := end - start + 1

			newNumbers = append(newNumbers, strconv.Itoa(length)[0])
			newNumbers = append(newNumbers, numbers[j])
		}

		numbers = newNumbers
	}

	fmt.Println(len(numbers))
}
