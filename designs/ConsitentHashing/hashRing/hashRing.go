package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"sync"
)

type HashRing struct {
	SortedHash   []uint32
	Hashmap      map[uint32]string
	ReplicaNodes int
	Mutex        sync.RWMutex
}

func NewHashRing(ReplicaNodes int) *HashRing {
	return &HashRing{
		SortedHash: []uint32{},
		Hashmap:    make(map[uint32]string),
		ReplicaNodes:    ReplicaNodes,
		Mutex:      sync.RWMutex{},
	}
}

func (hr *HashRing) AddNode(node string) {
	hr.Mutex.Lock()
	defer hr.Mutex.Unlock()

	for i := 0; i < hr.ReplicaNodes; i++ {
		ReplicaNodesNode := fmt.Sprintf("%s:%d", node, i)
		hash := crc32.ChecksumIEEE([]byte(ReplicaNodesNode))
		hr.Hashmap[hash] = node
		hr.SortedHash = append(hr.SortedHash, hash)
	}

	sort.Slice(hr.SortedHash, func(i, j int) bool {
		return hr.SortedHash[i] < hr.SortedHash[j]
	})
}

func (hr *HashRing) GetNode(key string) (string, bool) {
	hr.Mutex.RLock()
	defer hr.Mutex.RUnlock()

	hash := crc32.ChecksumIEEE([]byte(key))
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

func (hr *HashRing) RemoveNode(node string) {
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
