package algo

import (
	"container/list"

	. "github.com/noypi/graph/types"
)

type tarjanNode struct {
	index   int
	lowlink int
}

// TarjanStrongCC returns a list of sets of vertices. Each set
// is a strongly connected component of the graph.
func TarjanStrongCC(g Graph) *list.List {
	ccList := list.New()
	nodes := map[string]*tarjanNode{}
	stack := list.New()

	// Loop through every vertex of the graph. If the vertex
	// has not been visited yet, perform Tarjan’s strongly
	// connected component function with that vertex.
	itv := g.VerticesIter()
	for ; itv.Valid(); itv.Next() {
		v := itv.Value()
		if _, exists := nodes[v.StringID()]; !exists {
			tarjanStrongCC(g, v, ccList, stack, nodes)
		}
	}

	return ccList
}

// tarjanStrongCC implements Tarjan’s strongly connected components
// algorithm starting with the given vertex.
func tarjanStrongCC(
	g Graph,
	v Vertex,
	ccList *list.List,
	stack *list.List,
	nodes map[string]*tarjanNode,
) {
	// When this function is called it’s certain that the given
	// vertex has not been visited yet. Create a new struct with
	// the index and lowlink information for the vertex and thus
	// mark it as visited.
	nodes[v.StringID()] = &tarjanNode{
		index:   len(nodes),
		lowlink: len(nodes),
	}

	stack.PushBack(v)

	// Loop through every adjacent vertex.
	ithe := g.HalfedgesIter(v)
	for ; ithe.Valid(); ithe.Next() {
		he := ithe.Value()
		w := he.End()

		if _, exists := nodes[w]; !exists {
			// Call ourselves recursively if that adjacent
			// vertex has not been visited yet. That’s
			// basically a DFS.
			tarjanStrongCC(g, g.V(w), ccList, stack, nodes)

			// Can the adjacent vertex w reach a lower indexed
			// vertex than vertex v can? If yes v can reach that
			// vertex, too, since v and w are connected.
			if nodes[w].lowlink < nodes[v.StringID()].lowlink {
				nodes[v.StringID()].lowlink = nodes[w].lowlink
			}

			continue
		}

		// That vertex has already been visited. Check
		// if it’s in the stack and thus part of the path
		// to the adjacent vertex w.
		for e := stack.Front(); e != nil; e = e.Next() {
			if e.Value.(Vertex).StringID() != w {
				continue
			}

			// It’s part of the path. Does that vertex w
			// improve the lowest indexed vertex the current
			// vertex v can reach? If yes, update its
			// lowlink information accordingly.
			if nodes[w].index < nodes[v.StringID()].lowlink {
				nodes[v.StringID()].lowlink = nodes[w].index
			}

			break
		}
	}

	// If the lowest indexed vertex the current vertex v can
	// reach is itself it’s a root of the strongly connected
	// component. Create a new set of vertices and add all the
	// vertices from the stack to the set until the current
	// vertex v was popped from the stack.
	if nodes[v.StringID()].lowlink == nodes[v.StringID()].index {
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
