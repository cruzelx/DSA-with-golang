package main

import (
	"fmt"
	"testing"
)

func TestHashRing_AddNode(t *testing.T) {
	nodeReplicas := 3
	hr := NewHashRing(nodeReplicas)

	hr.AddNode("192.168.0.255:8080")
	hr.AddNode("192.168.0.100:8080")

	key := "Alex"

	node, _ := hr.GetNode(key)
	if node != "192.168.0.255:8080" {
		t.Errorf("Expected the key to be in %v node, Got %v", "192.168.0.255:8080", node)
	}

	// After new nodes are added, the node should be remapped to one of the new nodes.
	hr.AddNode("192.168.0.2:8080")
	hr.AddNode("192.168.0.4:8080")

	node, _ = hr.GetNode(key)
	if node != "192.168.0.2:8080" {
		t.Errorf("Expected the key to be in %v node, Got %v", "192.168.0.2:8080", node)
	}

	hr.RemoveNode("192.168.0.2:8080")
	node, _ = hr.GetNode(key)
	if node == "192.168.0.2:8080" {
		t.Errorf("Key %v mapped to a deleted node %v", key, "192.168.0.2:8080")
	}

}

func BenchmarkHashRing_AddNode(b *testing.B) {
	nodeReplicas := 3
	hr := NewHashRing(nodeReplicas)

	for i := 0; i < b.N; i++ {
		hr.AddNode(fmt.Sprintf("192.168.0.1:%d", i))
	}
}

func BenchmarkHashRing_GetNode(b *testing.B) {
	nodeReplicas := 3
	hr := NewHashRing(nodeReplicas)

	for i := 0; i < b.N; i++ {
		hr.AddNode(fmt.Sprintf("192.168.0.1:%d", i))
	}

	keys := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = fmt.Sprintf("Key%d", i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		hr.GetNode(keys[i])
	}
}
