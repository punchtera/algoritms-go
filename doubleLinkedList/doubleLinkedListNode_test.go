package doubleLinkedList

import (
	"strconv"
	"testing"
)

func TestDoubleLinkedList(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{
		head: &DoubleLinkedListNode{
			value: 6,
		},
	}

	want := strconv.Itoa(6)
	result := strconv.Itoa(doubleLinkedList.head.value)
	if result != want {
		t.Errorf("The head value = %q is different than the want = %q \n", result, want)
	}
}

func TestDoubleLinkedListShowAll(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{
		head: &DoubleLinkedListNode{
			value: 6,
		},
	}
	doubleLinkedList.AppendToTail(2)
	doubleLinkedList.AppendToTail(11)

	doubleLinkedList.ShowValues()
}

func TestDoubleLinkedListShowReverseAll(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{
		head: &DoubleLinkedListNode{
			value: 6,
		},
	}
	doubleLinkedList.AppendToTail(2)
	doubleLinkedList.AppendToTail(11)

	doubleLinkedList.ShowReverseValues()
}

func TestDoubleLinkedListAppendHead(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{
		head: &DoubleLinkedListNode{
			value: 6,
		},
	}
	doubleLinkedList.AppendToHead(2)
	doubleLinkedList.AppendToTail(11)

	doubleLinkedList.ShowValues()
	// Should display [2] -> [6] -> 11
}
