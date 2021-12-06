/*
package main

import (
	"fmt"
)

func main() {

}
*/

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
    // fmt.Println(lines)

    if err := writeLines(lines, "foo.out.txt"); err != nil {
        log.Fatalf("writeLines: %s", err)
    }
    
   //   var content []int
   /* 
    for s := range lines {
        // string to int
        i, _ := strconv.Atoi(s)
        content = append(content, i)
    }
 */   
    var v1 int
    var prev = lines[0]
    for g := 0; g < len(lines); g++ {
        if lines[g] > prev {
            v1++
        }
        prev = lines[g]
    }
    println(v1)
    
    var v2 int = 0
    prev1 := lines[0]
    
    for f := 0; f < len(lines)-3; f++ {
        s := lines[f] + lines[f+1] + lines[f+2]
        
        if s > prev1 {
            v2++
        }
        prev1 = s
    }
    fmt.Println(v2)
}
