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

func printList() {
	currNode := list.headNode
	for i := 1; i <= lenNode(list.headNode); i++ {
		if i == lenNode(list.headNode) {
			print(currNode.data)
		} else {
			print(currNode.data, " -> ")
		}
		currNode = currNode.next
	}
	println()
}

func DeleteElement(element int) (list1 *LinkedList) {
	currentPosition := list.headNode
	if list.le == 0 {
		return
	}
	for index := 0; index < list.le; index++ {
		if currentPosition.data == element {
			currentPosition = nil // set the position to nil
			for currentPosition.next != nil {
				currentPosition = currentPosition.next
			}
		}
	}
	list.le -= 1
	return
}

func replaceElement(element int, replacer int) {
	var currentPosition = list.headNode
	for index := 0; index < list.le; index++ {
		if currentPosition.data == element {
			currentPosition.data = replacer
		}
		currentPosition = currentPosition.next
	}
}

var list = LinkedList{le: 0}

func main() {
	//list := LinkedList{le: 0}
	fmt.Println(list.le)

	print("How many numbers do you want to enter? :")
	var ask0 int
	_, err := fmt.Scan(&ask0)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= ask0; i++ {
		print("Enter a number: ")
		var ask1 int
		fmt.Scan(&ask1)
		addElement(ask1)
		//local := addElement(ask1)
		//fmt.Println(local.le)
	}

	fmt.Println(lenNode(list.headNode))

	printList()

	var ask1 int
	print("Which element do you want to delete? :")
	fmt.Scan(&ask1)

	DeleteElement(ask1)

	printList()

	print("Which element you want to replace and with which element do you want it to be replaced? Enter the values with a space a <space> b:")
	fmt.Scan(&ask0, &ask1)
	println(ask0)
	println(ask1)
	replaceElement(ask0, ask1)

	printList()
	//fmt.Println(list.headNode.data)
	//fmt.Println(list.headNode.next.data)
}
