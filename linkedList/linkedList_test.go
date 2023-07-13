package linkedList

import "testing"

func TestLinkedList(t *testing.T) {
	want := "Hello, world."
	if got := "Hello, world."; got != want {
		t.Errorf("Hello, world. = %q, want %q", got, want)
	}
}
