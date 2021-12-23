package graphs

import (
	"testing"
)

func TestBellmanFord(t *testing.T) {
	graph := NewDigraph[string]()

	graph.AddEdge("a", "b", 1)
	graph.AddEdge("a", "c", 3)
	graph.AddEdge("b", "g", 5)
	graph.AddEdge("c", "g", 8)
	graph.AddEdge("g", "h", 6)
	graph.AddEdge("c", "d", -2)
	graph.AddEdge("g", "f", 4)
	graph.AddEdge("d", "f", 3)
	graph.AddEdge("d", "e", 5)

	path := BellmanFord(graph, "a", "e")
	if path == nil {
		t.Error("no result")
		t.FailNow()
	}

	result := []string{"a", "c", "d", "e"}
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
	graph := NewDigraph[string]()

	graph.AddEdge("a", "b", 6)
	graph.AddEdge("a", "c", 3)
	graph.AddEdge("c", "a", -4)

	path := BellmanFord(graph, "a", "b")
	if path != nil {
		t.Error("should return no result (negative weight cycle)")
	}
}
