package algo

import (
	. "github.com/noypi/graph/types"
)

type DFSWalkFunc func(Vertex, *bool)

func DFS(g Graph, start Vertex, walkFunc DFSWalkFunc) {
	visited := NewSet()
	stop := false
	dfs(g, start, visited, &stop, walkFunc)
}

func dfs(g Graph, start Vertex, visited *Set, stop *bool, walkFunc DFSWalkFunc) {
	visited.Add(start.StringID())

	walkFunc(start, stop)
	if *stop {
		return
	}

	ithe := g.HalfedgesIter(start)
	for ; ithe.Valid(); ithe.Next() {
		he := ithe.Value()
		if !visited.Contains(he.End()) {
			dfs(g, g.V(he.End()), visited, stop, walkFunc)
			if *stop {
				break
			}
		}
	}
}
