package graphs

import (
	"testing"
)

func TestBFS(t *testing.T) {
	graph := NewGraph()

	graph.AddEdge(1, 3, 0)
	graph.AddEdge(1, 2, 0)
	graph.AddEdge(3, 8, 0)
	graph.AddEdge(2, 12, 0)
	graph.AddEdge(12, 13, 0)

	predicate := func(v Vertex) bool {
		i := v.(int)
		return i > 10 && i%2 != 0
	}

	if result := BFS(graph, 1, predicate); result != 13 {
		t.Error("bad result")
	}
}
