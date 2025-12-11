package query

import (
	"mini-graphdb/graph"
)

// Traversal represents a graph traversal state
type Traversal struct {
	g       *graph.Graph
	results []*graph.Node
}

// New creates a new traversal starting with the valid graph
func New(g *graph.Graph) *Traversal {
	return &Traversal{
		g:       g,
		results: []*graph.Node{},
	}
}

// V starts the traversal with all nodes (or specific IDs if provided, implemented later)
// For now, V() loads ALL nodes.
func (t *Traversal) V() *Traversal {
	t.results = t.g.GetAllNodes()
	return t
}

// Has filters the current nodes by property key and value
func (t *Traversal) Has(key string, val interface{}) *Traversal {
	var next []*graph.Node
	for _, n := range t.results {
		if v, ok := n.Props[key]; ok && v == val {
			next = append(next, n)
		}
	}
	t.results = next
	return t
}

// HasLabel filters by node label
func (t *Traversal) HasLabel(label string) *Traversal {
	var next []*graph.Node
	for _, n := range t.results {
		if n.Label == label {
			next = append(next, n)
		}
	}
	t.results = next
	return t
}

// Out traverses outgoing edges with the given label to finding adjacent nodes
func (t *Traversal) Out(edgeLabel string) *Traversal {
	var next []*graph.Node
	// Use a map to deduplicate if multiple paths lead to same node?
	// Gremlin `out` usually returns ALL occurrences (bag semantics), not set.
	// But let's stick to simple list for now.

	for _, n := range t.results {
		neighbors := t.g.GetNeighbors(n.ID, edgeLabel)
		next = append(next, neighbors...)
	}
	t.results = next
	return t
}

// Values extracts property values from the current nodes
func (t *Traversal) Values(key string) []interface{} {
	var vals []interface{}
	for _, n := range t.results {
		if v, ok := n.Props[key]; ok {
			vals = append(vals, v)
		}
	}
	return vals
}

// Execute returns the final list of nodes
func (t *Traversal) Execute() []*graph.Node {
	return t.results
}
