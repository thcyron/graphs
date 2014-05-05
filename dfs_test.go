package graphs

import "testing"

func TestDFS(t *testing.T) {
	graph := NewDigraph()

	graph.AddEdge(1, 2, 0)
	graph.AddEdge(2, 3, 0)
	graph.AddEdge(3, 4, 0)
	graph.AddEdge(1, 5, 0)
	graph.AddEdge(5, 6, 0)
	graph.AddEdge(6, 3, 0)
	graph.AddEdge(1, 7, 0)

	walks := 0
	DFS(graph, 1, func(v Vertex, stop *bool) {
		walks++
	})

	if walks != 7 {
		t.Errorf("should visit 7 vertices; visited %d", walks)
	}

	walks = 0
	DFS(graph, 1, func(v Vertex, stop *bool) {
		walks++
		if v == 5 {
			*stop = true
		}
	})

	if walks != 5 {
		t.Errorf("should visit 5 vertices; visited %d", walks)
	}
}
