package main

import "fmt"

type LRU struct {
	capacity int
	bucket   map[int]*Node
	dll      *DoublyLinkedList
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		bucket:   make(map[int]*Node, capacity),
		dll:      NewDLL(),
	}
}

// When getting the key, if matched in hashtable,
// the node in the hastable is removed its position in doubly linked list and prepend to it

func (lru *LRU) Get(key int) int {
	if node, ok := lru.bucket[key]; ok {
		lru.dll.Remove(node)
		lru.dll.Prepend(node)
		return node.Value

	}
	return -1
}

// When putting the key-val, if matched in hash table, replace the value, remove the node and prepend to the list
// else check if the bucket has reached its limit and delete the least recently used node i.e the last node then
// a new node with key val is prepended in doubly linked list and reference it in hash table
func (lru *LRU) Put(key int, value int) {
	if node, ok := lru.bucket[key]; ok {
		node.Value = value
		lru.dll.Remove(node)
		lru.dll.Prepend(node)
		return
	} else {
		if len(lru.bucket) >= lru.capacity {
			delete(lru.bucket, lru.dll.Tail.Key)
			lru.dll.RemoveTail()
		}
		newNode := NewNode(key, value)
		lru.bucket[key] = newNode
		lru.dll.Prepend(newNode)

	}

}

func PrintDLL(dll *DoublyLinkedList) {
	curr := dll.Head

	for curr != nil {
		fmt.Print(curr, " ")
		curr = curr.Next

	}
	fmt.Println()
	fmt.Println()

}
func (lru *LRU) Print() {
	fmt.Println(lru.bucket)
	PrintDLL(lru.dll)
}
