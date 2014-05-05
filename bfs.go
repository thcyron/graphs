package graphs

import "container/list"

type BFSWalkFunc func(Vertex, *bool)

func BFS(g *Graph, start Vertex, walkFunc BFSWalkFunc) {
	queue := list.New()
	queue.PushFront(start)

	visited := NewSet()

	for f := queue.Front(); f != nil; f = queue.Front() {
		v := queue.Remove(f).(Vertex)

		stop := false
		walkFunc(v, &stop)
		if stop {
			return
		}

		visited.Add(v)

		g.AdjacentVertices(v).Each(func(e interface{}, stop *bool) {
			v := e.(Vertex)
			if !visited.Contains(v) {
				queue.PushBack(v)
			}
		})
	}
}
