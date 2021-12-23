package graphs

import "container/list"

type BFSWalkFunc[T Vertex] func(T, *bool)

func BFS[T Vertex](g *Graph[T], start T, walkFunc BFSWalkFunc[T]) {
	queue := list.New()
	queue.PushFront(start)

	visited := NewSet[T]()

	for f := queue.Front(); f != nil; f = queue.Front() {
		v := queue.Remove(f).(T)

		stop := false
		walkFunc(v, &stop)
		if stop {
			return
		}

		visited.Add(v)

		for he := range g.HalfedgesIter(v) {
			if !visited.Contains(he.End) {
				queue.PushBack(he.End)
			}
		}
	}
}
