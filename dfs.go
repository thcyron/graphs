package graphs

type DFSWalkFunc func(Vertex, *bool)

func DFS(g *Graph, start Vertex, walkFunc DFSWalkFunc) {
	visited := NewSet()
	stop := false
	dfs(g, start, visited, &stop, walkFunc)
}

func dfs(g *Graph, start Vertex, visited *Set, stop *bool, walkFunc DFSWalkFunc) {
	visited.Add(start)

	walkFunc(start, stop)
	if *stop {
		return
	}

	for he := range g.HalfedgesIter(start) {
		if !visited.Contains(he.End) {
			dfs(g, he.End, visited, stop, walkFunc)
			if *stop {
				break
			}
		}
	}
}
