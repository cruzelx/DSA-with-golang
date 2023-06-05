package main

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type HashRing struct {
	nodes      []string
	sortedHash []uint32
	hashmap    map[uint32]string
	replica    int
}

func NewHashRing(replica int) *HashRing {
	return &HashRing{
		sortedHash: []uint32{},
		hashmap:    make(map[uint32]string),
		replica:    replica,
	}
}

func (hr *HashRing) AddNode(node string) {
	for i := 0; i < hr.replica; i++ {
		replicaKey := fmt.Sprintf("%s:%d", node, i)
		hash := crc32.ChecksumIEEE([]byte(replicaKey))
		hr.hashmap[hash] = node
		hr.sortedHash = append(hr.sortedHash, hash)
	}

	sort.Slice(hr.sortedHash, func(i, j int) bool {
		return hr.sortedHash[i] < hr.sortedHash[j]
	})
}

func (hr *HashRing) getNode(key string) string {
	hash := crc32.ChecksumIEEE([]byte(key))
	index := sort.Search(len(hr.sortedHash), func(i int) bool { return hr.sortedHash[i] >= hash })
	if index == len(hr.sortedHash) {
		index = 0
	}
	return hr.hashmap[hr.sortedHash[index]]
}


