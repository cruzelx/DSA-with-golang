package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const (
	EXPIRED_PERCENTAGE_LIMIT = 25
	SAMPLING_FREQUENCY       = 10
	MAX_SAMPLING_SIZE        = 100
	MIN_SAMPLING_SIZE        = 10
)

// Data structure to represent LRU Cache
type LRU[K comparable, V any] struct {

	// Size limit of the cache
	Capacity int

	// Hashmap maps keys with their corresponding nodes in DLL
	Bucket map[K]*Node[K, V]

	// Doubly Linked List (DLL) to store key-value pairs
	// The key-value pairs are stored in the order of their frequency of use
	// From the most recently used to the least recently used
	// MRU->......->LRU
	Dll *DoublyLinkedList[K, V]

	// Sampling Size vaires based on percentage of expired keys in the sample
	SamplingSize int

	// Mutual exclusion lock to prevent multiple goroutines to access same resource at the same time (Race Condition)
	Mutex sync.Mutex

	// Manage concurrent expiration with wait groups
	Wg sync.WaitGroup
}

func NewLRU[K comparable, V any](capacity int) *LRU[K, V] {
	return &LRU[K, V]{
		Capacity:     capacity,
		Bucket:       make(map[K]*Node[K, V], capacity),
		Dll:          NewDLL[K, V](),
		SamplingSize: MIN_SAMPLING_SIZE,
		Mutex:        sync.Mutex{},
		Wg:           sync.WaitGroup{},
	}
}

// If cache-hit, node in the DLL is removed and moved to the front
func (lru *LRU[K, V]) Get(key K) (V, bool) {

	if node, ok := lru.Bucket[key]; ok {

		// Lazy delete: If found to be expired, remove and return "not found"
		if lru.isExpired(node) {
			lru.Mutex.Lock()

			delete(lru.Bucket, node.Key)
			lru.Dll.Remove(node)

			lru.Mutex.Unlock()
			return *new(V), false
		}

		node.TimeStamp = time.Now()
		lru.Mutex.Lock()

		lru.Dll.Remove(node)
		lru.Dll.Prepend(node)

		lru.Mutex.Unlock()
		return node.Value, true

	}
	return *new(V), false
}

// When putting the key-val, if matched in hash table, replace the value, remove the node and prepend to the list
// else check if the bucket has reached its limit and delete the least recently used node i.e the last node then
// a new node with key val is prepended in DLL and reference it in hash table
func (lru *LRU[K, V]) Put(key K, value V, ttl time.Duration) {
	lru.Mutex.Lock()
	defer lru.Mutex.Unlock()

	node, exists := lru.Bucket[key]
	// If the key already exists, update value and move the node to the front of DLL
	if exists {
		if lru.isExpired(node) {
			delete(lru.Bucket, node.Key)
			lru.Dll.Remove(node)

			// Add the new key-value node in the LRU
			newNode := NewNode(key, value, ttl)
			lru.Bucket[key] = newNode

			// Move the newly added node to the front
			lru.Dll.Prepend(newNode)
			return
		}

		node.Value = value
		node.TimeStamp = time.Now()
		lru.Dll.Remove(node)
		lru.Dll.Prepend(node)
	} else {
		// Check if the capcity has reached the capacity limit
		// If so, evict the least recently used node i.e. tail of the DLL
		if len(lru.Bucket) >= lru.Capacity {
			delete(lru.Bucket, lru.Dll.Tail.Key)
			lru.Dll.Remove(lru.Dll.Tail)
		}
		// Add the new key-value node in the LRU
		newNode := NewNode(key, value, ttl)
		lru.Bucket[key] = newNode

		// Move the newly added node to the front
		lru.Dll.Prepend(newNode)

	}

}

func (lru *LRU[K, V]) isExpired(node *Node[K, V]) bool {
	if node == nil || node.TTL < 0 {
		return false
	}

	// if time_stamp + TTL > current_time, key is expired
	return node.TimeStamp.Add(node.TTL).Before(time.Now())
}

// Concurrently expire keys SAMPLING_FREQUENCY times per second
func (lru *LRU[K, V]) RunActiveExpirationConcurrently(done chan bool) {
	log.Println("active expiration has started concurrently")

	// Sends tick continuously as per SAMPLING_FREQUENCY
	ticker := time.NewTicker(time.Second / SAMPLING_FREQUENCY)
	defer ticker.Stop()

	// Expire keys concurrently on every tick
	for {
		select {
		case <-done:
			log.Println("Stopping keys expiration....")
			lru.Wg.Wait()
			return
		case <-ticker.C:
			lru.Wg.Add(1)
			go lru.expireKeys()
		}
	}
}

// Expires keys based on random sampling of keys with TTL
func (lru *LRU[K, V]) expireKeys() {
	defer lru.Wg.Done()

	keys := lru.randomKeysWithTTLSampling()
	n := len(keys)

	if n == 0 {
		return
	}

	// Any key expired is removed from the cache and expired counter is increased
	expiredCount := 0
	for _, key := range keys {
		node := lru.Bucket[key]
		if lru.isExpired(node) {
			log.Printf("key: %v is expired...", key)

			lru.Mutex.Lock()

			delete(lru.Bucket, key)
			lru.Dll.Remove(node)

			lru.Mutex.Unlock()

			expiredCount++
		}
	}

	// If more than 25% of the sample has expired, increase the sampling size
	expiredPercentage := expiredCount * 100 / n

	if expiredPercentage > EXPIRED_PERCENTAGE_LIMIT {
		lru.SamplingSize = lru.SamplingSize * 2

		if lru.SamplingSize > MAX_SAMPLING_SIZE {
			lru.SamplingSize = MAX_SAMPLING_SIZE
		}
	}

}

// Gets the keys with valid TTL
// Requires lock to prevent race
func (lru *LRU[K, V]) randomKeysWithTTLSampling() []K {

	// lru.Mutex.Lock()
	// defer lru.Mutex.Unlock()

	counter := Min(len(lru.Bucket), lru.SamplingSize)

	keys := []K{}
	for key, node := range lru.Bucket {

		if node.TTL < 0 {
			continue
		}

		keys = append(keys, key)

		counter--
		if counter == 0 {
			break
		}
	}

	// Might not need random shuffle as looping through a map gives key,val in random order
	// rand.Shuffle(len(keys), func(i, j int) {
	// 	keys[i], keys[j] = keys[j], keys[i]
	// })

	return keys

}

func PrintDLL[K comparable, V any](dll *DoublyLinkedList[K, V]) {
	curr := dll.Head

	for curr != nil {
		fmt.Print(curr.Key, " <=> ")
		curr = curr.Next
	}

	fmt.Println()
	fmt.Println()

}
func (lru *LRU[K, V]) Print() {
	fmt.Println(lru.Bucket)
	PrintDLL(lru.Dll)
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
