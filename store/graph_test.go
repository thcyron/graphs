package store

import (
	"testing"

	. "github.com/noypi/graph/types"
)

func NewInMemoryGraphTest() Graph {
	fn := func(bb []byte) (Vertex, error) {
		return VertexString(bb), nil
	}

	g, _ := NewGraphInMemory(fn, true)

	/*path, _ := ioutil.TempDir(os.TempDir(), "test-graph")
	g, err := NewGraphLeveldb(path, fn, true)
	if nil != err {
		panic(err.Error())
	}*/

	return g
}

func TestGraphNVertices(t *testing.T) {
	graph := NewInMemoryGraphTest()

	if graph.NVertices() != 0 {
		t.Fatal("empty graph should not have any vertices")
	}

	graph.AddEdge(VertexString("a"), VertexString("b"), 0)
	if graph.NVertices() != 2 {
		t.Fatal("graph should have two vertices")
	}
}

func TestVerticesIter(t *testing.T) {
	graph := NewInMemoryGraphTest()

	graph.AddEdge(VertexString("a"), VertexString("b"), 0)
	graph.AddEdge(VertexString("b"), VertexString("c"), 0)

	vertices := 0

	it := graph.VerticesIter()
	for ; it.Valid(); it.Next() {
		vertices++
	}

	if vertices != graph.NVertices() {
		t.Fatal("wrong number of vertices")
	}
}

func TestGraphNEdges(t *testing.T) {
	graph := NewInMemoryGraphTest()

	if graph.NEdges() != 0 {
		t.Fatal("empty graph should not have any edges")
	}

	graph.AddEdge(VertexString("a"), VertexString("b"), 0)
	if graph.NEdges() != 1 {
		t.Fatal("graph should have one edge")
	}
}

func TestGraphEquals(t *testing.T) {
	g1 := NewInMemoryGraphTest()
	g2 := NewInMemoryGraphTest()

	if !g1.Equals(g2) {
		t.Fatal("two empty graphs should be equal")
	}

	g1.AddEdge(VertexString("a"), VertexString("b"), 0)
	if g1.Equals(g2) {
		t.Fatal("two graphs with different number of edges should not be equal")
	}

	g2.AddEdge(VertexString("a"), VertexString("b"), 0)
	if g1.NVertices() != g2.NVertices() {
		t.Fatal("two graphs with same edges should have equal nVertices")
	}
	if g1.NEdges() != g2.NEdges() {
		t.Fatal("two graphs with same edges should have equal nEdges")
	}

	a1 := g1.Adj()
	a2 := g2.Adj()

	it := a1.Iterator()
	for ; it.Valid(); it.Next() {
		set1 := it.Value()
		set2 := a2.EdgeSetOf(it.Key())

		if set1.Len() != set2.Len() {
			t.Fatal("set lengths should be equal")
		}
		if !it.Value().Equals(a2.EdgeSetOf(it.Key())) {
			t.Fatal("could not find it.Key=>", string(it.Key().StringID()))
		}
	}

	if !g1.Equals(g2) {
		t.Fatal("two graphs with same edges should be equal")
	}

	g1.AddEdge(VertexString("b"), VertexString("c"), 0)
	g2.AddEdge(VertexString("a"), VertexString("c"), 0)
	if g1.Equals(g2) {
		t.Fatal("two graphs with different edges should not be equal")
	}
}
