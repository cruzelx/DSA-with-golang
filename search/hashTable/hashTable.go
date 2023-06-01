package main

import "fmt"

type KeyVal struct {
	key   string
	value interface{}
}

type HashTable struct {
	bucket_size int
	filled_size int
	bucket      [][]KeyVal
}

func NewHashTable(bucket_size int) *HashTable {
	bucket := make([][]KeyVal, bucket_size)
	return &HashTable{
		bucket_size: bucket_size,
		bucket:      bucket,
		filled_size: 0,
	}
}

func RemoveAtIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func (ht *HashTable) _hash(key string) int {
	// hash := 0
	// g := 31

	// for _, c := range []rune(key) {
	// 	hash = g*hash + int(c)
	// }

	hash := 5381

	for _, c := range []rune(key) {
		char_code := int(c)
		hash = ((hash << 5) + hash) + char_code
	}
	// mul := 1

	// for _, c := range []rune(key) {
	// 	char_code := int(c)
	// 	hash += char_code
	// }

	// for better distribution of keys
	// for i, c := range []rune(key) {
	// 	char_code := int(c)
	// 	if i%4 == 0 {
	// 		mul = 1
	// 	} else {
	// 		mul *= 256
	// 	}
	// 	hash += char_code * mul
	// }
	return hash % ht.bucket_size
}

func (ht *HashTable) get(key string) interface{} {
	hash := ht._hash(key)

	if len(ht.bucket[hash]) > 0 {
		for _, v := range ht.bucket[hash] {
			if v.key == key {
				return v.value
			}
		}
	}
	return nil
}

// while setting check if the hash table is about to be filled
// calculate the load factor and if the load factor is gets to about 80% increase the bucket size and re-distribute the key-val pairs
func (ht *HashTable) set(key string, value interface{}) {

	load_factor := int(ht.filled_size * 100.0 / ht.bucket_size)

	fmt.Println("load factor: ", load_factor)

	if load_factor >= 80 {
		new_bucket_size := ht.bucket_size * 2
		new_hash_table := make([][]KeyVal, new_bucket_size)

		for _, v := range ht.bucket {
			for _, w := range v {
				hash := ht._hash(w.key)
				new_hash_table[hash] = append(new_hash_table[hash], KeyVal{key: w.key, value: w.value})
			}
		}
		ht.bucket = new_hash_table
	}

	hash := ht._hash(key)
	fmt.Println("Hash: ", hash)

	if len(ht.bucket[hash]) > 0 {
		for i, v := range ht.bucket[hash] {
			if v.key == key {
				ht.bucket[hash][i].value = value
				return
			}
		}

	}
	if len(ht.bucket[hash]) == 0 {
		ht.filled_size++
	}
	ht.bucket[hash] = append(ht.bucket[hash], KeyVal{key: key, value: value})

}

func (ht *HashTable) remove(key string) bool {
	hash := ht._hash(key)

	if len(ht.bucket[hash]) > 0 {
		for i, v := range ht.bucket[hash] {
			if v.key == key {
				ht.bucket[hash] = RemoveAtIndex(ht.bucket[hash], i)
				if len(ht.bucket[hash]) == 0 {
					ht.filled_size--
				}
				return true
			}
		}
	}
	return false
}

func (ht *HashTable) display() {
	for idx, val := range ht.bucket {
		fmt.Printf("%d : %v", idx, val)
		fmt.Println()
	}
}
