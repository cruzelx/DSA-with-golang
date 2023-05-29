package main

import "fmt"

type KeyVal struct {
	key   string
	value interface{}
}

type HashTable struct {
	bucket_size int
	bucket      [][]KeyVal
}

func NewHashTable(bucket_size int) *HashTable {
	bucket := make([][]KeyVal, bucket_size)
	return &HashTable{
		bucket_size: bucket_size,
		bucket:      bucket,
	}
}

func RemoveAtIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func (ht *HashTable) _hash(key string) int {
	hash := 0

	for _, c := range []rune(key) {
		char_code := int(c)
		hash += char_code
	}
	return hash % ht.bucket_size
}

func (ht *HashTable) get(key string) interface{} {
	hash := ht._hash(key)

	if ht.bucket[hash] != nil {
		for _, v := range ht.bucket[hash] {
			if v.key == key {
				return v.value
			}
		}
	}
	return nil
}

func (ht *HashTable) set(key string, value interface{}) {
	hash := ht._hash(key)

	if ht.bucket[hash] != nil {
		for i, v := range ht.bucket[hash] {
			if v.key == key {
				ht.bucket[hash][i] = KeyVal{key, value}
			} else {
				ht.bucket[hash] = append(ht.bucket[hash], KeyVal{key, value})
			}
			break
		}
	} else {
		ht.bucket[hash] = append(ht.bucket[hash], KeyVal{key: key, value: value})
	}
}

func (ht *HashTable) remove(key string) bool {
	hash := ht._hash(key)

	if ht.bucket[hash] != nil {
		for i, v := range ht.bucket[hash] {
			if v.key == key {
				ht.bucket[hash] = RemoveAtIndex(ht.bucket[hash], i)
				return true
			}
		}
	}
	return false

	// if key_val != (KeyVal{}) && key_val.key == key {
	// 	ht.bucket[hash] = KeyVal{}
	// 	return true
	// }

	// return false
}

func (ht *HashTable) display() {
	for idx, val := range ht.bucket {
		fmt.Printf("%d : %v", idx, val)
		fmt.Println()
	}
}
