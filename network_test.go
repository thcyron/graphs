package graphs

import (
	"testing"
)

func resultResidualNetwork() *Network {
	g := NewGraph()
	g.AddEdge("s", "a", 0)
	g.AddEdge("s", "b", 0)
	g.AddEdge("a", "t", 0)
	g.AddEdge("b", "t", 0)
	g.AddEdge("a", "b", 0)

	n := NewNetwork(g, "s", "t")
	n.SetCapacity("s", "a", 2)
	n.SetCapacity("b", "t", 2)
	n.SetCapacity("a", "b", 2)
	n.SetCapacity("t", "a", 1)
	n.SetCapacity("t", "b", 1)
	n.SetCapacity("a", "s", 1)
	n.SetCapacity("b", "s", 1)

	return n
}

func TestNetworkResidualNetwork(t *testing.T) {
	g := NewGraph()
	g.AddEdge("s", "a", 0)
	g.AddEdge("s", "b", 0)
	g.AddEdge("a", "t", 0)
	g.AddEdge("b", "t", 0)
	g.AddEdge("a", "b", 0)

	n := NewNetwork(g, "s", "t")
	n.SetFlowAndCapacity("s", "a", 1, 3)
	n.SetFlowAndCapacity("s", "b", 1, 1)
	n.SetFlowAndCapacity("a", "t", 1, 1)
	n.SetFlowAndCapacity("b", "t", 1, 3)
	n.SetFlowAndCapacity("a", "b", 0, 2)

	rn := n.ResidualNetwork()

	if !rn.Equals(resultResidualNetwork()) {
		t.Error("bad residual network")
	}
}
