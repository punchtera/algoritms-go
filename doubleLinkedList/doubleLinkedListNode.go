package doubleLinkedList

import (
	"fmt"
)

type DoubleLinkedListNode struct {
	value int
	prev  *DoubleLinkedListNode
	next  *DoubleLinkedListNode
}

type DoubleLinkedList struct {
	head *DoubleLinkedListNode
	tail *DoubleLinkedListNode
}

func (doubleLinkedList *DoubleLinkedList) AppendToTail(nextValue int) {
	newListNode := DoubleLinkedListNode{value: nextValue}

	var n = doubleLinkedList.head

	for n.next != nil {
		n = n.next
	}

	newListNode.prev = n
	n.next = &newListNode
	doubleLinkedList.tail = &newListNode
}

func (currentList *DoubleLinkedList) ShowValues() {
	var n = currentList.head

	fmt.Println("\n Begin: ")

	for n.next != nil {
		fmt.Printf("%d\n", n.value)
		n = n.next
	}
	fmt.Printf("%d\n", n.value)
	fmt.Println("End: ")
}
