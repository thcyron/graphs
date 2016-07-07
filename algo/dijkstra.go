package algo

import (
	"container/heap"
	"container/list"
	"math"

	. "github.com/noypi/graph/types"
)

type dijkstraNode struct {
	vertex      Vertex
	distance    float64
	predecessor *dijkstraNode
	index       int // Index of the node in the heap
}

type dijkstraPQ []*dijkstraNode

func (dpq dijkstraPQ) Len() int {
	return len(dpq)
}

func (dpq dijkstraPQ) Less(i, j int) bool {
	return dpq[i].distance < dpq[j].distance
}

func (dpq dijkstraPQ) Swap(i, j int) {
	dpq[i], dpq[j] = dpq[j], dpq[i]
	dpq[i].index, dpq[j].index = i, j
}

func (dpq *dijkstraPQ) Push(x interface{}) {
	dn := x.(*dijkstraNode)
	dn.index = len(*dpq)
	*dpq = append(*dpq, dn)
}

func (dpq *dijkstraPQ) Pop() interface{} {
	n := len(*dpq)
	node := (*dpq)[n-1]
	*dpq = (*dpq)[0 : n-1]
	return node
}

func Dijkstra(g Graph, start, end Vertex) *list.List {
	pq := dijkstraPQ{}
	nodes := map[string]*dijkstraNode{}

	heap.Init(&pq)

	itv := g.VerticesIter()
	for ; itv.Valid(); itv.Next() {
		v := itv.Value()
		dn := &dijkstraNode{
			vertex:   v,
			distance: math.Inf(1),
		}
		heap.Push(&pq, dn)
		nodes[v.StringID()] = dn
	}

	nodes[start.StringID()].distance = 0
	heap.Fix(&pq, nodes[start.StringID()].index)

	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*dijkstraNode)

		ithe := g.HalfedgesIter(v.vertex)
		for ; ithe.Valid(); ithe.Next() {
			he := ithe.Value()
			dn := nodes[he.End()]

			if dn == nil {
				continue
			}

			if v.distance+he.Cost() < dn.distance {
				dn.distance = v.distance + he.Cost()
				dn.predecessor = v
				heap.Fix(&pq, dn.index)
			}
		}

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
