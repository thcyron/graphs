package graphs

// Kruskal implements Kruskal’s algorithm. It returns a
// minimal spanning tree for the given graph.
func Kruskal(g *Graph) *Graph {
	tree := NewGraph()
	cc := map[Vertex]*Set{}

	for e := g.Edges.Front(); e != nil; e = e.Next() {
		edge := e.Value.(*Edge)

		// Initialize the start vertex’s connected component.
		if _, exists := cc[edge.Start]; !exists {
			set := NewSet()
			set.Add(edge.Start)
			cc[edge.Start] = set
		}

		// Do the same for the end vertex.
		if _, exists := cc[edge.End]; !exists {
			set := NewSet()
			set.Add(edge.End)
			cc[edge.End] = set
		}

		// If both the connected component of the start vertex
		// and the connected component of the end vertex are equal
		// a circle would occur. Better skip that edge.
		if cc[edge.Start].Equals(cc[edge.End]) {
			continue
		}

		// Add each vertex to the other’s connected component.
		cc[edge.Start].Add(edge.End)
		cc[edge.End].Add(edge.Start)

		// Update the connected components of all connected vertices.
		for _, s := range cc {
			if s.Contains(edge.Start) {
				s.Merge(cc[edge.Start])
			}
			if s.Contains(edge.End) {
				s.Merge(cc[edge.End])
			}
		}

		tree.AddEdge(edge.Start, edge.End, edge.Cost)
	}

	return tree
}
