package main

import "fmt"

type KeyVal struct {
	key   string
	value interface{}
}

type HashTable struct {
	bucket_size int
	bucket      []KeyVal
}

func NewHashTable() *HashTable {
	bucket_size := 20
	bucket := make([]KeyVal, bucket_size)
	return &HashTable{
		bucket_size: bucket_size,
		bucket:      bucket,
	}
}

func (ht *HashTable) _hash(key string) int {
	hash := 0

	for _, c := range []rune(key) {
		utf_16 := int(c)
		hash += utf_16
	}
	return hash % ht.bucket_size
}

func (ht *HashTable) get(key string) *KeyVal {
	hash := ht._hash(key)

	if ht.bucket[hash] != (KeyVal{}) {
		return &ht.bucket[hash]
	}
	return nil
}

func (ht *HashTable) set(key string, value interface{}) {
	hash := ht._hash(key)

	if ht.bucket[hash] == (KeyVal{}) {
		ht.bucket[hash] = KeyVal{key, value}
	} else {
		key_val := ht.bucket[hash]

		if key_val.key == key {
			ht.bucket[hash] = KeyVal{key, value}
		}
	}
}

func (ht *HashTable) remove(key string) bool {
	hash := ht._hash(key)
	key_val := ht.bucket[hash]

	if key_val != (KeyVal{}) && key_val.key == key {
		ht.bucket[hash] = KeyVal{}
		return true
	}

	return false
}

func (ht *HashTable) display() {
	for idx, val := range ht.bucket {
		fmt.Println("%d : %v", idx, val)
	}
}
