package algo

import (
	"testing"

	. "github.com/noypi/graph/store"
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

func TestBellmanFord(t *testing.T) {
	graph := NewInMemoryGraphTest()

	graph.AddEdge(VertexString("a"), VertexString("b"), 1)
	graph.AddEdge(VertexString("a"), VertexString("c"), 3)
	graph.AddEdge(VertexString("b"), VertexString("g"), 5)
	graph.AddEdge(VertexString("c"), VertexString("g"), 8)
	graph.AddEdge(VertexString("g"), VertexString("h"), 6)
	graph.AddEdge(VertexString("c"), VertexString("d"), -2)
	graph.AddEdge(VertexString("g"), VertexString("f"), 4)
	graph.AddEdge(VertexString("d"), VertexString("f"), 3)
	graph.AddEdge(VertexString("d"), VertexString("e"), 5)

	path := BellmanFord(graph, VertexString("a"), VertexString("e"))
	if path == nil {
		t.Error("no result")
		t.FailNow()
	}

	result := []Vertex{VertexString("a"), VertexString("c"), VertexString("d"), VertexString("e")}
	if len(path) != len(result) {
		t.Error("bad result")
		t.FailNow()
	}

	for i, v := range path {
		if v != result[i] {
			t.Errorf("bad vertex in path at index %d", i)
		}
	}
}

func TestBellmanFordNegWeightCycle(t *testing.T) {
	graph := NewInMemoryGraphTest()

	graph.AddEdge(VertexString("a"), VertexString("b"), 6)
	graph.AddEdge(VertexString("a"), VertexString("c"), 3)
	graph.AddEdge(VertexString("c"), VertexString("a"), -4)

	path := BellmanFord(graph, VertexString("a"), VertexString("b"))
	if path != nil {
		t.Error("should return no result (negative weight cycle)")
	}
}
