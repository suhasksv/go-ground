package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [1000][1000]bool

	fromreg := regexp.MustCompile("\\d+,\\d+")
	toreg := regexp.MustCompile("\\d+,\\d+")

	for scanner.Scan() {
		instruction := scanner.Text()

		fromstr := fromreg.FindString(instruction)
		tostr := toreg.FindString(instruction)

		from := strings.Split(fromstr, ",")
		to := strings.Split(tostr, ",")

		fromX, _ := strconv.Atoi(from[0])
		fromY, _ := strconv.Atoi(from[1])

		toX, _ := strconv.Atoi(to[0])
		toY, _ := strconv.Atoi(to[1])

		for x := fromX; x <= toX; x++ {
			for y := fromY; y <= toY; y++ {
				if strings.Contains(instruction, "off") {
					grid[x][y] = false
				} else if strings.Contains(instruction, "on") {
					grid[x][y] = true
				} else {
					grid[x][y] = !grid[x][y]
				}
			}
		}
	}

	t := 0

	for _, row := range grid {
		for _, col := range row {
			if col {
				t++
			}
		}
	}
	println(t)
}
