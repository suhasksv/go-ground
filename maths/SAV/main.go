package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const pi = 3.14

func main() {
	fmt.Println("Welcome to Surface Area Volume Calculator CLI")
	fmt.Println("What do u want to calculate?\nSupported -\nCube\nCuboid\nCylinder\nCone\nSphere\nHemisphere")
	var ask1 string
	fmt.Scan(&ask1)

	switch strings.ToLower(ask1) {
	case "cube":
		fmt.Println("What do you want to find?\nSupported - \nTSA\nCSA\nVolume")
		var ask string
		fmt.Scan(&ask)

		fmt.Print("Enter the value of Side:")
		var a int
		fmt.Scan(&a)
		cube(ask, a)

	case "cuboid":
		fmt.Println("What do you want to find?\nSupported - \nTSA\nCSA\nVolume")
		var ask string
		fmt.Scan(&ask)

		fmt.Print("Enter the vales of l, b, h with a space:")
		var l, b, h int
		fmt.Scan(&l, &b, &h)
		fmt.Println(cuboid(ask, l, b, h))

	case "cylinder":
		fmt.Println("What do you want to find?\nSupported - \nTSA\nCSA\nVolume")
		var ask string
		fmt.Scan(&ask)

		fmt.Print("Enter the values of r and h with a space:")
		var r, h float64
		fmt.Scan(&r, &h)
		fmt.Println(cylinder(ask, r, h))

	case "cone":
		fmt.Println("What do you want to find?\nSupported - \nTSA\nCSA\nVolume")
		var ask string
		fmt.Scan(&ask)

		fmt.Print("Enter the values of l, r wait a space:")
		var l, r float64
		fmt.Println(cone(ask, l, r))

	case "sphere":
		fmt.Println("What do you want to find?\nSupported - \nTSA\nCSA\nVolume")
		var ask string
		fmt.Scan(&ask)

		fmt.Print("Enter the value of r:")
		var r float64
		fmt.Scan(&r)
		fmt.Println(sphere(ask, r))

	case "hemisphere":
		fmt.Println("What do you want to find?\nSupported - \nTSA\nCSA\nVolume")
		var ask string
		fmt.Scan(&ask)

		fmt.Print("Enter the value of r:")
		var r float64
		fmt.Scan(&r)
		fmt.Println(hemisphere(ask, r))
	}

	callAgain()
}

func cube(ask string, a float64) float64 {
	switch ask {
	case "TSA":
		return 6 * a * a
	case "CSA":
		return 4 * a * a
	case "VOL":
		return a * a * a
	}
	return 0
}

func cuboid(ask string, l, b, h float64) float64 {
	switch ask {
	case "TSA":
		return 2 * ((l * b) + (b * h) + (l * h))
	case "CSA":
		return 2 * ((b * h) + (l * h))
	case "VOL":
		return l * b * h
	}
	return 0
}

func cylinder(ask string, r, h float64) float64 {
	switch ask {
	case "TSA":
		return 2*pi*r*r + 2*pi*r*h
	case "CSA":
		return 2 * pi * r * h
	case "VOL":
		return pi * r * r * h
	}
	return 0
}

func cone(ask string, l, r float64) float64 {
	switch ask {
	case "TSA":
		return pi * r * l + pi * r * r
	case "CSA":
		return pi * r * l
	case "VOL":
		h := math.Sqrt((l * l) + (r * r))
		return 0.33 * pi * r * r * h
	}
	return 0
}

func sphere(ask string, r float64) float64 {
	switch ask {
	case "TSA":
		return 4 * pi * r * r
	case "CSA":
		return 4 * pi * r * r
	case "VOL":
		return 1.33 * pi * r * r * r
	}
	return 0
}

func hemisphere(ask string, r float64) float64 {
	switch ask {
	case "TSA":
		return 3 * pi * r * r
	case "CSA":
		return 2 * pi * r * r
	case "VOL":
		return 0.66 * pi * r * r
	}
	return 0
}

func callAgain() {
	fmt.Print("Do you want to calculate another one? type 'y' to continue: ")
	var ask string
	fmt.Scan(&ask)
	if ask == "y" {
		main()
	} else {
		os.Exit(0)
	}
}
