package algo

import (
	"container/heap"
	"math"

	"github.com/noypi/graph/store"
	. "github.com/noypi/graph/types"
)

type primNode struct {
	vertex      Vertex
	cost        float64
	index       int // Index of the node in the heap
	visited     bool
	predecessor *primNode
}

type primPQ []*primNode

func (pq primPQ) Len() int {
	return len(pq)
}

func (pq primPQ) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq primPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

func (pq *primPQ) Push(x interface{}) {
	n := x.(*primNode)
	n.index = len(*pq)
	*pq = append(*pq, n)
}

func (pq *primPQ) Pop() interface{} {
	n := len(*pq)
	node := (*pq)[n-1]
	*pq = (*pq)[0 : n-1]
	return node
}

// Prim implements Prim’s algorithm. It returns a minimal spanning
// tree for the given graph, starting with vertex start.
func Prim(g Graph, start Vertex) Graph {
	tree, _ := store.NewGraphInMemory(g.GetVertexDeserializer(), g.IsDirected())
	nodes := map[string]*primNode{}
	pq := primPQ{}

	heap.Init(&pq)

	itv := g.VerticesIter()
	for ; itv.Valid(); itv.Next() {
		v := itv.Value()
		n := &primNode{
			vertex:  v,
			cost:    math.Inf(1),
			visited: false,
		}
		heap.Push(&pq, n)
		nodes[v.StringID()] = n
	}

	nodes[start.StringID()].cost = 0
	heap.Fix(&pq, nodes[start.StringID()].index)

	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*primNode)
		v.visited = true

		ithe := g.HalfedgesIter(v.vertex)
		for ; ithe.Valid(); ithe.Next() {
			he := ithe.Value()
			node := nodes[he.End()]
			if node.visited {
				continue
			}

			if he.Cost() < node.cost {
				node.cost = he.Cost()
				node.predecessor = v
				heap.Fix(&pq, node.index)
			}
		}
	}

	for _, node := range nodes {
		if node.predecessor != nil {
			tree.AddEdge(node.predecessor.vertex, node.vertex, node.cost)
		}
	}

	return tree
}
