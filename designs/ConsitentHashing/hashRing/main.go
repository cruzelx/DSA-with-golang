package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	hashRing := NewHashRing(3)

	hashRing.AddNode("192.168.0.1:8080")
	hashRing.AddNode("192.168.0.2:8080")
	hashRing.AddNode("192.168.0.3:8080")
	hashRing.AddNode("192.168.0.4:8080")
	hashRing.AddNode("192.168.0.5:8080")

	fmt.Println(hashRing.getNode("alex"))
	fmt.Println(hashRing.getNode("12345"))
	fmt.Println(hashRing.getNode("zebra"))
	fmt.Println(hashRing.getNode("birds"))

	hashRing.AddNode("192.168.0.6:8080")
	hashRing.AddNode("192.168.0.7:8080")
	hashRing.AddNode("192.168.0.8:8080")
	hashRing.AddNode("192.168.0.9:8080")

	fmt.Println(hashRing.getNode("gxjyaJP4so"))
	fmt.Println(hashRing.getNode("OpuIQFMcD6"))
	fmt.Println(hashRing.getNode("Jvn6P5RIhi"))

	// test(hashRing)
}

func test(hashRing *HashRing) {

	keys := []string{}

	for i := 0; i < 10; i++ {
		for j := 0; j < 256; j++ {
			hashRing.AddNode(fmt.Sprintf("192.168.%d.%d", i, j))
			keys = append(keys, genRandString(10))
		}
	}

	start := time.Now()

	for _, v := range keys {
		hashRing.getNode(v)
	}

	fmt.Printf("elapsed time: %v", time.Since(start))
	fmt.Println()
}

func genRandString(size int) string {
	rand.Seed(time.Now().UnixNano())

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, size)

	for i := 0; i < size; i++ {
		result[i] += charset[rand.Intn(len(charset))]
	}
	return string(result)
}
