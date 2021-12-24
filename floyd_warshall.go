package graphs

// FloydWarshall implements the Floyd–Warshall algorithm. It returns
// the cost matrix for each vertex to each other vertex of the given
// graph. It does not check for negative weight cycles.
func FloydWarshall[T Vertex](g *Graph[T]) map[T]map[T]float64 {
	m := make(map[T]map[T]float64)

	// Initialize matrix m.
	g.EachVertex(func(v T, _ func()) {
		m[v] = make(map[T]float64)

		g.EachHalfedge(v, func(he Halfedge[T], _ func()) {
			m[v][he.End] = he.Cost
		})
	})

	// For each vertex v check if using it as an intermediate
	// vertex in the path from u -> w (u -> v -> w) results in
	// a shorter path.
	g.EachVertex(func(v T, _ func()) {
		g.EachVertex(func(u T, _ func()) {
			g.EachVertex(func(w T, _ func()) {
				if _, exists := m[u][v]; !exists {
					return
				}
				if _, exists := m[v][w]; !exists {
					return
				}

				newCost := m[u][v] + m[v][w]

				if oldCost, exists := m[u][w]; exists {
					if newCost < oldCost {
						m[u][w] = newCost
					}
				} else {
					m[u][w] = newCost
				}
			})
		})
	})

	return m
}
