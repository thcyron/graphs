package algo

import (
	"testing"

	. "github.com/noypi/graph/types"
)

func TestDijkstra(t *testing.T) {
	graph := NewInMemoryGraphTest()

	graph.AddEdge(VertexString("a"), VertexString("b"), 1)
	graph.AddEdge(VertexString("a"), VertexString("c"), 3)
	graph.AddEdge(VertexString("b"), VertexString("g"), 5)
	graph.AddEdge(VertexString("c"), VertexString("g"), 8)
	graph.AddEdge(VertexString("g"), VertexString("h"), 6)
	graph.AddEdge(VertexString("c"), VertexString("d"), 2)
	graph.AddEdge(VertexString("g"), VertexString("f"), 4)
	graph.AddEdge(VertexString("d"), VertexString("f"), 3)
	graph.AddEdge(VertexString("d"), VertexString("e"), 5)

	path := Dijkstra(graph, VertexString("a"), VertexString("e"))
	result := []Vertex{VertexString("a"), VertexString("c"), VertexString("d"), VertexString("e")}

	i := 0
	for e := path.Front(); e != nil; e = e.Next() {
		if e.Value != result[i] {
			t.Errorf("bad vertex in path at index %d", i)
		}
		i++
	}

	if i != len(result) {
		t.Error("bad path")
	}
}
