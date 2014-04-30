package graphs

import (
	"testing"
)

func TestKruskal(t *testing.T) {
	graph := NewGraph()

	graph.AddEdge("a", "b", 1)
	graph.AddEdge("a", "c", 3)
	graph.AddEdge("b", "g", 5)
	graph.AddEdge("c", "g", 8)
	graph.AddEdge("g", "h", 6)
	graph.AddEdge("c", "d", 2)
	graph.AddEdge("g", "f", 4)
	graph.AddEdge("d", "f", 3)
	graph.AddEdge("d", "e", 5)

	tree := Kruskal(graph)

	result := NewGraph()
	result.AddEdge("a", "b", 1)
	result.AddEdge("c", "d", 2)
	result.AddEdge("a", "c", 3)
	result.AddEdge("d", "f", 3)
	result.AddEdge("g", "f", 4)
	result.AddEdge("d", "e", 5)
	result.AddEdge("g", "h", 6)

	if !tree.Equals(result) {
		t.Error("bad minimal spanning tree")
	}
}
