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

func (doubleLinkedList *DoubleLinkedList) AppendToHead(nextValue int) {
	newListNode := DoubleLinkedListNode{value: nextValue}
	var headNode = doubleLinkedList.head

	newListNode.next = headNode
	newListNode.prev = doubleLinkedList.tail
	headNode.prev = &newListNode

	doubleLinkedList.head = &newListNode
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

func (currentList *DoubleLinkedList) ShowReverseValues() {
	var n = currentList.tail
	fmt.Println("\n Begin: ")

	for n.prev != nil {
		fmt.Printf("%d\n", n.value)
		n = n.prev
	}
	fmt.Printf("%d\n", n.value)
	fmt.Println("End: ")
}
