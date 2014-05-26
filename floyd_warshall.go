package graphs

// FloydWarshall implements the Floyd–Warshall algorithm. It returns
// the cost matrix for each vertex to each other vertex of the given
// graph. It does not check for negative weight cycles.
func FloydWarshall(g *Graph) map[Vertex]map[Vertex]float64 {
	m := make(map[Vertex]map[Vertex]float64)

	// Initialize matrix m.
	for v := range g.VerticesIter() {
		m[v] = make(map[Vertex]float64)

		for he := range g.HalfedgesIter(v) {
			m[v][he.End] = he.Cost
		}
	}

	// For each vertex v check if using it as an intermediate
	// vertex in the path from u -> w (u -> v -> w) results in
	// a shorter path.
	for v := range g.VerticesIter() {
		for u := range g.VerticesIter() {
			for w := range g.VerticesIter() {
				if _, exists := m[u][v]; !exists {
					continue
				}
				if _, exists := m[v][w]; !exists {
					continue
				}

				newCost := m[u][v] + m[v][w]

				if oldCost, exists := m[u][w]; exists {
					if newCost < oldCost {
						m[u][w] = newCost
					}
				} else {
					m[u][w] = newCost
				}
			}
		}
	}

	return m
}
