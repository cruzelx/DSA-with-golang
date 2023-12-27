package main

import (
	"fmt"
	"hashTable/hasher"
	"math/rand"
	"testing"
	"time"
)

func TestHashTableGetSet(t *testing.T) {
	ht := NewHashTable(1, 80, hasher.Djb2)

	key := "apple"
	value := 22

	ht.Set(key, value)

	val := ht.Get(key)

	if val != value {
		t.Errorf("Expected value for key %s: %d, Got: %v", key, value, val)
	}

}

func TestHashTableRemove(t *testing.T) {
	ht := NewHashTable(1, 80, hasher.Djb2)

	key := "apple"
	value := 22

	ht.Set(key, value)
	removed := ht.Remove(key)

	if !removed {
		t.Errorf("Expected the key %s to be removed but failed", key)
	}

	key = "does not exist"
	removed = ht.Remove(key)
	if removed {
		t.Errorf("Expected remove operation for key %s: %v, Got: %v", key, false, true)
	}

}

func BenchmarkHashTableSet(b *testing.B) {
	hashFuncs := map[string]func(string) uint32{
		"djb2":        hasher.Djb2,
		"murmurHash3": hasher.MurmurHash3,
		"naiveHash":   hasher.NaiveHash,
		"fnv1a":       hasher.Fnv1a,
	}
	bucketSizes := []int{100, 1000, 10000}
	keyLengths := []int{10, 50, 100}

	for name, function := range hashFuncs {
		for _, size := range bucketSizes {
			for _, length := range keyLengths {
				ht := NewHashTable(size, 80, function)
				keys := generateRandomStrings(b.N, length)

				b.Run(fmt.Sprintf("Hash Function_%s_Bucket Size_%d_Key Length_%d", name, size, length), func(t *testing.B) {
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						ht.Set(keys[i], "Value")
					}
				})
			}
		}
	}
}

func BenchmarkHashTableGet(b *testing.B) {
	hashFuncs := map[string]func(string) uint32{
		"djb2":        hasher.Djb2,
		"murmurHash3": hasher.MurmurHash3,
		"naiveHash":   hasher.NaiveHash,
		"fnv1a":       hasher.Fnv1a,
	}
	bucketSizes := []int{100, 1000, 10000}
	keyLengths := []int{10, 50, 100}

	for name, function := range hashFuncs {
		for _, size := range bucketSizes {
			for _, length := range keyLengths {
				ht := NewHashTable(size, 80, function)
				keys := generateRandomStrings(b.N, length)

				for i := 0; i < b.N; i++ {
					ht.Set(keys[i], "Value")
				}
				b.Run(fmt.Sprintf("Hash Function_%s_Bucket Size_%d_Key Length_%d", name, size, length), func(t *testing.B) {
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						ht.Get(keys[i])
					}

				})
			}
		}
	}
}

func generateRandomStrings(count, maxLength int) []string {
	rand.Seed(time.Now().UnixNano())

	var result []string
	charset := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := 0; i < count; i++ {
		length := rand.Intn(maxLength) + 1 // Generating variable length
		str := make([]byte, length)
		for j := 0; j < length; j++ {
			str[j] = charset[rand.Intn(len(charset))]
		}
		result = append(result, string(str))
	}

	return result
}
