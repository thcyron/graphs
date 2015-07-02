package graphs

import (
	"container/list"
	"errors"
)

type NoDAGError struct {
	error
}

func initInEdges(g *Graph) map[Vertex]int {
	inEdges := make(map[Vertex]int)
	for e := range g.EdgesIter() {
		if _, ok := inEdges[e.Start]; !ok {
			inEdges[e.Start] = 0
		}
		if _, ok := inEdges[e.End]; ok {
			inEdges[e.End]++
		} else {
			inEdges[e.End] = 1
		}
	}
	return inEdges
}

func removeEdgesFromVertex(v Vertex, g *Graph, inEdges map[Vertex]int) {
	for outEdge := range g.HalfedgesIter(v) {
		neighbor := outEdge.End
		inEdges[neighbor]--
	}
}

func TopologicalSort(g *Graph) (topologicalOrder *list.List, topologicalClasses map[Vertex]int, err *NoDAGError) {
	inEdges := initInEdges(g)

	topologicalClasses = make(map[Vertex]int)
	topologicalOrder = list.New()
	err = nil
	tClass := 0
	for len(inEdges) > 0 {
		topClass := []Vertex{}
		for v, inDegree := range inEdges {
			if inDegree == 0 {
				topClass = append(topClass, v)
				topologicalClasses[v] = tClass
			}
		}
		if len(topClass) == 0 {
			err = &NoDAGError{errors.New("Graph is not a DAG.")}
			topologicalClasses = make(map[Vertex]int)
			topologicalOrder = list.New()
			return
		}
		for _, v := range topClass {
			removeEdgesFromVertex(v, g, inEdges)
			delete(inEdges, v)
			topologicalOrder.PushBack(v)
		}
		tClass++
	}

	return
}
