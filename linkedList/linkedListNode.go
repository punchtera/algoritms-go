package linkedList

import "fmt"

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

type LinkedList struct {
	head LinkedListNode
}

func (currentListNode LinkedListNode) AppendToTail(nextValue int) {
	newListNode := LinkedListNode{value: nextValue}

	var n = &currentListNode

	for n.next != nil {
		n = n.next
	}
	n.next = &newListNode
}

func (currentList LinkedList) ShowValues() {
	var n = &currentList.head

	fmt.Println("Begin: ")
	for n.next != nil {
		fmt.Printf("%d", n.value)
	}
	fmt.Println("End: ")
}
