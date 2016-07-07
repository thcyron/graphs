package algo

import (
	"testing"

	. "github.com/noypi/graph/types"
)

func TestFloydWarshall(t *testing.T) {
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

	m := FloydWarshall(graph)
	if c := m["a"]["e"]; c != 10 {
		t.Errorf("bad shortest cost %f for a-e", c)
	}
}
