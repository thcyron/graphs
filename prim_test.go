package graphs

import (
	"testing"
)

func TestPrim(t *testing.T) {
	graph := NewGraph[string]()

	graph.AddEdge("a", "b", 8)
	graph.AddEdge("a", "c", 5)
	graph.AddEdge("b", "c", 10)
	graph.AddEdge("b", "d", 2)
	graph.AddEdge("b", "e", 18)
	graph.AddEdge("c", "d", 3)
	graph.AddEdge("c", "f", 16)
	graph.AddEdge("d", "e", 12)
	graph.AddEdge("d", "f", 30)
	graph.AddEdge("d", "g", 14)
	graph.AddEdge("e", "g", 4)
	graph.AddEdge("f", "g", 26)

	tree := Prim(graph, "g")
	if tree == nil {
		t.Error("no result")
		t.FailNow()
	}

	result := NewGraph[string]()
	result.AddEdge("g", "e", 4)
	result.AddEdge("e", "d", 12)
	result.AddEdge("d", "b", 2)
	result.AddEdge("d", "c", 3)
	result.AddEdge("c", "a", 5)
	result.AddEdge("c", "f", 16)

	if !tree.Equals(result) {
		t.Error("bad minimal spanning tree")
	}
}
