package main

import "fmt"

func main() {
    minNo := mini(10, 90, 9, 200, 6, 900)
    fmt.Println(minNo)
}

func mini(minf ...int) int {
    min := minf[0]

    for _, i := range minf {
        if i < min {
            min = i
        }
    }
    return min
}
