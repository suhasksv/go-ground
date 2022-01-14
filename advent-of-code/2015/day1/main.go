package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//  "strconv"
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

	if err := writeLines(lines, "foo.out.txt"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
	fmt.Println(lines)
	var floor int
	//	var data []string
	//	for _, j := range lines {
	//	data = append(data, strings.SplitAfter(j, ""))
	//	}

	for i, _ := range lines {
		if lines[i] == "(" {
			floor++
		} else if lines[i] == ")" {
			floor--
		}
	}
	fmt.Println(floor)

}
