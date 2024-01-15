package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	manager := NewManager()
	// defer manager.Close()

	var wg sync.WaitGroup

	wg.Wait()

	manager.AddNode("192.168.0.1:8081")
	manager.AddNode("192.168.0.99:9000")
	manager.AddNode("192.168.0.200:8000")
	manager.AddNode("192.168.0.100:8089")

	for i := 0; i < 10; i++ {
		manager.Wg.Add(1)
		go func(i int) {
			for j := 0; j < 100; j++ {
				manager.PutKey(KeyVal{fmt.Sprintf("key%v:%v", rand.Intn(i+j+1), rand.Intn(i*j+1)), "value"})
				// log.Printf("[%v] putting....%v\n", i, fmt.Sprintf("key%v:%v", i+j, i*j))
				time.Sleep(time.Millisecond * 10)
			}
			manager.Wg.Done()
		}(i)
	}

	// for i := 0; i < 10; i++ {
	// 	manager.Wg.Add(1)
	// 	go func(i int) {
	// 		for j := 0; j < 100; j++ {
	// 			manager.GetKey(fmt.Sprintf("key%v:%v", rand.Intn(i+j+1), rand.Intn(i*j+1)))
	// 			log.Printf("[%v] getting....%v\n", i, fmt.Sprintf("key%v:%v", i+j, i*j))
	// 			// time.Sleep(time.Millisecond)
	// 		}
	// 		manager.Wg.Done()
	// 	}(i)
	// }

	time.Sleep(time.Millisecond * 500)

	go func() {

		manager.AddNode("192.168.101.205:3000")
		// time.Sleep(time.Millisecond * 50)
	}()

	go func() {

		manager.RemoveNode("192.168.0.1:8081")
		time.Sleep(time.Millisecond * 50)

		manager.RemoveNode("192.168.0.99:9000")
	}()

	// manager.Wg.Wait()

	time.Sleep(time.Second * 1)

	manager.Close()
	PrintMapAsTable(&manager.Nodes)
	visualizeHashRing(manager.HashRing)

}
