package main

import (
	"time"
)

type Node[K comparable, V any] struct {
	Key   K
	Value V

	// Time of creation or last access
	TimeStamp time.Time

	// Time-To-Live; keys expire if the current time > Timestamp + TTL
	TTL time.Duration

	Prev *Node[K, V]
	Next *Node[K, V]
}

type DoublyLinkedList[K comparable, V any] struct {
	Head *Node[K, V]
	Tail *Node[K, V]
}

func NewNode[K comparable, V any](key K, value V, ttl time.Duration) *Node[K, V] {
	return &Node[K, V]{
		Key:       key,
		Value:     value,
		TTL:       ttl,
		TimeStamp: time.Now(),
	}
}

func NewDLL[K comparable, V any]() *DoublyLinkedList[K, V] {
	return &DoublyLinkedList[K, V]{}
}

func (dll *DoublyLinkedList[K, V]) Prepend(node *Node[K, V]) {

	// If the DLL is empty
	if dll.Head == nil {
		dll.Head, dll.Tail = node, node
		return
	}

	node.Next = dll.Head
	node.Prev = nil
	dll.Head.Prev = node
	dll.Head = node

}

func (dll *DoublyLinkedList[K, V]) Remove(node *Node[K, V]) {

	if node == dll.Head {
		dll.Head = node.Next
	}
	if node == dll.Tail {
		dll.Tail = node.Prev
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = nil

}
