package algo

import (
	. "github.com/noypi/graph/types"
)

// FloydWarshall implements the Floyd–Warshall algorithm. It returns
// the cost matrix for each vertex to each other vertex of the given
// graph. It does not check for negative weight cycles.
func FloydWarshall(g Graph) map[string]map[string]float64 {
	m := make(map[string]map[string]float64)

	// Initialize matrix m.
	itv := g.VerticesIter()
	for ; itv.Valid(); itv.Next() {
		v := itv.Value()
		m[v.StringID()] = make(map[string]float64)

		ithe := g.HalfedgesIter(v)
		for ; ithe.Valid(); ithe.Next() {
			he := ithe.Value()
			m[v.StringID()][he.End()] = he.Cost()
		}
	}

	// For each vertex v check if using it as an intermediate
	// vertex in the path from u -> w (u -> v -> w) results in
	// a shorter path.
	itv2 := g.VerticesIter()
	for ; itv2.Valid(); itv2.Next() {
		v := itv2.Value()
		itu := g.VerticesIter()
		for ; itu.Valid(); itu.Next() {
			u := itu.Value()
			itw := g.VerticesIter()
			for ; itw.Valid(); itw.Next() {
				w := itw.Value()
				if _, exists := m[u.StringID()][v.StringID()]; !exists {
					continue
				}
				if _, exists := m[v.StringID()][w.StringID()]; !exists {
					continue
				}

				newCost := m[u.StringID()][v.StringID()] + m[v.StringID()][w.StringID()]

				if oldCost, exists := m[u.StringID()][w.StringID()]; exists {
					if newCost < oldCost {
						m[u.StringID()][w.StringID()] = newCost
					}
				} else {
					m[u.StringID()][w.StringID()] = newCost
				}
			}
		}
	}

	return m
}
