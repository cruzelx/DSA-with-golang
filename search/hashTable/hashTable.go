package main

import (
	"fmt"
	"hashTable/hasher"
)

// Data structure to store key-val pairs.
type KeyVal struct {
	//  Key is of type string but it can be of any valid types supported by golang.
	// To accomodate for different Key type include hash functions for each type with switch case in _hash() function.
	Key string

	// Store any type of value
	Value interface{}
}

type HashTable struct {
	// Represents the total capacity of the HashTable
	BucketSize int

	// Represents the current number of the slots filled.

	FilledSize int

	// Store the data in array to compensate for collision
	Bucket [][]KeyVal

	// If (FilledSize*100)/BucketSize >= load factor, double the bucket size and rehash
	LoadFactor int

	// Function to calculate hash
	HashFunc func(string) uint32
}

func NewHashTable(BucketSize int, LoadFactor int, HashFunc func(string) uint32) *HashTable {
	return &HashTable{
		BucketSize: BucketSize,
		Bucket:     make([][]KeyVal, BucketSize),
		FilledSize: 0,
		LoadFactor: LoadFactor,
		HashFunc:   HashFunc,
	}
}

func (ht *HashTable) _hash(Key string) int {
	// Double hashing for better distribution at the expense of speed
	// Since, it uses two hash functions, resizing is costly
	// Upon collision new hash is calculated using enhanced double hashing (hash1+i*hash2+(i*i*i-i)/6) MOD tableSize
	// https://en.wikipedia.org/wiki/Double_hashing

	h1 := int(ht.HashFunc(Key) % uint32(ht.BucketSize))
	h2 := int(hasher.UnknownHash(Key) % uint32(ht.BucketSize))

	i := 0
	for len(ht.Bucket[h1]) > 0 && ht.Bucket[h1][0].Key != Key {
		h1 = (h1 + i*h2 + (i*i*i-i)/6) % ht.BucketSize
		if i == ht.BucketSize {
			break
		}
		i++
	}
	return h1

	// hash := ht.HashFunc(Key)
	// return int(hash % uint32(ht.BucketSize))
}

func (ht *HashTable) Get(Key string) interface{} {
	hash := ht._hash(Key)

	if len(ht.Bucket[hash]) > 0 {
		for _, v := range ht.Bucket[hash] {
			if v.Key == Key {
				return v.Value
			}
		}
	}
	return nil
}

func (ht *HashTable) Set(Key string, Value interface{}) {

	// While setting, check if the hash table is about to be filled
	// Calculate the load and if the load is about to reach the %load factor increase the Bucket size and re-distribute the Key-val pairs
	load := ht.FilledSize * 100 / ht.BucketSize

	if load >= ht.LoadFactor {
		// fmt.Println("Rehashing.....")

		// If load reaches given load factor, double the bucket size
		ht.BucketSize = ht.BucketSize * 2
		temp := ht.Bucket
		ht.Bucket = make([][]KeyVal, ht.BucketSize)

		// Rehash every key in the bucket compensating new size
		for _, v := range temp {
			for _, w := range v {
				hash := ht._hash(w.Key)
				ht.Bucket[hash] = append(ht.Bucket[hash], KeyVal{Key: w.Key, Value: w.Value})
			}
		}
	}

	hash := ht._hash(Key)

	// If Key exists, just update the value
	if len(ht.Bucket[hash]) > 0 {
		for i, v := range ht.Bucket[hash] {
			if v.Key == Key {
				ht.Bucket[hash][i].Value = Value
				return
			}
		}

	}

	// If the slot is empty, append to it
	if len(ht.Bucket[hash]) == 0 {
		ht.FilledSize++
	}
	ht.Bucket[hash] = append(ht.Bucket[hash], KeyVal{Key: Key, Value: Value})

}

func (ht *HashTable) Remove(Key string) bool {
	hash := ht._hash(Key)

	if len(ht.Bucket[hash]) > 0 {
		for i, v := range ht.Bucket[hash] {
			if v.Key == Key {
				// Remove the match item from the slot
				ht.Bucket[hash] = append(ht.Bucket[hash][:i], ht.Bucket[hash][:i+1]...)

				// If removing item empties the slot decrease FilledSize
				if len(ht.Bucket[hash]) == 0 {
					ht.FilledSize--
				}
				return true
			}
		}
	}
	return false
}

func (ht *HashTable) Display() {
	for idx, val := range ht.Bucket {
		fmt.Printf("%d : %v", idx, val)
		fmt.Println()
	}
}
