package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func BenchmarkManager_PutKey(b *testing.B) {
	m := NewManager()
	m.AddNode("192.168.0.1:8000")
	m.AddNode("192.168.1.101:9000")
	m.AddNode("192.168.101.5:3000")

	keyvals := make([]KeyVal, b.N)
	for i := 0; i < b.N; i++ {
		keyvals[i] = KeyVal{Key: fmt.Sprintf("Key:%d", i), Value: "value"}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.PutKey(keyvals[i])
	}
}

func BenchmarkManager_GetKey(b *testing.B) {
	m := NewManager()

	m.AddNode("192.168.0.1:8000")
	m.AddNode("192.168.1.101:9000")
	m.AddNode("192.168.101.5:3000")

	keyvals := make([]KeyVal, b.N)
	for i := 0; i < b.N; i++ {
		keyvals[i] = KeyVal{Key: fmt.Sprintf("Key:%d", i), Value: "value"}
	}

	for i := 0; i < b.N; i++ {
		m.PutKey(keyvals[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.GetKey(keyvals[i].Key)
	}
}

func BenchmarkRebalance(b *testing.B) {
	manager := NewManager()

	// Add nodes
	nodesToAdd := []string{"192.168.0.1:8000", "192.168.1.101:9000", "192.168.101.5:3000"}
	for _, node := range nodesToAdd {
		manager.AddNode(node)
	}

	// Generate and insert sample key-value pairs
	for i := 0; i < b.N; i++ {
		keyVal := KeyVal{Key: fmt.Sprintf("key%d", i), Value: "value"}
		manager.PutKey(keyVal)
	}

	b.ResetTimer()

	// Run the benchmark with concurrency
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			keyValsToRebalance, _ := manager.Nodes.Load(nodesToAdd[rand.Intn(len(nodesToAdd))])
			manager.KeysToRebalance <- keyValsToRebalance.([]KeyVal)
		}()
	}

	wg.Wait()
	manager.Close()
}
