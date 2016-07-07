package algo

import (
	"container/list"
	"errors"

	. "github.com/noypi/graph/types"
)

var ErrNoDAG = errors.New("graphs: graph is not a DAG")

func TopologicalSort(g Graph) (topologicalOrder *list.List, topologicalClasses map[string]int, err error) {
	inEdges := make(map[string]int)
	ite := g.EdgesIter()
	for ; ite.Valid(); ite.Next() {
		e := ite.Value()
		if _, ok := inEdges[e.Start()]; !ok {
			inEdges[e.Start()] = 0
		}
		if _, ok := inEdges[e.End()]; ok {
			inEdges[e.End()]++
		} else {
			inEdges[e.End()] = 1
		}
	}

	removeEdgesFromVertex := func(vid string) {
		v := g.V(vid)
		itoutEdge := g.HalfedgesIter(v)
		for ; itoutEdge.Valid(); itoutEdge.Next() {
			outEdge := itoutEdge.Value()
			neighbor := outEdge.End()
			inEdges[neighbor]--
		}
	}

	topologicalClasses = make(map[string]int)
	topologicalOrder = list.New()
	tClass := 0
	for len(inEdges) > 0 {
		topClass := []string{}
		for v, inDegree := range inEdges {
			if inDegree == 0 {
				topClass = append(topClass, v)
				topologicalClasses[v] = tClass
			}
		}
		if len(topClass) == 0 {
			err = ErrNoDAG
			topologicalClasses = make(map[string]int)
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
