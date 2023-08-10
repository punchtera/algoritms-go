package linkedList

import (
	"fmt"
	"strconv"
)

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

type LinkedList struct {
	head LinkedListNode
}

func (currentListNode *LinkedListNode) AppendToTail(nextValue int) {
	newListNode := LinkedListNode{value: nextValue}

	var n = currentListNode

	for n.next != nil {
		n = n.next
	}
	n.next = &newListNode
	fmt.Printf("%v", strconv.Itoa(n.value))
	fmt.Printf("%v", n.next.value)
}

func (currentList *LinkedList) ShowValues() {
	var n = &currentList.head

	fmt.Println("\n Begin: ")
	for n.next != nil {
		fmt.Printf("%d", n.value)
	}
	fmt.Println("\n End: ")
}

func (currentList *LinkedList) GetLength() int {
	var n = &currentList.head
	var length = 1

	for n.next != nil {
		length += 1
		n = n.next
	}

	return length
}
