package doubleLinkedList

import (
	"fmt"
	"strconv"
)

type DoubleLinkedListNode struct {
	value int
	prev  *DoubleLinkedListNode
	next  *DoubleLinkedListNode
}

type DoubleLinkedList struct {
	head DoubleLinkedListNode
}

func (currentListNode *DoubleLinkedListNode) AppendToTail(nextValue int) {
	newListNode := DoubleLinkedListNode{value: nextValue}

	var n = currentListNode

	for n.next != nil {
		n = n.next
	}

	newListNode.prev = n
	n.next = &newListNode
	fmt.Printf("%v", strconv.Itoa(n.value))
	fmt.Printf("%v", n.next.value)
}

func (currentList *DoubleLinkedList) ShowValues() {
	var n = &currentList.head

	fmt.Println("\n Begin: ")

	for n.next != nil {
		fmt.Printf("%d", n.value)
		n = n.next
	}
	fmt.Println("\n End: ")
}
