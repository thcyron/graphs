package graphs

import (
	"container/list"
	"errors"
)

var ErrNoDAG = errors.New("graphs: graph is not a DAG")

func TopologicalSort[T Vertex](g *Graph[T]) (topologicalOrder *list.List, topologicalClasses map[T]int, err error) {
	inEdges := make(map[T]int)
	g.EachEdge(func(e Edge[T], _ func()) {
		if _, ok := inEdges[e.Start]; !ok {
			inEdges[e.Start] = 0
		}
		if _, ok := inEdges[e.End]; ok {
			inEdges[e.End]++
		} else {
			inEdges[e.End] = 1
		}
	})

	removeEdgesFromVertex := func(v T) {
		g.EachHalfedge(v, func(outEdge Halfedge[T], _ func()) {
			neighbor := outEdge.End
			inEdges[neighbor]--
		})
	}

	topologicalClasses = make(map[T]int)
	topologicalOrder = list.New()
	tClass := 0
	for len(inEdges) > 0 {
		topClass := []T{}
		for v, inDegree := range inEdges {
			if inDegree == 0 {
				topClass = append(topClass, v)
				topologicalClasses[v] = tClass
			}
		}
		if len(topClass) == 0 {
			err = ErrNoDAG
			topologicalClasses = make(map[T]int)
			topologicalOrder = list.New()
			return
		}
		for _, v := range topClass {
			removeEdgesFromVertex(v)
			delete(inEdges, v)
			topologicalOrder.PushBack(v)
		}
		tClass++
	}

	return
}
