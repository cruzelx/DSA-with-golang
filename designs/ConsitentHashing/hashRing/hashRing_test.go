package main

import (
	"testing"
)

func TestHashRing_AddGetRemoveNode(t *testing.T) {
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
	hr := NewHashRing(10)
	nodes := []string{"192.168.0.1:8080", "192.168.0.101:9000", "192.168.100.200:3000", "192.168.110.5:4000", "192.168.1.120:5000"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hr.AddNode(nodes[i%len(nodes)])
	}
}

func BenchmarkHashRing_GetNode(b *testing.B) {
	hr := NewHashRing(10)
	hr.AddNode("192.168.0.1:8080")
	hr.AddNode("192.168.0.101:9000")
	keys := []string{"key1", "key2", "key3", "key4", "key5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hr.GetNode(keys[i%len(keys)])
	}
}

func BenchmarkHashRing_RemoveNode(b *testing.B) {
	hr := NewHashRing(10)
	hr.AddNode("192.168.0.1:8080")
	hr.AddNode("192.168.0.101:9000")
	hr.AddNode("192.168.100.200:3000")
	nodes := []string{"192.168.0.1:8080", "192.168.0.101:9000", "192.168.100.200:3000"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hr.RemoveNode(nodes[i%len(nodes)])
	}
}
