# Mini-GraphDB

Mini-GraphDB is a lightweight, in-memory graph database implemented in Go. It supports core graph data structures (nodes, edges), property indexing for fast lookups, and a Gremlin-style query language for graph traversals.

## Features

- **In-Memory Storage**: Fast and simple in-memory graph structure.
- **Graph Primitives**: Support for Nodes and Edges with labels and arbitrary property maps.
- **Indexing**: Automatic indexing of node properties for O(1) retrieval.
- **Query Engine**: Chainable query builder inspired by Gremlin (e.g., `g.V().Has(...).Out(...)`).
- **Algorithms**: Built-in BFS (Breadth-First Search) for shortest path calculations.
- **DFS Traversal**: Generic Depth-First Search traversal support.

## Getting Started

### Prerequisites

- Go 1.18 or higher

### Installation

Clone the repository:

```bash
git clone https://github.com/abhi3114-glitch/Mini-GraphDB.git
cd Mini-GraphDB
```

### Usage

Run the demo application to see the graph database in action:

```bash
go run main.go
```

### Example Code

Here is a simple example of how to use Mini-GraphDB programmatically:

```go
package main

import (
	"fmt"
	"mini-graphdb/graph"
	"mini-graphdb/query"
)

func main() {
	g := graph.NewGraph()

	// Add Nodes
	g.AddNode("1", "person", map[string]interface{}{"name": "Alice"})
	g.AddNode("2", "person", map[string]interface{}{"name": "Bob"})

	// Add Edge
	g.AddEdge("101", "knows", "1", "2", nil)

	// Query
	results := query.New(g).V().Has("name", "Alice").Out("knows").Execute()
	
	for _, n := range results {
		fmt.Printf("Alice knows: %v\n", n.Props["name"])
	}
}
```

## Project Structure

- `graph/`: Core data structures (Graph, Node, Edge) and indexing logic.
- `query/`: Query engine implementation.
- `algo/`: Graph algorithms (BFS, DFS).
- `main.go`: Demo application.

## License

This project is open source and available under the MIT License.
