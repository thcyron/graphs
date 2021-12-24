package graphs

type DFSWalkFunc[T Vertex] func(T, *bool)

func DFS[T Vertex](g *Graph[T], start T, walkFunc DFSWalkFunc[T]) {
	visited := NewSet[T]()
	stop := false
	dfs(g, start, visited, &stop, walkFunc)
}

func dfs[T Vertex](g *Graph[T], start T, visited *Set[T], stop *bool, walkFunc DFSWalkFunc[T]) {
	visited.Add(start)

	walkFunc(start, stop)
	if *stop {
		return
	}

	g.EachHalfedge(start, func(he Halfedge[T], innerStop func()) {
		if !visited.Contains(he.End) {
			dfs(g, he.End, visited, stop, walkFunc)
			if *stop {
				innerStop()
			}
		}
	})
}
