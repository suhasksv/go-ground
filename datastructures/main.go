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

	for i := 1; i <= ask; i++ {
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
