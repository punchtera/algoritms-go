package linkedList

import (
	"strconv"
	"testing"
)

func TestLinkedList(t *testing.T) {
	linkedList := LinkedList{
		head: LinkedListNode{
			value: 6,
			next:  nil,
		},
	}

	want := strconv.Itoa(6)
	result := strconv.Itoa(linkedList.head.value)
	if result != want {
		t.Errorf("The head value is different value = %q, want = %q \n", result, want)
	}
}
