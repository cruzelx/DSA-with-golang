package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"sync"
)

// Data structure that represents Consistent Hashring
// The hash ring is a sorted list (Ascending) of nodes, where each node is associated with a hash value.
type HashRing struct {

	// Sorted list of hash values of virtual/replica nodes.
	SortedHash []uint32

	// Map of hash values of virtual/replica nodesto original nodes.
	// Eg: replica:a:1 & replica:a:2=> original node:a
	Hashmap map[uint32]string

	// Number of replica/virtual nodes for each original node in the hash ring.
	ReplicaNodes int

	// Mutual Exclusion Lock for the hash ring.
	Mutex sync.RWMutex
}

func NewHashRing(ReplicaNodes int) *HashRing {
	return &HashRing{
		SortedHash:   []uint32{},
		Hashmap:      make(map[uint32]string),
		ReplicaNodes: ReplicaNodes,
		Mutex:        sync.RWMutex{},
	}
}

// Add a node to the hash ring.
// The node is associated with a hash value, which is calculated using the CRC32 algorithm.
// The hash value is used to determine the position of the node in the hash ring.
func (hr *HashRing) AddNode(node string) {

	// Prevent concurrent writes to the hash ring.
	hr.Mutex.Lock()
	defer hr.Mutex.Unlock()

	for i := 0; i < hr.ReplicaNodes; i++ {

		// Create number of replica nodes for each node in the hash ring.
		ReplicaNodes := fmt.Sprintf("%s:%d", node, i)
		hash := crc32.ChecksumIEEE([]byte(ReplicaNodes))
		hr.Hashmap[hash] = node
		hr.SortedHash = append(hr.SortedHash, hash)
	}

	// Sort the hash ring in ascending order.
	sort.Slice(hr.SortedHash, func(i, j int) bool {
		return hr.SortedHash[i] < hr.SortedHash[j]
	})
}

// Get the node associated with the given key.
func (hr *HashRing) GetNode(key string) (string, bool) {

	// Prevent concurrent reads to the hash ring.
	hr.Mutex.RLock()
	defer hr.Mutex.RUnlock()

	hash := crc32.ChecksumIEEE([]byte(key))

	// Find the first index in the sorted hash ring that is greater than or equal to the hash value.
	index := sort.Search(len(hr.SortedHash), func(i int) bool { return hr.SortedHash[i] >= hash })
	if index == len(hr.SortedHash) {
		index = 0
	}

	if len(hr.SortedHash) == 0 {
		return "", false
	}

	if val, ok := hr.Hashmap[hr.SortedHash[index]]; ok {
		return val, true
	} else {
		return "", false
	}
}

// Remove a node from the hash ring.
// The node is removed from the hash ring by removing all replica nodes associated with the node.
func (hr *HashRing) RemoveNode(node string) {

	// Prevent concurrent writes to the hash ring
	hr.Mutex.Lock()
	defer hr.Mutex.Unlock()

	newSortedHash := []uint32{}
	for _, hash := range hr.SortedHash {
		if hr.Hashmap[hash] != node {
			newSortedHash = append(newSortedHash, hash)
		} else {
			delete(hr.Hashmap, hash)
		}
	}
	hr.SortedHash = newSortedHash
}
