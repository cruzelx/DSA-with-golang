package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	lru := NewLRU[int, interface{}](10)

	done := make(chan bool)
	go lru.RunActiveExpirationConcurrently(done)
	defer func() {
		done <- true
		close(done)
	}()

	var wg sync.WaitGroup

	for th := 0; th < 10; th++ {
		wg.Add(1)
		go func(threadId int) {
			defer wg.Done()

			for i := 0; i < 100; i++ {
				var t time.Duration
				n := rand.Intn(100)
				t = time.Duration(i%10) * time.Second

				time.Sleep(time.Duration(i%2) * time.Second)
				log.Printf("thread: %v, putting: %v", threadId, n)
				lru.Put(n, n, t)
				PrintDLL(lru.Dll)
			}
		}(th)
	}

	// Put values
	lru.Put(1, 1, -1)
	lru.Put(2, 2, time.Duration(0)*time.Second)
	lru.Put(3, 3, time.Duration(1)*time.Second)
	lru.Put(4, 4, time.Duration(2)*time.Second)
	lru.Put(5, 5, time.Duration(3)*time.Second)
	lru.Put(6, 6, time.Duration(4)*time.Second)

	time.Sleep(time.Second * 10)

	lru.Put(7, 7, time.Duration(3)*time.Second)
	time.Sleep(time.Second * 3)

	lru.Get(7)
	wg.Wait()
	PrintDLL(lru.Dll)
}
