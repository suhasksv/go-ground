package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"strconv"
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
type cod struct {
	x, y int
}

func main() {
	lines, err := readLines("foo.in.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	// fmt.Println(lines)

	if err := writeLines(lines, "foo.out.txt"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}

	homes := make(map[cod]bool)
	coord := cod{0, 0}

	homes[coord] = true

	var data []string
	for _, j := range lines {
		data = append(data, strings.Split(j, "")...)
	}
	// TODO: Part 1
	for _, i := range data {
		switch i {
		case ">":
			coord = cod{coord.x + 1, coord.y}
		case "<":
			coord = cod{coord.x - 1, coord.y}
		case "^":
			coord = cod{coord.x, coord.y - 1}
		case "v":
			coord = cod{coord.x, coord.y + 1}
		}
		homes[coord] = true
	}
	fmt.Println("Part 1:", len(homes))
}
