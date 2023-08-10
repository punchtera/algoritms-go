package linkedList

import (
	"fmt"
	"strconv"
	"testing"
)

func TestLinkedList(t *testing.T) {
	linkedList := LinkedList{
		head: LinkedListNode{
			value: 6,
		},
	}

	want := strconv.Itoa(6)
	result := strconv.Itoa(linkedList.head.value)
	if result != want {
		t.Errorf("The head value = %q is different than the want = %q \n", result, want)
	}
}

func TestLinkedList_GetLength(t *testing.T) {

	linkedList := LinkedList{
		head: LinkedListNode{
			value: 6,
		},
	}

	linkedList.head.AppendToTail(4)
	linkedList.head.AppendToTail(3)
	linkedList.head.AppendToTail(2)

	result := linkedList.GetLength()
	linkedList.ShowValues()

	want := 4
	if result != want {
		t.Errorf("The head value = %q is different than the want = %q \n", result, want)
	}
}

func TestLinkedList_GetLengthFull(t *testing.T) {

	secondNode := LinkedListNode{
		value: 4,
	}

	firstNode := LinkedListNode{
		value: 5,
	}

	headNode := LinkedListNode{
		value: 3,
	}

	firstNode.next = &secondNode
	headNode.next = &firstNode

	var n = &headNode
	i := 0
	for n.next != nil {
		i = +1
		fmt.Printf("Value -> %v\n", n.value)
		n = n.next
	}
	fmt.Printf("Value -> %v\n", n.value)
	fmt.Printf("Index : %v\n", i+1)
}
