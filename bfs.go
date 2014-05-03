package graphs

import "container/list"

type BFSPredicate func(Vertex) bool

func BFS(g *Graph, start Vertex, predicate BFSPredicate) Vertex {
	queue := list.New()
	queue.PushFront(start)

	visited := NewSet()

	for f := queue.Front(); f != nil; f = queue.Front() {
		v := queue.Remove(f).(Vertex)

		if predicate(v) {
			return v
		}

		visited.Add(v)

		for vv, _ := range *g.AdjacentVertices(v) {
			v2 := vv.(Vertex)
			if !visited.Contains(v2) {
				queue.PushBack(v2)
			}
		}
	}

	return nil
}
