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
	graph.AddEdge(13, 14, 0)

	var result Vertex
	walks := 0

	BFS(graph, 1, func(v Vertex, stop *bool) {
		walks++
		if i := v.(int); i > 10 && i%2 != 0 {
			result = v
			*stop = true
		}
	})

	if result != 13 {
		t.Error("bad result")
	}

	if walks != 6 {
		t.Error("should visit 6 vertices")
	}
}
