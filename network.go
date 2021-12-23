package graphs

type Network[T Vertex] struct {
	Graph    *Graph[T]
	Source   T
	Sink     T
	Flow     map[Edge[T]]uint
	Capacity map[Edge[T]]uint
}

func NewNetwork[T Vertex](graph *Graph[T], source, sink T) *Network[T] {
	return &Network[T]{
		Graph:    graph,
		Source:   source,
		Sink:     sink,
		Flow:     make(map[Edge[T]]uint),
		Capacity: make(map[Edge[T]]uint),
	}
}

func (n *Network[T]) SetCapacity(v, w T, c uint) {
	e := Edge[T]{v, w, 0}
	n.Capacity[e] = c
}

func (n *Network[T]) SetFlow(v, w T, f uint) {
	e := Edge[T]{v, w, 0}
	n.Flow[e] = f
}

func (n *Network[T]) SetFlowAndCapacity(v, w T, f, c uint) {
	e := Edge[T]{v, w, 0}
	n.Flow[e] = f
	n.Capacity[e] = c
}

func (n *Network[T]) ResidualNetwork() *Network[T] {
	rg := NewGraph[T]()

	for v := range n.Graph.VerticesIter() {
		rg.AddVertex(v)
	}

	an := NewNetwork(rg, n.Source, n.Sink)

	for e := range n.Graph.EdgesIter() {
		f := n.Flow[e]

		if f > 0 {
			rg.AddEdge(e.End, e.Start, 0)
			an.Capacity[Edge[T]{e.End, e.Start, 0}] = f
		}

		if c := n.Capacity[e]; f < c {
			rg.AddEdge(e.Start, e.End, 0)
			an.Capacity[Edge[T]{e.Start, e.End, 0}] = c - f
		}
	}

	return an
}

func (n *Network[T]) Equals(n2 *Network[T]) bool {
	if !n.Graph.Equals(n2.Graph) {
		return false
	}

	if n.Source != n2.Source || n.Sink != n2.Sink {
		return false
	}

	if len(n.Flow) != len(n2.Flow) {
		return false
	}

	if len(n.Capacity) != len(n2.Capacity) {
		return false
	}

	for k, v := range n.Flow {
		if vv, exists := n2.Flow[k]; !exists || v != vv {
			return false
		}
	}

	for k, v := range n.Capacity {
		if vv, exists := n2.Capacity[k]; !exists || v != vv {
			return false
		}
	}

	return true
}
