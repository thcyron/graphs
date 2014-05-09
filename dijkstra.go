package graphs

import (
	"container/heap"
	"container/list"
	"math"
)

type djikstraNode struct {
	vertex      Vertex
	distance    float64
	predecessor *djikstraNode
	index       int // Index of the node in the heap
}

type djikstraPQ []*djikstraNode

func (dpq djikstraPQ) Len() int {
	return len(dpq)
}

func (dpq djikstraPQ) Less(i, j int) bool {
	return dpq[i].distance < dpq[j].distance
}

func (dpq djikstraPQ) Swap(i, j int) {
	dpq[i], dpq[j] = dpq[j], dpq[i]
	dpq[i].index, dpq[j].index = i, j
}

func (dpq *djikstraPQ) Push(x interface{}) {
	dn := x.(*djikstraNode)
	dn.index = len(*dpq)
	*dpq = append(*dpq, dn)
}

func (dpq *djikstraPQ) Pop() interface{} {
	n := len(*dpq)
	node := (*dpq)[n-1]
	*dpq = (*dpq)[0 : n-1]
	return node
}

func Dijkstra(g *Graph, start, end Vertex) *list.List {
	pq := djikstraPQ{}
	nodes := map[Vertex]*djikstraNode{}

	heap.Init(&pq)

	g.Vertices.Each(func(e interface{}, stop *bool) {
		v := e.(Vertex)
		dn := &djikstraNode{
			vertex:   v,
			distance: math.Inf(1),
		}
		heap.Push(&pq, dn)
		nodes[v] = dn
	})

	nodes[start].distance = 0
	heap.Fix(&pq, nodes[start].index)

	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*djikstraNode)

		for he := range g.HalfedgesIter(v.vertex) {
			dn := nodes[he.End]

			if dn == nil {
				continue
			}

			if math.IsInf(dn.distance, 1) {
				dn.distance = he.Cost
				dn.predecessor = v
				heap.Fix(&pq, dn.index)
			} else {
				newCost := dn.distance + he.Cost
				if newCost < dn.distance {
					dn.distance = newCost
					dn.predecessor = v
					heap.Fix(&pq, dn.index)
				}
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
