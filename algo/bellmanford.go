package algo

import (
	"math"

	. "github.com/noypi/graph/types"
)

type bellmanFordNode struct {
	cost        float64
	predecessor string //vertex id
	v           Vertex
}

// BellmanFord implements the Bellman-Ford algorithm. It returns
// the shortest-weight path from start to end vertex as a slice,
// or nil if the graph contains a negative-weight cycle.
func BellmanFord(g Graph, start, end Vertex) []Vertex {
	nodes := map[string]*bellmanFordNode{}

	itv := g.VerticesIter()
	for ; itv.Valid(); itv.Next() {
		v := itv.Value()
		nodes[v.StringID()] = &bellmanFordNode{
			cost:        math.Inf(1),
			predecessor: "",
			v:           itv.Value(),
		}
	}
	nodes[start.StringID()].cost = 0

	n := g.NVertices()
	for i := 0; i < n-1; i++ {
		itedge := g.EdgesIter()
		for ; itedge.Valid(); itedge.Next() {
			e := itedge.Value()
			c := nodes[e.Start()].cost + e.Cost()
			if c < nodes[e.End()].cost {
				nodes[e.End()].cost = c
				nodes[e.End()].predecessor = e.Start()
			}
		}
	}

	// Check for negative-weight cycles.
	itedge := g.EdgesIter()
	for ; itedge.Valid(); itedge.Next() {
		e := itedge.Value()
		if nodes[e.Start()].cost+e.Cost() < nodes[e.End()].cost {
			return nil
		}
	}

	i := 0
	for v := end.StringID(); v != ""; v = nodes[v].predecessor {
		i++
	}

	path := make([]Vertex, i)
	for v := end.StringID(); v != ""; v = nodes[v].predecessor {
		i--
		if node, has := nodes[v]; has {
			path[i] = node.v
		}
	}

	return path
}
