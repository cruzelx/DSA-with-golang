package main

import (
	"log"
	"sync"
)

const (

	// Number of virtual/replica nodes for each original node
	NumReplicaNodes = 5

	// Maximum size of key-value pairs passed to rebalance routine
	// Batching allows for efficient rebalancing
	RebalanceBatchSize = 10
)

// Data structure to hold key-value pairs
type KeyVal struct {
	Key   string
	Value string
}

// Manager is a data structure that manages a hash ring and distributes key-value pairs in the nodes of the hash ring.
type Manager struct {

	// Map keys are nodes, and values are slices of key-value pairs
	// sync.Map for locking during concurrent access
	Nodes sync.Map

	// HashRing is a hash ring that maps keys to nodes
	HashRing *HashRing

	// Mutex is a mutex for locking during concurrent access
	Mutex sync.RWMutex

	// KeysToRebalance is a channel that receives slices of key-value pairs to rebalance
	KeysToRebalance chan []KeyVal

	// Done is a channel that is closed when the manager is closed
	Done chan struct{}

	// Manage go routines throughout the app
	Wg sync.WaitGroup
}

func NewManager() *Manager {
	m := &Manager{
		HashRing:        NewHashRing(NumReplicaNodes),
		Mutex:           sync.RWMutex{},
		KeysToRebalance: make(chan []KeyVal),
		Done:            make(chan struct{}),
		Wg:              sync.WaitGroup{},
	}

	// Start a go routine to rebalance keys in the hash ring when nodes are added/removed to/from the manager
	go m.rebalance()
	return m

}

// Uses channel to receive slice of key-val pairs from AddNode and RemoveNode method
func (m *Manager) rebalance() {

	for {
		select {

		// Receive slices of key-value pairs to rebalance
		case keyValsToRebalance := <-m.KeysToRebalance:
			m.Wg.Add(1)
			go m.rebalanceKeys(keyValsToRebalance)

		// Receive a signal to close the manager
		case <-m.Done:

			// Close the channel to receive slices of key-value pairs to rebalance
			close(m.KeysToRebalance)
			return

		}
	}

}

// Rebalances the keys in the hash ring by distributing them evenly among the nodes.
func (m *Manager) rebalanceKeys(keyValsToRebalance []KeyVal) {

	// Remove go routine from the wait group when rebalancing is complete
	defer m.Wg.Done()

	for _, kv := range keyValsToRebalance {
		if node, ok := m.HashRing.GetNode(kv.Key); ok {

			// Load the slice of key-value pairs from the nodes
			// If the node is not found, create a new slice
			if nodes, found := m.Nodes.LoadOrStore(node, []KeyVal{kv}); found {

				// Lock the mutex during concurrent update to the key-value pairs in nodes
				m.Mutex.Lock()

				// Update the slice of key-value pairs in nodes
				keyVals := nodes.([]KeyVal)
				keyVals = append(keyVals, kv)
				m.Mutex.Unlock()

				// Store the updated slice of key-value pairs in nodes
				m.Nodes.Store(node, keyVals)
			}

		}
	}
}

// Exits the manager when all the go routines are done rebalancing keys in the hash ring
func (m *Manager) Close() {
	m.Done <- struct{}{}
	m.Wg.Wait()

	close(m.Done)
}

// Adds a key-value pair to the nodes in manager through hash ring.
func (m *Manager) PutKey(keyVal KeyVal) {

	// Get the corresponding node in the hash ring for the key-value pair
	if node, ok := m.HashRing.GetNode(keyVal.Key); ok {

		// Load the slice of key-value pairs from the nodes
		// If the node is not found, create a new slice
		if nodes, loaded := m.Nodes.LoadOrStore(node, []KeyVal{keyVal}); loaded {

			// Lock the mutex during concurrent access to the key-value pairs in nodes
			m.Mutex.Lock()

			// Update the slice of key-value pairs in nodes
			keyvals := nodes.([]KeyVal)
			keyvals = append(keyvals, keyVal)
			m.Mutex.Unlock()

			// Store the updated slice of key-value pairs in nodes
			m.Nodes.Store(node, keyvals)
		}
	}

}

// Get a key-value pair from the nodes in manager through hash ring.
func (m *Manager) GetKey(key string) (KeyVal, bool) {

	// Get the corresponding node in the hash ring for the key
	if node, ok := m.HashRing.GetNode(key); ok {

		// Load key-value pairs from the corresponding nodes
		if nodes, found := m.Nodes.Load(node); found {
			keyvals := nodes.([]KeyVal)

			for _, kv := range keyvals {
				if kv.Key == key {
					return kv, true
				}
			}
		}
	}

	return KeyVal{}, false

}

// Adds a node to the manager and rebalances keys in the hash ring.
func (m *Manager) AddNode(node string) {

	// Initialize the node with empty slice
	m.Nodes.Store(node, []KeyVal{})

	// Get the corresponding node in the hash ring for the node to be added.
	// Since GetNode is used to get the place in the hashring where the Key falls in,
	// node can be expressed as key to find its place in the hashring
	// The added node falls in the region of some other node and that node is the node to rebalance
	if nodeToRebalance, ok := m.HashRing.GetNode(node); ok {

		// Add the node to the hash ring
		m.HashRing.AddNode(node)

		// Get the key-value pairs from the node to rebalance and send them to the channel to rebalance keys in the hash ring.
		if nodes, found := m.Nodes.Load(nodeToRebalance); found {
			keyValsToRebalance := nodes.([]KeyVal)

			m.Wg.Add(1)
			go func() {
				defer m.Wg.Done()

				// Send keys to rebalance in batch via goroutine
				// This allows rebalance go routine to process a batch of data at a time
				// and prevent a long slice of key-value pairs while adding and removing nodes to occupy the go routine.
				for i := 0; i < len(keyValsToRebalance); i += RebalanceBatchSize {
					batch := keyValsToRebalance[i:Min(i+RebalanceBatchSize, len(keyValsToRebalance))]
					m.KeysToRebalance <- batch
					log.Printf(" Node added: %v, rebalancing %v keys, node to rebalance: %v", node, len(batch), nodeToRebalance)
				}
			}()
		}

	} else {

		// If this is the first node then just add the node to the hash ring.
		m.HashRing.AddNode(node)

	}

}

// Removes a node from the manager and rebalances keys in the hash ring.
func (m *Manager) RemoveNode(node string) {

	// First, remove the node from the hash ring
	m.HashRing.RemoveNode(node)

	// Get the corresponding node in the hash ring for the node to be removed.
	// Since GetNode is used to get the place in the hashring where the Key falls in,
	// node can be expressed as key to find its place in the hashring
	// The removed node had fallen in to some other node's region when it was added in the hash ring.
	// After removal, the keys of removed node must be mapped to the original node.
	if nodeToRebalance, ok := m.HashRing.GetNode(node); ok {

		// Load key-value pairs from the node to be deleted to rebalance
		if nodes, found := m.Nodes.Load(node); found {

			// Delete the node from the manager
			m.Nodes.Delete(node)

			keyValsToRebalance := nodes.([]KeyVal)

			m.Wg.Add(1)
			go func() {
				defer m.Wg.Done()

				// Send keys to rebalance in batch via goroutine
				// This allows rebalance go routine to process a batch of data at a time
				// and prevent a long slice of key-value pairs while adding and removing nodes to occupy the go routine.
				for i := 0; i < len(keyValsToRebalance); i += RebalanceBatchSize {
					batch := keyValsToRebalance[i:Min(i+RebalanceBatchSize, len(keyValsToRebalance))]
					m.KeysToRebalance <- batch
					log.Printf(" Node removed: %v, rebalancing %v keys, node to rebalance: %v", node, len(batch), nodeToRebalance)
				}

			}()
		}

	}

}
