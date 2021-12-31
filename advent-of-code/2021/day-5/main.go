package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	// fmt.Println(lines)

	if err := writeLines(lines, "foo.out.txt"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}

	type P struct {
		x1, y1 int
		x2, y2 int
	}

	var entries = make([]P, len(lines))
	vhlines := []P{}

	maxX := 0
	maxY := 0
	for i, l := range lines {
		fmt.Sscanf(l, "%d,%d -> %d,%d",
			&entries[i].x1,
			&entries[i].y1,
			&entries[i].x2,
			&entries[i].y2)

		if entries[i].x1 > maxX {
			maxX = entries[i].x1
		}
		if entries[i].x2 > maxX {
			maxX = entries[i].x2
		}
		if entries[i].y1 > maxY {
			maxY = entries[i].y1
		}
		if entries[i].y2 > maxY {
			maxY = entries[i].y2
		}

		if entries[i].x1 == entries[i].x2 || entries[i].y1 == entries[i].y2 {
			vhlines = append(vhlines, entries[i])
		}
	}
	fmt.Println(vhlines)
	fmt.Println(maxX, maxY)

	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}

	for _, l := range entries {
		if l.x1 == l.x2 {
			y1 := l.y1
			y2 := l.y2
			if y2 < y1 {
				y1 = l.y1
				y2 = l.y2
			}
			for y := y1; y <= y2; y++ {
				grid[y][l.x1]++
			}
		} else if l.y1 == l.y2 {
			x1 := l.x1
			x2 := l.x2
			if x2 < x1 {
				x1 = l.x2
				x2 = l.x1
			}
			for x := x1; x <= x2; x++ {
				grid[l.y1][x]++
			}
		} else {
			dx := int(float64(l.x2-l.x1) / math.Abs(float64(l.x2-l.x1)))
			dy := int(float64(l.y2-l.y1) / math.Abs(float64(l.y2-l.y1)))

			y := l.y1
			x := l.x1

			for y != l.y2+dy {
				grid[y][x]++
				y += dy
				x += dx
			}
		}
	}
	count := 0
	for _, row := range grid {
		for _, v := range row {
			if v >= 2 {
				count++
			}
		}
	}
	fmt.Println(count)
}
