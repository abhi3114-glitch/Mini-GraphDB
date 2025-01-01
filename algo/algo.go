package algo

import (
	"fmt"
	"mini-graphdb/graph"
)

// BFSShortestPath finds the shortest path between startID and endID using BFS.
// Returns a slice of NodeIDs representing the path.
func BFSShortestPath(g *graph.Graph, startID, endID string) ([]string, error) {
	// Check if nodes exist
	if _, err := g.GetNode(startID); err != nil {
		return nil, err
	}
	if _, err := g.GetNode(endID); err != nil {
		return nil, err
	}

	queue := []string{startID}
	visited := make(map[string]bool)
	visited[startID] = true

	// parent map to reconstruct path: nodeID -> parentID
	parent := make(map[string]string)

	found := false
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == endID {
			found = true
			break
		}

		// Get outgoing neighbors
		neighbors := g.GetNeighbors(curr, "")
		for _, n := range neighbors {
			if !visited[n.ID] {
				visited[n.ID] = true
				parent[n.ID] = curr
				queue = append(queue, n.ID)
			}
		}
	}

	if !found {
		return nil, fmt.Errorf("path not found between %s and %s", startID, endID)
	}

	// Reconstruct path
	path := []string{}
	curr := endID
	for curr != "" {
		path = append([]string{curr}, path...) // Prepend
		if curr == startID {
			break
		}
		curr = parent[curr]
	}

	return path, nil
}

// DFSTraversal performs a Depth-First Search starting from startID.
// It calls the visit function for each visited node.
func DFSTraversal(g *graph.Graph, startID string, visit func(n *graph.Node)) error {
	startNode, err := g.GetNode(startID)
	if err != nil {
		return err
	}

	visited := make(map[string]bool)
	stack := []*graph.Node{startNode}

	for len(stack) > 0 {
		// Pop
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[curr.ID] {
			visited[curr.ID] = true
			visit(curr)

			// neighbors
			neighbors := g.GetNeighbors(curr.ID, "")
			// Push neighbors (in reverse order to preserve some order if needed, but not strictly required for generic DFS)
			for i := len(neighbors) - 1; i >= 0; i-- {
				stack = append(stack, neighbors[i])
			}
		}
	}
	return nil
}
