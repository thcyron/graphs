package graphs

type Network struct {
	Graph    *Graph
	Source   Vertex
	Sink     Vertex
	Flow     map[Edge]uint
	Capacity map[Edge]uint
}

func NewNetwork(graph *Graph, source, sink Vertex) *Network {
	return &Network{
		Graph:    graph,
		Source:   source,
		Sink:     sink,
		Flow:     make(map[Edge]uint),
		Capacity: make(map[Edge]uint),
	}
}

func (n *Network) SetCapacity(v, w Vertex, c uint) {
	e := Edge{v, w, 0}
	n.Capacity[e] = c
}

func (n *Network) SetFlow(v, w Vertex, f uint) {
	e := Edge{v, w, 0}
	n.Flow[e] = f
}

func (n *Network) SetFlowAndCapacity(v, w Vertex, f, c uint) {
	e := Edge{v, w, 0}
	n.Flow[e] = f
	n.Capacity[e] = c
}

func (n *Network) ResidualNetwork() *Network {
	rg := NewGraph()

	for v := range n.Graph.VerticesIter() {
		rg.AddVertex(v)
	}

	an := NewNetwork(rg, n.Source, n.Sink)

	for e := range n.Graph.EdgesIter() {
		f := n.Flow[e]

		if f > 0 {
			rg.AddEdge(e.End, e.Start, 0)
			an.Capacity[Edge{e.End, e.Start, 0}] = f
		}

		if c := n.Capacity[e]; f < c {
			rg.AddEdge(e.Start, e.End, 0)
			an.Capacity[Edge{e.Start, e.End, 0}] = c - f
		}
	}

	return an
}

func (n *Network) Equals(n2 *Network) bool {
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
