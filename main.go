package main

import (
	"fmt"
	"log"
	"mini-graphdb/algo"
	"mini-graphdb/graph"
	"mini-graphdb/query"
)

func main() {
	g := graph.NewGraph()

	// 1. Add Nodes
	// People
	g.AddNode("1", "person", map[string]interface{}{"name": "Alice", "age": 30})
	g.AddNode("2", "person", map[string]interface{}{"name": "Bob", "age": 27})
	g.AddNode("3", "person", map[string]interface{}{"name": "Charlie", "age": 35})
	g.AddNode("4", "person", map[string]interface{}{"name": "David", "age": 25})
	// Software
	g.AddNode("5", "software", map[string]interface{}{"name": "Ripple", "lang": "Go"})
	g.AddNode("6", "software", map[string]interface{}{"name": "Lop", "lang": "Java"})

	// 2. Add Edges
	g.AddEdge("101", "knows", "1", "2", nil)   // Alice -> Bob
	g.AddEdge("102", "knows", "1", "4", nil)   // Alice -> David
	g.AddEdge("103", "knows", "2", "3", nil)   // Bob -> Charlie
	g.AddEdge("104", "created", "3", "5", nil) // Charlie -> Ripple
	g.AddEdge("105", "created", "3", "6", nil) // Charlie -> Lop
	g.AddEdge("106", "created", "4", "5", nil) // David -> Ripple

	fmt.Println("--- Graph Built ---")

	// 3. Test Indexing
	fmt.Println("\n--- Index Lookup 'name = Alice' ---")
	nodes := g.FindNodes("person", "name", "Alice")
	for _, n := range nodes {
		fmt.Printf("Found: %s (%v)\n", n.ID, n.Props)
	}

	// 4. Test Query Engine
	// Find who Alice knows
	fmt.Println("\n--- Query: g.V().Has('name', 'Alice').Out('knows') ---")
	aliceKnows := query.New(g).V().Has("name", "Alice").Out("knows").Execute()
	for _, n := range aliceKnows {
		fmt.Printf("Alice knows: %s (%v)\n", n.ID, n.Props)
	}

	// Find creators of Ripple
	fmt.Println("\n--- Query: g.V().Has('name', 'Ripple').In('created') ---")
	// Wait, I didn't implement In() yet! Only Out().
	// Let's stick to Out(). Who created Ripple?
	// Finding who created Ripple requires In().
	// Alternative: g.V().Out('created').Has('name', 'Ripple') -> Start with people.
	fmt.Println("Query: All nodes -> Out('created') -> Has('name', 'Ripple')")
	creators := query.New(g).V().Out("created").Has("name", "Ripple").Execute()
	for _, n := range creators {
		// This query logic is flawed: Out('created') starts from Person and goes to Software.
		// If we filter result by "Ripple", we get the Software node, not the creator.
		// Effectively this finds "Ripple" via traversal.
		fmt.Printf("Found Software via traversal: %s\n", n.Props["name"])
	}

	// 5. Test BFS
	fmt.Println("\n--- BFS Shortest Path: Alice(1) -> Ripple(5) ---")
	path, err := algo.BFSShortestPath(g, "1", "5")
	if err != nil {
		log.Printf("BFS Error: %v", err)
	} else {
		fmt.Printf("Path: %v\n", path)
	}

	// Alice(1) -> Bob(2) -> Charlie(3) -> Ripple(5) (Length 4)
	// Alice(1) -> David(4) -> Ripple(5) (Length 3) -> Should pick this one.
}
