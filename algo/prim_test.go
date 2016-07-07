package algo

import (
	"testing"

	. "github.com/noypi/graph/types"
)

func TestPrim(t *testing.T) {
	graph := NewInMemoryGraphNotDirectedTest()

	graph.AddEdge(VertexString("a"), VertexString("b"), 8)
	graph.AddEdge(VertexString("a"), VertexString("c"), 5)
	graph.AddEdge(VertexString("b"), VertexString("c"), 10)
	graph.AddEdge(VertexString("b"), VertexString("d"), 2)
	graph.AddEdge(VertexString("b"), VertexString("e"), 18)
	graph.AddEdge(VertexString("c"), VertexString("d"), 3)
	graph.AddEdge(VertexString("c"), VertexString("f"), 16)
	graph.AddEdge(VertexString("d"), VertexString("e"), 12)
	graph.AddEdge(VertexString("d"), VertexString("f"), 30)
	graph.AddEdge(VertexString("d"), VertexString("g"), 14)
	graph.AddEdge(VertexString("e"), VertexString("g"), 4)
	graph.AddEdge(VertexString("f"), VertexString("g"), 26)

	tree := Prim(graph, VertexString("g"))
	if tree == nil {
		t.Error("no result")
		t.FailNow()
	}

	result := NewInMemoryGraphNotDirectedTest()
	result.AddEdge(VertexString("g"), VertexString("e"), 4)
	result.AddEdge(VertexString("e"), VertexString("d"), 12)
	result.AddEdge(VertexString("d"), VertexString("b"), 2)
	result.AddEdge(VertexString("d"), VertexString("c"), 3)
	result.AddEdge(VertexString("c"), VertexString("a"), 5)
	result.AddEdge(VertexString("c"), VertexString("f"), 16)

	if !tree.Equals(result) {
		t.Error("bad minimal spanning tree")
	}
}
