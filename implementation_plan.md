# Mini-GraphDB Implementation Plan

## Goal Description
Build a lightweight, in-memory graph database in Go with support for nodes, edges, labels, properties, indexing, and a Gremlin-style query language.

## Proposed Changes

### Core Graph (`graph/`)
#### [NEW] [graph.go](file:///c:/PROJECTS/Mini-GraphDB/graph/graph.go)
- Structs: `Node`, `Edge`, `Graph`
- `Node`: ID (string/int), Label (string), Props (map[string]interface{})
- `Edge`: ID, Label, FromID, ToID, Props
- `Graph`:
    - `nodes`: map[ID]*Node
    - `edges`: map[ID]*Edge
    - `adj`: map[NodeID][]EdgeID (adjacency list for fast traversal)
    - `indexes`: map[Label]map[PropKey]map[Value][]NodeID

### Query Engine (`query/`)
#### [NEW] [query.go](file:///c:/PROJECTS/Mini-GraphDB/query/query.go)
- Chainable query builder struct `Traversal`
- Steps: `V()`, `Has(key, val)`, `Out(label)`, `In(label)`, `Values(key)`
- formatting: `g.V().Has("name", "alice").Out("knows")`

### Algorithms (`algo/`)
#### [NEW] [bfs.go](file:///c:/PROJECTS/Mini-GraphDB/algo/bfs.go)
- `ShortestPath(g *Graph, startID, endID)` returning path of nodes/edges.

## Verification Plan
### Automated Tests
- Unit tests for `Graph` CRUD.
- Unit tests for `Query` execution (mocking graph or using real graph).
- Unit tests for `BFS` correctness on known graph structures.
