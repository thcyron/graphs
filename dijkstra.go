package graphs

import (
	"container/heap"
	"container/list"
	"math"
)

type dijkstraNode[T Vertex] struct {
	vertex      T
	distance    float64
	predecessor *dijkstraNode[T]
	index       int // Index of the node in the heap
}

type dijkstraPQ[T Vertex] []*dijkstraNode[T]

func (dpq dijkstraPQ[T]) Len() int {
	return len(dpq)
}

func (dpq dijkstraPQ[T]) Less(i, j int) bool {
	return dpq[i].distance < dpq[j].distance
}

func (dpq dijkstraPQ[T]) Swap(i, j int) {
	dpq[i], dpq[j] = dpq[j], dpq[i]
	dpq[i].index, dpq[j].index = i, j
}

func (dpq *dijkstraPQ[T]) Push(x interface{}) {
	dn := x.(*dijkstraNode[T])
	dn.index = len(*dpq)
	*dpq = append(*dpq, dn)
}

func (dpq *dijkstraPQ[T]) Pop() interface{} {
	n := len(*dpq)
	node := (*dpq)[n-1]
	*dpq = (*dpq)[0 : n-1]
	return node
}

func Dijkstra[T Vertex](g *Graph[T], start, end T) *list.List {
	pq := dijkstraPQ[T]{}
	nodes := map[T]*dijkstraNode[T]{}

	heap.Init(&pq)

	g.EachVertex(func(v T, _ func()) {
		dn := &dijkstraNode[T]{
			vertex:   v,
			distance: math.Inf(1),
		}
		heap.Push(&pq, dn)
		nodes[v] = dn
	})

	nodes[start].distance = 0
	heap.Fix(&pq, nodes[start].index)

	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*dijkstraNode[T])

		g.EachHalfedge(v.vertex, func(he Halfedge[T], _ func()) {
			dn := nodes[he.End]

			if dn == nil {
				return
			}

			if v.distance+he.Cost < dn.distance {
				dn.distance = v.distance + he.Cost
				dn.predecessor = v
				heap.Fix(&pq, dn.index)
			}
		})

		if v.vertex == end {
			l := list.New()
			for e := v; e != nil; e = e.predecessor {
				l.PushFront(e.vertex)
			}
			return l
		}
	}

	return nil
}
