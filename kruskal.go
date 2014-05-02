package graphs

// Kruskal implements Kruskal’s algorithm. It returns a
// minimal spanning tree for the given graph.
func Kruskal(g *Graph) *Graph {
	tree := NewGraph()
	cc := map[Vertex]int{}
	ccid := 1

	for _, edge := range g.SortedEdges() {
		// Add the start vertex to the connected
		// component if it isn’t included yet.
		if _, exists := cc[edge.Start]; !exists {
			cc[edge.Start] = ccid
			ccid++
		}

		// If both the start and end vertex are in the
		// same connected component a cycle would occur,
		// so don’t add that edge to the spanning tree.
		if cc[edge.Start] == cc[edge.End] {
			continue
		}

		// If the end vertex has a valid connected component
		// set all vertices with that ID to the ID of the
		// start vertex, set it to the ID of the start vertex
		// otherwise.
		if cc[edge.End] != 0 {
			endid := cc[edge.End]
			for v, id := range cc {
				if id == endid {
					cc[v] = cc[edge.Start]
				}
			}
		} else {
			cc[edge.End] = cc[edge.Start]
		}

		tree.AddEdge(edge.Start, edge.End, edge.Cost)
	}

	return tree
}
