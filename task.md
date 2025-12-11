# Tasks: Mini-GraphDB

- [x] Core Graph Data Structure <!-- id: 0 -->
    - [x] Define `Node` and `Edge` structs with properties and labels <!-- id: 1 -->
    - [x] Implement `Graph` struct with in-memory storage (maps) <!-- id: 2 -->
    - [x] Implement CRUD operations (AddNode, AddEdge, GetNode, GetEdge) <!-- id: 3 -->
- [x] Indexing System <!-- id: 4 -->
    - [x] Implement property indexing (map of value -> node IDs) <!-- id: 5 -->
    - [x] Integrate indexing with Graph CRUD operations <!-- id: 6 -->
- [x] Query Engine (Gremlin-style) <!-- id: 7 -->
    - [x] Implement query builder/traversal struct (`g.V()`, `.Has()`, `.Out()`, etc.) <!-- id: 8 -->
    - [x] Implement step execution logic <!-- id: 9 -->
- [x] Graph Algorithms <!-- id: 10 -->
    - [x] Implement BFS Shortest Path <!-- id: 11 -->
    - [x] Implement DFS (generic traversal or specific use case) <!-- id: 12 -->
- [x] Verification & Demo <!-- id: 13 -->
    - [x] Unit tests for core graph, index, and query engine <!-- id: 14 -->
    - [x] Create a demo script showing graph creation, querying, and algorithms <!-- id: 15 -->
