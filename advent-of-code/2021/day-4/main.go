package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	//fmt.Println(lines)

	var numbers []int

	for _, i := range strings.Split(lines[0], ",") {
		j, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal("\n\n\n", err)
		}
		numbers = append(numbers, j)
	}
	//fmt.Println(numbers)

	var bingoBoard [][]int

	for i := 2; i < len(lines); i += 6 {
		var board []int
		for _, s := range strings.Split(strings.Join(lines[i:i+5], " "), " ") {

			if s == "" {
				continue
			}
			s, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			board = append(board, s)
		}
		bingoBoard = append(bingoBoard, board)

		if len(board) != 25 {
			log.Fatal(board)
		}
		bingoBoard = append(bingoBoard, board)
	}
	// fmt.Println(bingoBoard)

	for _, n := range numbers {
		for _, b := range bingoBoard {
			for i, v := range b {
				if v == n {
					b[i] = 0
					break
				}

			}
			fmt.Println()
			// TODO: check for win
			if chkWin(b) {
				fmt.Println(b)
			}
		}
	}
}

func chkWin(b []int) bool {
	win := true
	for j := 0; j < 25; j += 5 {
		for i := j; i < j+5; i++ {
			if b[i] != 0 {
				win = false
			}
		}
		if win {
			return true
		}
	}
	return false
}
