package doubleLinkedList

import (
	"strconv"
	"testing"
)

func TestDoubleLinkedList(t *testing.T) {
	linkedList := LinkedList{
		head: DoubleLinkedListNode{
			value: 6,
		},
	}

	want := strconv.Itoa(6)
	result := strconv.Itoa(linkedList.head.value)
	if result != want {
		t.Errorf("The head value = %q is different than the want = %q \n", result, want)
	}
}
