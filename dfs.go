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

	set := g.AdjacentVertices(start)

	for v, _ := range *set {
		vv := v.(Vertex)
		if !visited.Contains(vv) {
			dfs(g, vv, visited, stop, walkFunc)
			if *stop {
				return
			}
		}
	}
}
