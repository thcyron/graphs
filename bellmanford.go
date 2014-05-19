package graphs

import "math"

type bellmanFordNode struct {
	cost        float64
	predecessor Vertex
}

// BellmanFord implements the Bellman-Ford algorithm. It returns
// the shortest-weight path from start to end vertex as a slice,
// or nil if the graph contains a negative-weight cycle.
func BellmanFord(g *Graph, start, end Vertex) []Vertex {
	nodes := map[Vertex]*bellmanFordNode{}

	for v := range g.VerticesIter() {
		nodes[v] = &bellmanFordNode{
			cost:        math.Inf(1),
			predecessor: nil,
		}
	}
	nodes[start].cost = 0

	n := g.NVertices()
	for i := 0; i < n-1; i++ {
		for e := range g.EdgesIter() {
			c := nodes[e.Start].cost + e.Cost
			if c < nodes[e.End].cost {
				nodes[e.End].cost = c
				nodes[e.End].predecessor = e.Start
			}
		}
	}

	// Check for negative-weight cycles.
	for e := range g.EdgesIter() {
		if nodes[e.Start].cost+e.Cost < nodes[e.End].cost {
			return nil
		}
	}

	i := 0
	for v := end; v != nil; v = nodes[v].predecessor {
		i++
	}

	path := make([]Vertex, i)
	for v := end; v != nil; v = nodes[v].predecessor {
		i--
		path[i] = v
	}

	return path
}
