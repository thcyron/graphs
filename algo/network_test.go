package algo

import (
	"testing"

	. "github.com/noypi/graph/store"
	. "github.com/noypi/graph/types"
)

func NewInMemoryGraphNotDirectedTest() Graph {
	fn := func(bb []byte) (Vertex, error) {
		return VertexString(bb), nil
	}

	g, _ := NewGraphInMemory(fn, false)

	/*path, _ := ioutil.TempDir(os.TempDir(), "test-graph")
	g, err := NewGraphLeveldb(path, fn, true)
	if nil != err {
		panic(err.Error())
	}*/

	return g
}

func resultResidualNetwork() *Network {
	g := NewInMemoryGraphNotDirectedTest()
	g.AddEdge(VertexString("s"), VertexString("a"), 0)
	g.AddEdge(VertexString("s"), VertexString("b"), 0)
	g.AddEdge(VertexString("a"), VertexString("t"), 0)
	g.AddEdge(VertexString("b"), VertexString("t"), 0)
	g.AddEdge(VertexString("a"), VertexString("b"), 0)

	n := NewNetwork(g, VertexString("s"), VertexString("t"))
	n.SetCapacity(VertexString("s"), VertexString("a"), 2)
	n.SetCapacity(VertexString("b"), VertexString("t"), 2)
	n.SetCapacity(VertexString("a"), VertexString("b"), 2)
	n.SetCapacity(VertexString("t"), VertexString("a"), 1)
	n.SetCapacity(VertexString("t"), VertexString("b"), 1)
	n.SetCapacity(VertexString("a"), VertexString("s"), 1)
	n.SetCapacity(VertexString("b"), VertexString("s"), 1)

	return n
}

func TestNetworkResidualNetwork(t *testing.T) {
	g := NewInMemoryGraphNotDirectedTest()
	g.AddEdge(VertexString("s"), VertexString("a"), 0)
	g.AddEdge(VertexString("s"), VertexString("b"), 0)
	g.AddEdge(VertexString("a"), VertexString("t"), 0)
	g.AddEdge(VertexString("b"), VertexString("t"), 0)
	g.AddEdge(VertexString("a"), VertexString("b"), 0)

	n := NewNetwork(g, VertexString("s"), VertexString("t"))
	n.SetFlowAndCapacity(VertexString("s"), VertexString("a"), 1, 3)
	n.SetFlowAndCapacity(VertexString("s"), VertexString("b"), 1, 1)
	n.SetFlowAndCapacity(VertexString("a"), VertexString("t"), 1, 1)
	n.SetFlowAndCapacity(VertexString("b"), VertexString("t"), 1, 3)
	n.SetFlowAndCapacity(VertexString("a"), VertexString("b"), 0, 2)

	rn := n.ResidualNetwork()
	rrn2 := resultResidualNetwork()
	if !rn.Equals(rrn2) {
		t.Fatal("bad residual network")
	}
}
