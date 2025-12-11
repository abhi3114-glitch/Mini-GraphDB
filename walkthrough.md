# Mini-GraphDB Walkthrough

We have successfully built a lightweight, in-memory graph database in Go.

## Features Implemented
- **Core Graph**: `Node` and `Edge` support with property maps.
- **Indexing**: Automatic indexing of all properties for O(1) lookups.
- **Query Engine**: Gremlin-style chainable builder (`g.V().Has().Out()`).
- **Algorithms**: BFS Shortest Path finding.

## Verification Results

We ran a demo script (`main.go`) creating a social graph of People needing Software.

### Graph Structure
- **Nodes**: Alice, Bob, Charlie, David (Person); Ripple, Lop (Software).
- **Edges**: Knows, Created.

### Test Output
```text
--- Graph Built ---

--- Index Lookup 'name = Alice' ---
Found: 1 (map[age:30 name:Alice])

--- Query: g.V().Has('name', 'Alice').Out('knows') ---
Alice knows: 2 (map[age:27 name:Bob])
Alice knows: 4 (map[age:25 name:David])

--- Query: All nodes -> Out('created') -> Has('name', 'Ripple') ---
Found Software via traversal: Ripple

--- BFS Shortest Path: Alice(1) -> Ripple(5) ---
Path: [1 4 5]
```
*Successfully found the shortest path from Alice to Ripple via David (Length 3), ignoring the longer path via Bob & Charlie.*

## Code Structure
- `graph/`: Core `Graph`, `Node`, `Edge` structs and Indexing logic.
- `query/`: Query builder and execution.
- `algo/`: Graph algorithms (BFS).
- `main.go`: API demonstration.
