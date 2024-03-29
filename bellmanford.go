package graphs

import "math"

type bellmanFordNode[T Vertex] struct {
	cost        float64
	predecessor *T
}

// BellmanFord implements the Bellman-Ford algorithm. It returns
// the shortest-weight path from start to end vertex as a slice,
// or nil if the graph contains a negative-weight cycle.
func BellmanFord[T Vertex](g *Graph[T], start, end T) []T {
	nodes := map[T]*bellmanFordNode[T]{}

	g.EachVertex(func(v T, _ func()) {
		nodes[v] = &bellmanFordNode[T]{
			cost:        math.Inf(1),
			predecessor: nil,
		}
	})
	nodes[start].cost = 0

	n := g.NVertices()
	for i := 0; i < n-1; i++ {
		g.EachEdge(func(e Edge[T], _ func()) {
			c := nodes[e.Start].cost + e.Cost
			if c < nodes[e.End].cost {
				nodes[e.End].cost = c
				start := e.Start
				nodes[e.End].predecessor = &start
			}
		})
	}

	// Check for negative-weight cycles.
	hasNegativeWeightCycle := false
	g.EachEdge(func(e Edge[T], stop func()) {
		if nodes[e.Start].cost+e.Cost < nodes[e.End].cost {
			hasNegativeWeightCycle = true
			stop()
		}
	})
	if hasNegativeWeightCycle {
		return nil
	}

	i := 0
	for v := &end; v != nil; v = nodes[*v].predecessor {
		i++
	}

	path := make([]T, i)
	for v := &end; v != nil; v = nodes[*v].predecessor {
		i--
		path[i] = *v
	}

	return path
}
