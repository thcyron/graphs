package algo

import (
	"github.com/noypi/graph/store"
	. "github.com/noypi/graph/types"
)

type Network struct {
	Graph    Graph
	Source   Vertex
	Sink     Vertex
	Flow     map[string]uint // by edge id
	Capacity map[string]uint // by edge id
}

func NewNetwork(graph Graph, source, sink Vertex) *Network {
	return &Network{
		Graph:    graph,
		Source:   source,
		Sink:     sink,
		Flow:     make(map[string]uint),
		Capacity: make(map[string]uint),
	}
}

func (n *Network) SetCapacity(v, w Vertex, c uint) {
	e := EdgeBase{S: v.StringID()}
	e.E = w.StringID()
	e.C = 0
	n.Capacity[e.StringID()] = c
}

func (n *Network) SetFlow(v, w Vertex, f uint) {
	e := EdgeBase{S: v.StringID()}
	e.E = w.StringID()
	e.C = 0
	n.Flow[e.StringID()] = f
}

func (n *Network) SetFlowAndCapacity(v, w Vertex, f, c uint) {
	e := EdgeBase{S: v.StringID()}
	e.E = w.StringID()
	e.C = 0
	n.Flow[e.StringID()] = f
	n.Capacity[e.StringID()] = c
}

func (n *Network) ResidualNetwork() *Network {
	rg, _ := store.NewGraphInMemory(n.Graph.GetVertexDeserializer(), n.Graph.IsDirected())

	itv := n.Graph.VerticesIter()
	for ; itv.Valid(); itv.Next() {
		v := itv.Value()
		rg.AddVertex(v)
	}

	an := NewNetwork(rg, n.Source, n.Sink)

	ite := n.Graph.EdgesIter()
	for ; ite.Valid(); ite.Next() {
		e := ite.Value()
		f := n.Flow[e.StringID()]

		if f > 0 {
			rg.AddEdge(rg.V(e.End()), rg.V(e.Start()), 0)
			an.Capacity[NewEdge(e.End(), e.Start(), 0).StringID()] = f
		}

		if c := n.Capacity[e.StringID()]; f < c {
			rg.AddEdge(rg.V(e.Start()), rg.V(e.End()), 0)
			an.Capacity[NewEdge(e.Start(), e.End(), 0).StringID()] = c - f
		}
	}

	return an
}

func (n *Network) Equals(n2 *Network) bool {
	if !n.Graph.Equals(n2.Graph) {
		return false
	}

	if n.Source.StringID() != n2.Source.StringID() || n.Sink.StringID() != n2.Sink.StringID() {
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
