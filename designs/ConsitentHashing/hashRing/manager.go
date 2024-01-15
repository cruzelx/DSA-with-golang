package main

import (
	"log"
	"sync"
)

const (
	NumReplicaNodes    = 15
	RebalanceBatchSize = 10
)

type KeyVal struct {
	Key   string
	Value string
}

type Manager struct {
	Nodes           sync.Map
	HashRing        *HashRing
	Mutex           sync.RWMutex
	KeysToRebalance chan []KeyVal
	Done            chan struct{}
	Wg              sync.WaitGroup
}

// make struct channel

func NewManager() *Manager {
	m := &Manager{
		HashRing:        NewHashRing(NumReplicaNodes),
		Mutex:           sync.RWMutex{},
		KeysToRebalance: make(chan []KeyVal),
		Done:            make(chan struct{}),
		Wg:              sync.WaitGroup{},
	}
	go m.rebalance()
	return m

}

func (m *Manager) rebalance() {

	for {
		select {
		case keyValsToRebalance := <-m.KeysToRebalance:
			m.Wg.Add(1)
			go m.rebalanceKeys(keyValsToRebalance)
		case <-m.Done:
			close(m.KeysToRebalance)
			return

		}
	}

}

func (m *Manager) rebalanceKeys(keyValsToRebalance []KeyVal) {
	defer m.Wg.Done()

	for _, kv := range keyValsToRebalance {
		if node, ok := m.HashRing.GetNode(kv.Key); ok {
			if nodes, found := m.Nodes.LoadOrStore(node, []KeyVal{kv}); found {
				m.Mutex.Lock()
				keyVals := nodes.([]KeyVal)
				keyVals = append(keyVals, kv)
				m.Mutex.Unlock()
				m.Nodes.Store(node, keyVals)
			}

		}
	}
}

func (m *Manager) Close() {
	m.Done <- struct{}{}
	m.Wg.Wait()

	close(m.Done)
}

func (m *Manager) PutKey(keyVal KeyVal) {

	if node, ok := m.HashRing.GetNode(keyVal.Key); ok {
		if nodes, loaded := m.Nodes.LoadOrStore(node, []KeyVal{keyVal}); loaded {
			m.Mutex.Lock()
			keyvals := nodes.([]KeyVal)
			keyvals = append(keyvals, keyVal)
			m.Mutex.Unlock()
			m.Nodes.Store(node, keyvals)
		}
	}

}

func (m *Manager) GetKey(key string) (KeyVal, bool) {

	if node, ok := m.HashRing.GetNode(key); ok {
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

func (m *Manager) AddNode(node string) {
	m.Nodes.Store(node, []KeyVal{})

	if nodeToRebalance, ok := m.HashRing.GetNode(node); ok {
		m.HashRing.AddNode(node)

		if nodes, found := m.Nodes.Load(nodeToRebalance); found {
			keyValsToRebalance := nodes.([]KeyVal)

			// send keys to rebalance in batch via goroutine
			m.Wg.Add(1)
			go func() {
				defer m.Wg.Done()
				for i := 0; i < len(keyValsToRebalance); i += RebalanceBatchSize {
					batch := keyValsToRebalance[i:Min(i+RebalanceBatchSize, len(keyValsToRebalance))]
					m.KeysToRebalance <- batch
					log.Printf(" Node added: %v, rebalancing %v keys, node to rebalance: %v", node, len(batch), nodeToRebalance)
				}
			}()
		}

	} else {
		m.HashRing.AddNode(node)

	}

}

func (m *Manager) RemoveNode(node string) {
	m.HashRing.RemoveNode(node)

	if nodeToRebalance, ok := m.HashRing.GetNode(node); ok {
		if nodes, found := m.Nodes.Load(node); found {
			m.Nodes.Delete(node)

			keyValsToRebalance := nodes.([]KeyVal)

			// send keys to rebalance in batch via goroutine
			m.Wg.Add(1)
			go func() {
				defer m.Wg.Done()

				for i := 0; i < len(keyValsToRebalance); i += RebalanceBatchSize {
					batch := keyValsToRebalance[i:Min(i+RebalanceBatchSize, len(keyValsToRebalance))]
					m.KeysToRebalance <- batch
					log.Printf(" Node removed: %v, rebalancing %v keys, node to rebalance: %v", node, len(batch), nodeToRebalance)
				}

			}()
		}

	}

}
