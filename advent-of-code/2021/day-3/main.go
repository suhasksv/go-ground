package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
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
    
    var bits = make([]int, len(lines[0]));
    for _, n := range lines {
        for i := range n {
            if n[i] == '1' {
                bits[i]++
            }
        }
    }
    fmt.Println(bits)
    l := len(lines)
    gamma := ""
    epsilon := ""
    for j := 0; j < len(lines[0]); j++ {
        if bits[j] > (l - bits[j]) {
            gamma += "1"
            epsilon += "0"
        } else if (l - bits[j]) > bits[j] {
            gamma += "0"
            epsilon += "1"
        }
    }
    fmt.Println("gamma:", gamma, "epsilon:", epsilon, "\nAnswer:", 0b100111000110*0b011000111001)
}
