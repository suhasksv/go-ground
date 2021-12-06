package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
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
    
    var horizontal int
    var depth int
    
    var h int
    var d int
    var aim int
    
    type ins struct {
        str string
        value int
    }
    
    var data []ins
    
    for _, s := range lines {
        split := strings.Split(s, " ")
        val, _ := strconv.Atoi(split[1])
        data = append(data, ins{split[0], val})
    }

    for _, i := range data {
        switch i.str {
            case "forward":
               horizontal += i.value
               h += i.value
               d += i.value * aim
            case "down":
                depth += i.value
                aim += i.value
            case "up":
                depth -= i.value
                aim -= i.value
        }
    }
    fmt.Println("Answer is:", horizontal * depth)
    fmt.Println("Answer is:", h * d)
}

