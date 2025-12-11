package graph

import (
	"fmt"
	"sync"
)

// Node represents a vertex in the graph
type Node struct {
	ID    string
	Label string
	Props map[string]interface{}
}

// Edge represents a connection between two nodes
type Edge struct {
	ID    string
	Label string
	From  string
	To    string
	Props map[string]interface{}
}

// Graph is the in-memory graph database
type Graph struct {
	mu     sync.RWMutex
	nodes  map[string]*Node
	edges  map[string]*Edge
	outAdj map[string][]string
	inAdj  map[string][]string

	// Index: Label -> PropertyKey -> Value -> []NodeID
	indexes map[string]map[string]map[interface{}][]string
}

// NewGraph creates a new empty graph
func NewGraph() *Graph {
	return &Graph{
		nodes:   make(map[string]*Node),
		edges:   make(map[string]*Edge),
		outAdj:  make(map[string][]string),
		inAdj:   make(map[string][]string),
		indexes: make(map[string]map[string]map[interface{}][]string),
	}
}

// AddNode adds a node to the graph and updates indexes
func (g *Graph) AddNode(id, label string, props map[string]interface{}) (*Node, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if _, exists := g.nodes[id]; exists {
		return nil, fmt.Errorf("node with ID %s already exists", id)
	}

	n := &Node{
		ID:    id,
		Label: label,
		Props: props,
	}
	g.nodes[id] = n
	g.outAdj[id] = []string{}
	g.inAdj[id] = []string{}

	// Update Indexes
	for key, val := range props {
		g.addToIndex(label, key, val, id)
	}

	return n, nil
}

// addToIndex helper (caller must hold lock)
func (g *Graph) addToIndex(label, key string, val interface{}, nodeID string) {
	if _, ok := g.indexes[label]; !ok {
		g.indexes[label] = make(map[string]map[interface{}][]string)
	}
	if _, ok := g.indexes[label][key]; !ok {
		g.indexes[label][key] = make(map[interface{}][]string)
	}
	g.indexes[label][key][val] = append(g.indexes[label][key][val], nodeID)
}

// GetNode retrieves a node by ID
func (g *Graph) GetNode(id string) (*Node, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	n, exists := g.nodes[id]
	if !exists {
		return nil, fmt.Errorf("node with ID %s not found", id)
	}
	return n, nil
}

// AddEdge adds an edge between two nodes
func (g *Graph) AddEdge(id, label, fromID, toID string, props map[string]interface{}) (*Edge, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if _, exists := g.edges[id]; exists {
		return nil, fmt.Errorf("edge with ID %s already exists", id)
	}

	if _, exists := g.nodes[fromID]; !exists {
		return nil, fmt.Errorf("from node %s does not exist", fromID)
	}
	if _, exists := g.nodes[toID]; !exists {
		return nil, fmt.Errorf("to node %s does not exist", toID)
	}

	e := &Edge{
		ID:    id,
		Label: label,
		From:  fromID,
		To:    toID,
		Props: props,
	}
	g.edges[id] = e

	g.outAdj[fromID] = append(g.outAdj[fromID], id)
	g.inAdj[toID] = append(g.inAdj[toID], id)

	return e, nil
}

// GetEdge retrieves an edge by ID
func (g *Graph) GetEdge(id string) (*Edge, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	e, exists := g.edges[id]
	if !exists {
		return nil, fmt.Errorf("edge with ID %s not found", id)
	}
	return e, nil
}

// FindNodes uses the index to find nodes by label and property
func (g *Graph) FindNodes(label, key string, val interface{}) []*Node {
	g.mu.RLock()
	defer g.mu.RUnlock()

	var results []*Node

	// Check index
	if labelIdx, ok := g.indexes[label]; ok {
		if keyIdx, ok := labelIdx[key]; ok {
			if ids, ok := keyIdx[val]; ok {
				for _, id := range ids {
					if n, exists := g.nodes[id]; exists {
						results = append(results, n)
					}
				}
				return results
			}
		}
	}

	// Fallback scan if not indexed?
	// For now, we are indexing EVERYTHING on AddNode, so if it's not in index, it's not there.
	// However, complex queries might need scans. But for FindNodes(label, key, val) we rely on index.
	return results
}

// GetAllNodes returns all nodes in the graph (scans everything)
func (g *Graph) GetAllNodes() []*Node {
	g.mu.RLock()
	defer g.mu.RUnlock()

	nodes := make([]*Node, 0, len(g.nodes))
	for _, n := range g.nodes {
		nodes = append(nodes, n)
	}
	return nodes
}

// GetNeighbors returns outgoing neighbor nodes for a given node ID, optionally filtering by edge label
func (g *Graph) GetNeighbors(nodeID string, edgeLabel string) []*Node {
	g.mu.RLock()
	defer g.mu.RUnlock()

	var neighbors []*Node
	edgeIDs, ok := g.outAdj[nodeID]
	if !ok {
		return neighbors
	}

	for _, eID := range edgeIDs {
		edge, exists := g.edges[eID]
		if !exists {
			continue
		}
		if edgeLabel != "" && edge.Label != edgeLabel {
			continue
		}

		if targetNode, ok := g.nodes[edge.To]; ok {
			neighbors = append(neighbors, targetNode)
		}
	}
	return neighbors
}
