package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	// "flag"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func main() {
	lines, err := readLines("foo.in.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	//fmt.Println(lines)

	if err := writeLines(lines, "foo.out.txt"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}

	var data []string
	for _, j := range lines {
		data = append(data, strings.SplitAfter(j, "\n")...)
	}
	var sum int
	var totalRibbon int
	for _, s := range data {
		// paper
		split := strings.Split(s, "x")
		l, _ := strconv.Atoi(split[0])
		b, _ := strconv.Atoi(split[1])
		w, _ := strconv.Atoi(split[2])
		var data []int
		data = append(data, l*b, l*w, b*w)
		sum += (2*l*b) + (2*l*w) + (2*b*w) + minimum(data)
		data = nil

		// ribbon
		data = append(data, l+l+w+w, l+l+b+b, w+w+b+b)
		ribbon := minimum(data)
		bow := l * b * w
		totalRibbon += ribbon + bow
		data = nil
	}
	fmt.Println("Puzzle 1 answer:", sum)
	fmt.Println("Puzzle 2 answer:", totalRibbon)
}

func minimum(arr []int) int {
	min := arr[0]
	for i, j := range arr {
		if j < min {
			min = arr[i]
		}
		continue
	}
	return min
}
