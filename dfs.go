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

	g.AdjacentVertices(start).Each(func(e interface{}, vstop *bool) {
		v := e.(Vertex)
		if !visited.Contains(v) {
			dfs(g, v, visited, stop, walkFunc)
			if *stop {
				*vstop = true
			}
		}
	})
}
