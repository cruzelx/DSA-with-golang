package main

import (
	"fmt"
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

	fmt.Println(hashRing)

}
