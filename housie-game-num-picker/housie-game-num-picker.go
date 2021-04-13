package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

var (
	number []int
	choice int
	i      int
)

func main() {
	fmt.Println("\nWelcome to Quantum Tricks Housie Game Number Picker\nLet's get started...\n")
	var f string
	fmt.Println("Do you want to start the Game? :")
	f = strings.ToLower(f)
	if f == "no" || f == "n" {
		fmt.Println("Thanks for choosing Quantum Tricks Housie Game Number Picker")
		os.Exit(0)
	}

	for {
		choice = randInt(1, 100)

		if notIn(choice) == true {
			number = append(number, choice)
			fmt.Println("Picked #:", choice)
			fmt.Println(number)

			var ans string
			fmt.Print("Do you want to print next number? : ")
			fmt.Scan(&ans)
			ans = strings.ToLower(ans)
			if ans == "clear" || ans == "" {
				c := exec.Command("clear")
				c.Stdout = os.Stdout
				c.Run()
			} else if ans == "find" {
				find()
			}

			if len(number) == 99 {
				fmt.Println("\nAttention please\nAll numbers have been drawn\nGame Over\nThanks for using QT Housie Game Number picker..\n")
				break
			}
		}
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func notIn(a int) bool {
	for _, i = range number {
		if a != i {
			return true
			break
		}
		return false
	}
	return true
}

func find() {
	fmt.Println("\nWelcome to QT number finder")
	for true {
		var ask int
		fmt.Print("\nEnter the number that you want to find: ")
		fmt.Scan(&ask)
		for j := 0; j <= len(number)-1; j++ {
			if number[j] == ask {
				fmt.Printf("%d is has been drawn..\n", ask)
				break
			} else {
				fmt.Printf("%d is not yet drawn..\n Please wait patiently\n", ask)
			}
			break
		}
		var que string
		fmt.Print("Do you want to find another number? Enter 'y' to continue:")
		fmt.Scan(&que)
		strings.ToLower(que)

		if que == "y" {
			continue
		}
		break
	}
}
