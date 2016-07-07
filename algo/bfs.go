package algo

import (
	"container/list"

	. "github.com/noypi/graph/types"
)

type BFSWalkFunc func(Vertex, *bool)

func BFS(g Graph, start Vertex, walkFunc BFSWalkFunc) {
	queue := list.New()
	queue.PushFront(start.StringID())

	visited := NewSet()

	for f := queue.Front(); f != nil; f = queue.Front() {
		vid := queue.Remove(f).(string)
		v := g.V(vid)

		stop := false
		walkFunc(v, &stop)
		if stop {
			return
		}

		visited.Add(v)

		ithe := g.HalfedgesIter(v)
		for ; ithe.Valid(); ithe.Next() {
			he := ithe.Value()
			if !visited.Contains(he.End()) {
				queue.PushBack(he.End())
			}
		}
	}
}
