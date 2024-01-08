package main

import (
	"testing"
)

func TestDLLPrepend(t *testing.T) {
	dll := NewDLL[int, interface{}]()

	node := NewNode[int, interface{}](1, 10, -1)
	dll.Prepend(node)

	if dll.Head != node {
		t.Error("Head should be equal to node after prepend")
	}
}

func TestDLLRemove(t *testing.T) {
	dll := NewDLL[int, interface{}]()

	node1 := NewNode[int, interface{}](1, 100, -1)
	node2 := NewNode[int, interface{}](2, 200, -1)

	dll.Prepend(node1)
	dll.Prepend(node2)

	dll.Remove(node1)

	if dll.Head != node2 {
		t.Error("Head should be node2 after removing node1")
	}

	if node1.Next != nil || node1.Prev != nil {
		t.Error("Node1 should not have any previous or next value after removal")
	}
}
