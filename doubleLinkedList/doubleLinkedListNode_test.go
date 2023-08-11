package doubleLinkedList

import (
	"strconv"
	"testing"
)

func TestDoubleLinkedList(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{
		head: DoubleLinkedListNode{
			value: 6,
		},
	}

	want := strconv.Itoa(6)
	result := strconv.Itoa(doubleLinkedList.head.value)
	if result != want {
		t.Errorf("The head value = %q is different than the want = %q \n", result, want)
	}
}

func TestDoubleLinkedListShow(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{
		head: DoubleLinkedListNode{
			value: 6,
		},
	}

	doubleLinkedList.ShowValues()

}
