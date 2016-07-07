package algo

import (
	"testing"

	. "github.com/noypi/graph/types"
)

func TestKruskal(t *testing.T) {
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

	tree := Kruskal(graph)

	result := NewInMemoryGraphTest()
	result.AddEdge(VertexString("a"), VertexString("b"), 1)
	result.AddEdge(VertexString("c"), VertexString("d"), 2)
	result.AddEdge(VertexString("a"), VertexString("c"), 3)
	result.AddEdge(VertexString("d"), VertexString("f"), 3)
	result.AddEdge(VertexString("g"), VertexString("f"), 4)
	result.AddEdge(VertexString("d"), VertexString("e"), 5)
	result.AddEdge(VertexString("g"), VertexString("h"), 6)

	if !tree.Equals(result) {
		t.Error("bad minimal spanning tree")
	}
}
