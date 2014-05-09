package graphs

import "container/list"

type tarjanNode struct {
	index   int
	lowlink int
}

// TarjanStrongCC returns a list of sets of vertices. Each set
// is a strongly connected component of the graph.
func TarjanStrongCC(g *Graph) *list.List {
	ccList := list.New()
	nodes := map[Vertex]*tarjanNode{}
	stack := list.New()

	// Loop through every vertex of the graph. If the vertex
	// has not been visited yet, perform Tarjan’s strongly
	// connected component function with that vertex.
	for v := range g.VerticesIter() {
		if _, exists := nodes[v]; !exists {
			tarjanStrongCC(g, v, ccList, stack, nodes)
		}
	}

	return ccList
}

// tarjanStrongCC implements Tarjan’s strongly connected components
// algorithm starting with the given vertex.
func tarjanStrongCC(
	g *Graph,
	v Vertex,
	ccList *list.List,
	stack *list.List,
	nodes map[Vertex]*tarjanNode,
) {
	// When this function is called it’s certain that the given
	// vertex has not been visited yet. Create a new struct with
	// the index and lowlink information for the vertex and thus
	// mark it as visited.
	nodes[v] = &tarjanNode{
		index:   len(nodes),
		lowlink: len(nodes),
	}

	stack.PushBack(v)

	// Loop through every adjacent vertex.
	for he := range g.HalfedgesIter(v) {
		w := he.End

		if _, exists := nodes[w]; !exists {
			// Call ourselves recursively if that adjacent
			// vertex has not been visited yet. That’s
			// basically a DFS.
			tarjanStrongCC(g, w, ccList, stack, nodes)

			// Can the adjacent vertex w reach a lower indexed
			// vertex than vertex v can? If yes v can reach that
			// vertex, too, since v and w are connected.
			if nodes[w].lowlink < nodes[v].lowlink {
				nodes[v].lowlink = nodes[w].lowlink
			}

			continue
		}

		// That vertex has already been visited. Check
		// if it’s in the stack and thus part of the path
		// to the adjacent vertex w.
		for e := stack.Front(); e != nil; e = e.Next() {
			if e.Value.(Vertex) != w {
				continue
			}

			// It’s part of the path. Does that vertex w
			// improve the lowest indexed vertex the current
			// vertex v can reach? If yes, update its
			// lowlink information accordingly.
			if nodes[w].index < nodes[v].lowlink {
				nodes[v].lowlink = nodes[w].index
			}

			break
		}
	}

	// If the lowest indexed vertex the current vertex v can
	// reach is itself it’s a root of the strongly connected
	// component. Create a new set of vertices and add all the
	// vertices from the stack to the set until the current
	// vertex v was popped from the stack.
	if nodes[v].lowlink == nodes[v].index {
		set := NewSet()
		for {
			w := stack.Remove(stack.Back()).(Vertex)
			set.Add(w)

			if w == v {
				break
			}
		}
		ccList.PushBack(set)
	}
}
