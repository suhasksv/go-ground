/*
package main

import (
	"fmt"
	"os"
)

func main() {
	// feb := "february"
	_31months := []string{"january", "march", "may", "july", "august", "october", "december"}
	_30months := []string{"april", "june", "september", "november"}

	var date, year int
	var month string
	fmt.Print("Enter date DD MM YYYY: ")
	fmt.Scan(&date, &month, &year)
	if date > 31 {
		fmt.Println("Please enter correct date 1 - 31")
		os.Exit(1)
	}
	for _, i := range _31months {
		if month == i {
			days := 31 - date
			if days < 0 {
				fmt.Println("Please enter correct date respective of the month entered")
				os.Exit(1)
			} else {
				fmt.Println("Days left in ", month, "are:", days)
			}
		} else if true == true {
			for _, j := range _30months {
				if month == j {
					days := 30 - date
					if days < 0 {
						fmt.Println("Plz enter correct date respective of the month entered")
						os.Exit(1)
					} else {
						fmt.Println("Days left in ", month, "are: ", days)
						os.Exit(0)
					}
				}
			}
		} else {
			if leapYear(year) {
				days := 29 - date
				if days < 0 {
					fmt.Println("Please enter correct date respective of the month entered")
					os.Exit(1)
				} else {
					fmt.Println("Days left in", month, "are: ", days)
					//	os.Exit(0)
				}
			} else {
				days := 28 - date
				if days < 0 {
					fmt.Println("Please enter correct date with respective of the month entered")
					os.Exit(1)
				} else {
					fmt.Println("Days left in", month, "are: ", days)
				}
			}
		}
	}
}

func leapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	} else {
		return false
	}
}
*/
package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	le       int
	headNode *Node
}

func lenNode(node *Node) int {
	if node == nil {
		return 0
	} else {
		return 1 + lenNode(node.next)
	}
}

func addElement(element int) (list1 *LinkedList) {
	list1 = &list
	newNode := Node{
		data: element,
	}

	if list.le == 0 {
		list.headNode = &newNode
		list.le++
		return
	} else {
		currentNode := list.headNode
		for position := 0; position < list.le; position++ {
			if currentNode.next == nil {
				currentNode.next = &newNode
				list.le++ // increase the le
				return
			}
			currentNode = currentNode.next
		}
	}
	return
}

var list LinkedList = LinkedList{le: 0}

func main() {
	//list := LinkedList{le: 0}
	fmt.Println(list.le)

	print("How many numbers do you want to enter? :")
	var ask int
	fmt.Scan(&ask)

	for i := 0; i <= ask; i++ {
		print("Enter a number: ")
		var ask1 int
		fmt.Scan(&ask1)
		local := addElement(ask1)
		fmt.Println(local.le)
	}

	fmt.Println(lenNode(list.headNode))
	fmt.Println(list.headNode.data)
	fmt.Println(list.headNode.next.data)
}
