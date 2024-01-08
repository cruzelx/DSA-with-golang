package main

import (
	"testing"

	lru "github.com/hashicorp/golang-lru/v2"
)

func TestLRU_GetPut(t *testing.T) {
	lru := NewLRU[int, interface{}](3)

	_, ok := lru.Get(1)

	if ok {
		t.Errorf("Error getting from cache: Expected %v, Got: %v", false, ok)
	}

	lru.Put(1, 1, -1)
	val, ok := lru.Get(1)

	if !ok {
		t.Errorf("Error getting from cache: Expected %d, Got: %d", 1, val)
	}

}

func TestLRU(t *testing.T) {
	lru := NewLRU[int, interface{}](3)

	lru.Put(1, 1, -1)
	lru.Put(2, 2, -1)
	lru.Put(3, 3, -1)

	// Sequence: 3->2->1
	validateLRU(t, lru, []int{3, 2, 1})

	lru.Get(2)

	// Sequence: 2->3->1
	validateLRU(t, lru, []int{2, 3, 1})

	lru.Put(4, 4, -1)

	//Sequence: 4->2->3
	validateLRU(t, lru, []int{4, 2, 3})

	lru.Get(3)

	//Sequence: 3->4->2
	validateLRU(t, lru, []int{3, 4, 2})

}

func validateLRU[K comparable, V any](t *testing.T, lru *LRU[K, V], expectedSequence []K) {
	curr := lru.Dll.Head

	for _, key := range expectedSequence {
		if curr == nil {
			t.Errorf("Expected node with key %v, but reached the end ", key)
			return
		}

		if curr.Key != key {
			t.Errorf("Expected the node with key %v, got a node with key %v", key, curr.Key)
			return
		}
		curr = curr.Next
	}

	if curr != nil {
		t.Errorf("Expected the end of DLL, but found more nodes")
	}
}

// func TestLRU_ConcurrentExpire(t *testing.T) {
// 	lru := NewLRU[int, interface{}](5)

// 	lru.Put(1, 1, -1)
// 	lru.Put(2, 2, time.Duration(5)*time.Second)
// 	lru.Put(3, 3, time.Duration(7)*time.Second)

// 	done := make(chan bool)
// 	defer close(done)

// 	go lru.RunActiveExpirationConcurrently(done)

// 	time.Sleep(time.Duration(6) * time.Second)

// 	_, ok := lru.Get(1)
// 	if !ok {
// 		t.Errorf("Expected a value for key:%d", 1)
// 	}

// 	_, ok = lru.Get(2)
// 	if ok {
// 		t.Errorf("Expected key %v to have expired already", 2)
// 	}

// 	_, ok = lru.Get(3)
// 	if !ok {
// 		t.Errorf("Expected a value for key:%d", 3)
// 	}
// }

// Benchmarks

// func BenchmarkCustomLRU_ConcurrentExpire(b *testing.B) {
// 	lru := NewLRU[int, interface{}](100)

// 	lru.Put(1, 1, -1)
// 	lru.Put(2, 2, time.Duration(5)*time.Second)
// 	lru.Put(3, 3, time.Duration(7)*time.Second)

// 	done := make(chan bool)
// 	defer close(done)

// 	b.ResetTimer()
// 	b.RunParallel(func(p *testing.PB) {
// 		for p.Next() {
// 			go lru.RunActiveExpirationConcurrently(done)
// 		}
// 	})

// }

func BenchmarkCustomLRU_Put(b *testing.B) {
	lru := NewLRU[int, interface{}](100)

	for i := 0; i < b.N; i++ {
		lru.Put(i, "Value", -1)
	}
}

func BenchmarkHashicorpLRU_Put(b *testing.B) {
	lru, _ := lru.New[int, interface{}](100)

	for i := 0; i < b.N; i++ {
		lru.Add(i, "Value")
	}
}

func BenchmarkCustomLRU_Get(b *testing.B) {
	lru := NewLRU[int, interface{}](100)

	for i := 0; i < b.N; i++ {
		lru.Put(i, "Value", -1)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lru.Get(i)
	}
}

func BenchmarkHashicorpLRU_Get(b *testing.B) {
	lru, _ := lru.New[int, interface{}](100)

	for i := 0; i < b.N; i++ {
		lru.Add(i, "Value")
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lru.Get(i)
	}

}
